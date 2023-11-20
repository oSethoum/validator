**Validator - Validation package for Go**

[![License](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

Validator is a powerful and flexible Go package that simplifies the task of validating struct fields using struct tags. With this package, you can easily define validation rules for your struct fields by adding simple and intuitive struct tags.

## Installation

To use Validator in your Go project, you need to have Go installed and set up on your system. Once you have Go ready, you can install Validator using the `go get` command:

```bash
go get github.com/oSethoum/validator
```

## Features

-  **Easy-to-Use**: Validator is designed to be user-friendly and straightforward. You can define validation rules directly in your struct using struct tags, making the code clean and maintainable.

-  **Extensible**: The package is designed to be easily extensible, allowing you to create custom validation functions for specific use cases.

-  **Field-Level Validation**: Validator operates on a per-field basis, enabling you to apply different validation rules to different struct fields.

-  **Customizable Error Messages**: You can customize the error messages returned by Validator to provide more context and clarity to end-users.

-  **Zero Dependencies**: Validator has no external dependencies, keeping your project's dependency tree clean and tidy.

## How to Use

Using Validator in your Go project is simple. First, import the `validator` package into your code:

```go
import "github.com/oSethoum/validator"
```

Next, add struct tags to the fields you want to validate. These struct tags should be in the format `validate:"rule"`, where `rule` represents the validation rule you want to apply to that field.

Here's an example of a struct definition with validation rules:

```go
type User struct {
    ID       int    `validate:"max=25"`
    Name     string `validate:"required,min=2,max=50"`
    Email    string `validate:"required,email"`
    Age      int    `validate:"required,min=18"`
}
```

In this example, we have defined a `User` struct with four fields: `ID`, `Name`, `Email`, and `Age`. We've applied various validation rules to each field:

-  The `ID` field must be present (required).
-  The `Name` field must be present, and its length must be between 2 and 50 characters (inclusive).
-  The `Email` field must be present and must be a valid email address.
-  The `Age` field must be present and greater than 18.

Once you have defined your struct and added the necessary struct tags, you can use the `validator.Validate` function to validate the struct:

```go
user := User{
    ID:    1,
    Name:  "John Doe",
    Email: "john.doe@example.com",
    Age:   25,
}

err := validator.Struct(user)
if err != nil {
    // Handle validation errors
    fmt.Println("Validation failed:", err)
    return
}

// Proceed with your business logic for the valid user data.
fmt.Println("User data is valid.")
```

That's it! You've successfully integrated the Validator package into your project, and now your struct fields will be automatically validated based on the defined rules.

**Supported Validations**

Validator package supports the following validation rules. These rules can be used as struct tags to specify the validation criteria for individual struct fields:

-  `required`: The field must be present and cannot be empty or zero-value.
-  `alpha`: The field must contain only alphabetical characters (a-z and A-Z).
-  `alphaSpace`: The field must contain only alphabetical characters (a-z and A-Z) and optional space to separate between them.
-  `alphaNumeric`: The field must contain only alphanumeric characters (a-z, A-Z, and 0-9).
-  `numeric`: The field must contain only numeric characters (0-9).
-  `number`: The field must be a valid number (supports negative numbers, decimals, and scientific notation).
-  `hexadecimal`: The field must be a valid hexadecimal number (e.g., "0x1A", "1a").
-  `hexColor`: The field must be a valid hexadecimal color code (e.g., "#FF0000", " #f00").
-  `rgb`: The field must be a valid RGB color (e.g., "rgb(255, 0, 0)").
-  `rgba`: The field must be a valid RGBA color (e.g., "rgba(255, 0, 0, 0.5)").
-  `hsl`: The field must be a valid HSL color (e.g., "hsl(0, 100%, 50%)").
-  `hsla`: The field must be a valid HSLA color (e.g., "hsla(0, 100%, 50%, 0.5)").
-  `email`: The field must be a valid email address.
-  `cron`: The field must be a valid cron expression.
-  `min`: The field must be a numeric value and greater than or equal to the specified minimum value.
-  `max`: The field must be a numeric value and less than or equal to the specified maximum value.
-  `len`: The field must have a length equal to the specified value (applicable to strings, arrays, slices, maps).
-  `minLen`: The field must have a length greater than or equal to the specified minimum value (applicable to strings, arrays, slices, maps).
-  `maxLen`: The field must have a length less than or equal to the specified maximum value (applicable to strings, arrays, slices, maps).
-  `match`: The field must match the provided regular expression pattern.
-  `oneOf`: The field value must be one of the specified valid values.
-  `in`: The field value must be in the values of the param.
-  `out`: The field value must not contain any of the param values.
-  `include`: the value must include all values of the param list.
-  `exclude`: the value must not include any of the param list values

Each validation rule can be combined with other rules and options using commas. For example, to apply multiple validations to a field, you can use:

```go
type User struct {
    Name  string `validate:"required,alpha,minLen=3,maxLen=50"`
    Email string `validate:"required,email"`
    Age   int    `validate:"required,numeric,min=18"`
}
```

The above struct specifies that the `Name` field must be present, contain only alphabetical characters, and have a minimum length of 3 and a maximum length of 50 characters. The `Email` field must be present and be a valid email address, while the `Age` field must be present, contain only numeric characters, and be greater than or equal to 18.

---

Generated by Chat Gpt
