package validators

import (
	"github.com/go-playground/validator/v10"

	"github.com/orange-cloudavenue/common-go/internal/regex"
)

var (
	// ReExport ListCavResourceNames for external use.
	ListCavResourceNames = regex.ListCavResourceNames

	// CAVResourceName is a validator that checks if a string is a valid CAV resource name.
	// The list of valid resource names is defined in regex.ListCavResourceNames.
	// Usage: `validate:"resource_name=resource_key"`
	// E.g. `validate:"resource_name=edgegateway"`
	CAVResourceName = &CustomValidator{
		Key: "resource_name",
		Func: func(fl validator.FieldLevel) bool {
			for _, resource := range regex.ListCavResourceNames {
				if fl.Param() == resource.Key {
					return resource.RegexP.MatchString(fl.Field().String())
				}
			}
			return false
		},
	}
)
