package validator

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

type NameCase = string

const (
	Camel  NameCase = "camel"
	Pascal NameCase = "pascal"
)

type Constraint struct {
	Tag   string `json:"-"`
	Kind  string `json:"kind,omitempty"`
	Param any    `json:"param,omitempty"`
}

type FieldError struct {
	Field      string       `json:"field,omitempty"`
	Value      any          `json:"value,omitempty"`
	Struct     string       `json:"struct,omitempty"`
	Violations []Constraint `json:"violations,omitempty"`
}

type Error struct {
	FieldsErrors []FieldError `json:"fieldsErrors,omitempty"`
}

func (e *Error) Error() string {
	errs := []string{}
	for _, err := range e.FieldsErrors {
		vErrs := []string{}
		for _, v := range err.Violations {
			vErrs = append(vErrs, fmt.Sprintf("kind: %s, param: %+v", v.Kind, v.Param))
		}
		errs = append(errs, fmt.Sprintf("\nfield: %s \nstruct: %s \nvalue: %+v \nviolations: %s\n", err.Field, err.Struct, err.Value, strings.Join(vErrs, " | ")))
	}
	return strings.Join(errs, "")
}

func Struct(s any, nameCase ...string) error {
	nc := Camel
	if len(nameCase) > 0 {
		nc = nameCase[0]
	}

	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)
	if t.Kind() == reflect.Pointer {
		t = t.Elem()
		v = v.Elem()
	}
	fieldsErrors := []FieldError{}
	for i := 0; i < t.NumField(); i++ {
		ft := t.Field(i)
		fv := v.Field(i)

		if ft.Anonymous && ft.Type.Kind() == reflect.Struct {
			err := Struct(fv.Interface())
			if err != nil {
				e := err.(*Error)
				fieldsErrors = append(fieldsErrors, e.FieldsErrors...)
			}
		}

		tag, ok := ft.Tag.Lookup("validate")
		if !ok {
			continue
		}

		constraints := parseConstraints(tag)
		fieldError := FieldError{
			Field:  ft.Name,
			Struct: t.Name(),
		}
		if nc == Camel {
			fieldError.Field = camel(ft.Name)
		}
		violations := []Constraint{}

		if len(constraints) == 0 {
			continue
		}

		if isString(ft.Type) {
			value, ok := getStringValue(v.Field(i))
			if !ok {
				if strings.Contains(tag, required) {
					violations = append(violations, Constraint{
						Tag:  required,
						Kind: required,
					})
				}
			} else {
				fieldError.Value = value
				for _, constraint := range constraints {
					if constraint.Param != nil {
						switch constraint.Kind {
						case minLen:
							param, ok := getIntParam(constraint.Param)
							if !ok {
								panic(fmt.Sprintf("validate: struct %s field %s tag %s invalid param %v", t.Name(), ft.Name, constraint.Tag, constraint.Param))
							}
							constraint.Param = param
							if int64(len(value)) < param {
								violations = append(violations, constraint)
							}
						case maxLen:
							param, ok := getIntParam(constraint.Param)
							if !ok {
								panic(fmt.Sprintf("validate: struct %s field %s tag %s invalid param %v", t.Name(), ft.Name, constraint.Tag, constraint.Param))
							}
							constraint.Param = param
							if int64(len(value)) > param {
								violations = append(violations, constraint)
							}
						case length:
							param, ok := getIntParam(constraint.Param)
							if !ok {
								panic(fmt.Sprintf("validate: struct %s field %s tag %s invalid param %v", t.Name(), ft.Name, constraint.Tag, constraint.Param))
							}
							constraint.Param = param
							if int64(len(value)) != param {
								violations = append(violations, constraint)
							}

						case oneOf:
							param := getOneOfString(constraint.Param)
							if param == nil {
								panic(fmt.Sprintf("validate: struct %s field %s tag %s invalid param %v", t.Name(), ft.Name, constraint.Tag, constraint.Param))
							}
							constraint.Param = param
							if !inArray(param, value) {
								violations = append(violations, constraint)
							}
						}
					} else {
						value, _ := getStringValue(v.Field(i))
						if exp, ok := regexMap[constraint.Kind]; ok {
							if !exp.MatchString(value) {
								violations = append(violations, constraint)
							}
						} else {
							param, ok := getStringParam(constraint.Param)
							if ok {
								panic(fmt.Sprintf("validate: struct %s field %s tag %s invalid param %v", t.Name(), ft.Name, constraint.Tag, constraint.Param))
							}
							switch constraint.Kind {
							case match:
								if !regexp.MustCompile(param).MatchString(value) {
									violations = append(violations, constraint)
								}
							}
						}
					}

				}
			}
		}

		if isInt(ft.Type) {
			value, ok := getIntValue(v.Field(i))
			if !ok {
				if strings.Contains(tag, required) {
					violations = append(violations, Constraint{
						Tag:  required,
						Kind: required,
					})
				}
			} else {
				for _, constraint := range constraints {
					switch constraint.Kind {
					case min:
						param, ok := getIntParam(constraint.Param)
						if !ok {
							panic(fmt.Sprintf("validate: struct %s field %s tag %s invalid param %v", t.Name(), ft.Name, constraint.Tag, constraint.Param))
						}
						constraint.Param = param
						if value <= param {
							violations = append(violations, constraint)
						}
					case max:
						param, ok := getIntParam(constraint.Param)
						if !ok {
							panic(fmt.Sprintf("validate: struct %s field %s tag %s invalid param %v", t.Name(), ft.Name, constraint.Tag, constraint.Param))
						}
						constraint.Param = param
						if value > param {
							violations = append(violations, constraint)
						}
					}
				}
			}
			fieldError.Value = value
		}

		if isUint(ft.Type) {
			value, ok := getUintValue(v.Field(i))
			if !ok {
				if strings.Contains(tag, required) {
					violations = append(violations, Constraint{
						Tag:  required,
						Kind: required,
					})
				}
			} else {
				for _, constraint := range constraints {
					switch constraint.Kind {
					case min:
						param, ok := getUintParam(constraint.Param)
						if !ok {
							panic(fmt.Sprintf("validate: struct %s field %s tag %s invalid param %v", t.Name(), ft.Name, constraint.Tag, constraint.Param))
						}
						constraint.Param = param
						if value <= param {
							violations = append(violations, constraint)
						}
					case max:
						param, ok := getUintParam(constraint.Param)
						if !ok {
							panic(fmt.Sprintf("validate: struct %s field %s tag %s invalid param %v", t.Name(), ft.Name, constraint.Tag, constraint.Param))
						}
						constraint.Param = param
						if value > param {
							violations = append(violations, constraint)
						}
					}
				}
			}
			fieldError.Value = value
		}

		if isFloat(ft.Type) {
			value, ok := getFloatValue(v.Field(i))
			if !ok {
				if strings.Contains(tag, required) {
					violations = append(violations, Constraint{
						Tag:  required,
						Kind: required,
					})
				}
			} else {
				for _, constraint := range constraints {
					switch constraint.Kind {
					case min:
						param, ok := getFloatParam(constraint.Param)
						if !ok {
							panic(fmt.Sprintf("validate: struct %s field %s tag %s invalid param %v", t.Name(), ft.Name, constraint.Tag, constraint.Param))
						}
						constraint.Param = param
						if value <= param {
							violations = append(violations, constraint)
						}
					case max:
						param, ok := getFloatParam(constraint.Param)
						if !ok {
							panic(fmt.Sprintf("validate: struct %s field %s tag %s invalid param %v", t.Name(), ft.Name, constraint.Tag, constraint.Param))
						}
						constraint.Param = param
						if value > param {
							violations = append(violations, constraint)
						}
					}
				}
			}
			fieldError.Value = value
		}
		if len(violations) > 0 {
			fieldError.Violations = violations
			fieldsErrors = append(fieldsErrors, fieldError)
		}
	}

	if len(fieldsErrors) > 0 {
		return &Error{
			FieldsErrors: fieldsErrors,
		}
	}
	return nil
}
