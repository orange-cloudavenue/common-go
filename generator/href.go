package generator

import "github.com/brianvoe/gofakeit/v7"

func init() {
	gofakeit.AddFuncLookup("href_uuid", gofakeit.Info{
		Category: "custom",
		Display:  "href_uuid",
		Generate: func(f *gofakeit.Faker, _ *gofakeit.MapParams, _ *gofakeit.Info) (any, error) {
			return f.URL() + "/" + f.UUID(), nil
		},
	})
}
