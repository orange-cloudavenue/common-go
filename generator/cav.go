package generator

import (
	"github.com/brianvoe/gofakeit/v7"

	"github.com/orange-cloudavenue/common-go/internal/regex"
)

func init() {
	// Register the URN generator function
	gofakeit.AddFuncLookup("edgegateway_name", gofakeit.Info{
		Category:    "cloudavenue",
		Display:     "EdgeGateway",
		Description: "Generate a new EdgeGateway name",
		Example:     "tn01e02ocb0001234spt101",
		Output:      "string",
		Generate: func(f *gofakeit.Faker, _ *gofakeit.MapParams, _ *gofakeit.Info) (any, error) {
			// Generate a random EdgeGateway name based on the regex pattern
			return f.Regex(regex.EdgeGatewayNameRegexString), nil
		},
	})

	gofakeit.AddFuncLookup("t0_name", gofakeit.Info{
		Category:    "cloudavenue",
		Display:     "T0Name",
		Description: "Generate a new T0 name",
		Example:     "pr01e02ocb0001234spt101",
		Output:      "string",
		Generate: func(f *gofakeit.Faker, _ *gofakeit.MapParams, _ *gofakeit.Info) (any, error) {
			// Generate a random T0 name based on the regex pattern
			return f.Regex(regex.T0NameRegexString), nil
		},
	})
}
