package main

import (
	"fmt"

	"github.com/dreamph/validation"
)

type Wrapper struct {
	Attr1      int          `json:"attr1"`
	Type       int          `json:"type"`
	FieldOne   FieldOne     `json:"fieldOne"`
	FieldTwo   *FieldTwo    `json:"fieldTwo"`
	FieldThree []FieldThree `json:"fieldThree"`
}

type FieldOne struct {
	FieldThree string `json:"fieldThree"`
	FieldFour  string `json:"fieldFour"`
}

type FieldTwo struct {
	FieldFive int `json:"fieldFive"`
}

type FieldThree struct {
	Data int `json:"data"`
}

func main() {
	request := Wrapper{
		Attr1: 0,
		Type:  3,

		FieldOne: FieldOne{
			FieldThree: "Test",
			FieldFour:  "",
		},

		FieldTwo: &FieldTwo{
			FieldFive: 16,
		},

		FieldThree: []FieldThree{
			{Data: -1},
		},
	}

	validationBuilder := validation.NewStructValidationBuilder(&request)
	validationBuilder.AddRequiredFieldRules(
		validation.Field(&request.Type, validation.Required, validation.In(1, 2, 3)),
	)

	validationBuilder.AddFieldRules(
		validation.Field(&request.Attr1, validation.Required),
	)

	validationBuilder.AddFieldRules(
		validation.StructField[FieldOne](&request.FieldOne, func(value FieldOne) error {
			return validation.ValidateStruct(&value,
				validation.Field(&value.FieldThree, validation.Required),
				validation.Field(&value.FieldFour, validation.Required),
			)
		}),
		validation.StructField[FieldTwo](&request.FieldTwo, func(value FieldTwo) error {
			return validation.ValidateStruct(&value,
				validation.Field(&value.FieldFive, validation.Required),
			)
		}),
		validation.ArrayField[FieldThree](&request.FieldThree, func(value FieldThree, i int) error {
			return validation.ValidateStruct(&value,
				validation.Field(&value.Data, validation.Required, validation.Min(0)),
			)
		}),
	)

	err := validationBuilder.Validate()
	if err != nil {
		fmt.Println(err)
	}
}
