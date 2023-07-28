package validator

const (
	required     = "required"
	alpha        = "alpha"
	alphaNumeric = "alphaNumeric"
	numeric      = "numeric"
	number       = "number"
	hexadecimal  = "hexadecimal"
	hexColor     = "hexColor"
	rgb          = "rgb"
	rgba         = "rgba"
	hsl          = "hsl"
	hsla         = "hsla"
	email        = "email"
	cron         = "cron"
	min          = "min"
	max          = "max"
	length       = "len"
	minLen       = "minLen"
	maxLen       = "maxLen"
	match        = "match"
	oneOf        = "oneOf"
	// TODO implement the rest
	in      = "in"
	out     = "out"
	include = "include"
	exclude = "exclude"
)

const (
	typeString    = "string"
	typeStringPtr = "*string"
	typeInt       = "int"
	typeIntPtr    = "*int"
	typeUint      = "uint"
	typeUintPtr   = "*uint"
	typeFloat     = "float"
	typeFloatPtr  = "*float"
)
