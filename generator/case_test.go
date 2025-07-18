package generator

import (
	"testing"

	"github.com/orange-cloudavenue/common-go/internal/regex"
)

func TestGenerator_Case(t *testing.T) {
	type CStruct struct {
		Camel  string `fake:"{camelCase}"`
		Pascal string `fake:"{PascalCase}"`
		Snake  string `fake:"{snake_case}"`
		Upper  string `fake:"{UPPER_CASE}"`
		Kebab  string `fake:"{kebab-case}"`
	}

	var st CStruct

	err := Struct(&st)
	if err != nil {
		t.Fatalf("Failed to create template: %v", err)
	}

	if !regex.CamelCaseRegex().MatchString(st.Camel) {
		t.Fatalf("Expected CamelCase name to match regex, got %s", st.Camel)
	}
	if !regex.PascalCaseRegex().MatchString(st.Pascal) {
		t.Fatalf("Expected PascalCase name to match regex, got %s", st.Pascal)
	}
	if !regex.SnakeCaseRegex().MatchString(st.Snake) {
		t.Fatalf("Expected snake_case name to match regex, got %s", st.Snake)
	}
	if !regex.UpperCaseRegex().MatchString(st.Upper) {
		t.Fatalf("Expected UPPER_CASE name to match regex, got %s", st.Upper)
	}
	if !regex.KebabCaseRegex().MatchString(st.Kebab) {
		t.Fatalf("Expected kebab-case name to match regex, got %s", st.Kebab)
	}
}
