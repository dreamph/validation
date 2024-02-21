package validation

import (
	"strings"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/pkg/errors"
)

type DynamicRule struct {
	RuleKey         string
	ValidationRules map[string][]validation.Rule
	FormatData      func(value string) string
}

func (d DynamicRule) Get() []validation.Rule {
	return d.ValidationRules[d.RuleKey]
}

type validateRule struct {
	rule func() error
}

func WithRule(rule func() error) validation.Rule {
	return &validateRule{rule: rule}
}

func (l *validateRule) Validate(value interface{}) error {
	return l.rule()
}

func validateInRule(equalsIgnoreCase bool, values ...string) RuleFunc {
	return func(value interface{}) error {
		valStr, ok := value.(string)
		if !ok {
			return errors.New("must be a string")
		}
		if valStr == "" {
			return nil
		}
		for _, m := range values {
			if equalsIgnoreCase {
				if strings.EqualFold(m, valStr) {
					return nil
				}
			} else {
				if m == valStr {
					return nil
				}
			}
		}
		return errors.New("invalid value")
	}
}
