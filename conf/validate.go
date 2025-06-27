package conf

import "github.com/meta-atman/gopkg/validation"

// validate validates the value if it implements the Validator interface.
func validate(v any) error {
	if val, ok := v.(validation.Validator); ok {
		return val.Validate()
	}

	return nil
}
