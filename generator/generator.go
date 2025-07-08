package generator

import (
	"github.com/brianvoe/gofakeit/v7"
)

// Reimporting gofakeit functions to the package level for access with generator package
var (
	Struct   = gofakeit.Struct
	Template = gofakeit.Template
)
