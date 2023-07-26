package validator_test

import (
	"testing"

	"github.com/oSethoum/validator"
)

type User struct {
	Name   string `validate:"minLen=3;maxLen=10"`
	Email  string `validate:"minLen=5;email"`
	Age    int    `validate:"min=17;max=35"`
	Status string `validate:"oneOf=active,away"`
}

func TestStruct(T *testing.T) {
	err := validator.Struct(User{Name: "cccc", Status: "away", Email: "kkk@gmail.com", Age: 18})
	if err != nil {
		T.Error(err)
	}
}
