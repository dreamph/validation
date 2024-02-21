package validation

import (
	"github.com/dreamph/validation/rules"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func NewStructValidationBuilder(structPtr interface{}, fields ...*FieldRules) *Builder {
	return &Builder{structPtr: structPtr, fields: fields}
}

type Builder struct {
	structPtr      any
	fields         []*FieldRules
	requiredFields []*FieldRules
}

func (v *Builder) AddFieldRules(fields ...*FieldRules) *Builder {
	if len(fields) > 0 {
		v.fields = append(v.fields, fields...)
	}
	return v
}

func (v *Builder) AddRequiredFieldRules(fields ...*FieldRules) *Builder {
	if len(fields) > 0 {
		v.requiredFields = append(v.requiredFields, fields...)
	}
	return v
}

func (v *Builder) Validate() error {
	if len(v.requiredFields) > 0 {
		err := ValidateStruct(v.structPtr,
			v.requiredFields...,
		)
		if err != nil {
			return err
		}
	}

	return ValidateStruct(v.structPtr,
		v.fields...,
	)
}

func StructField[T any](fieldPtr any, rule func(value T) error) *FieldRules {
	return validation.Field(fieldPtr, rules.ByRule(rule))
}

func ArrayField[T any](fieldPtr any, rule func(value T, i int) error) *FieldRules {
	return validation.Field(fieldPtr, rules.ByRuleArray(rule))
}
