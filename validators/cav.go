package validators

import (
	"github.com/go-playground/validator/v10"

	"github.com/orange-cloudavenue/common-go/internal/regex"
)

var (
	// CAVEdgeGatewayName is a validator that checks if a string is a valid CAV Edge Gateway name.
	CAVEdgeGatewayName = &CustomValidator{
		Key: "edgegateway_name",
		Func: func(fl validator.FieldLevel) bool {
			return regex.EdgeGatewayNameRegex().MatchString(fl.Field().String())
		},
	}

	CAVT0Name = &CustomValidator{
		Key: "t0_name",
		Func: func(fl validator.FieldLevel) bool {
			return regex.T0NameRegex().MatchString(fl.Field().String())
		},
	}
)
