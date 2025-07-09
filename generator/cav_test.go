package generator

import (
	"testing"

	"github.com/orange-cloudavenue/common-go/internal/regex"
)

func TestGenerator_EdgeGatewayName(t *testing.T) {
	type CStruct struct {
		Name string `fake:"{edgegateway_name}"`
	}

	var st CStruct

	err := Struct(&st)
	if err != nil {
		t.Fatalf("Failed to create template: %v", err)
	}

	if !regex.EdgeGatewayNameRegex().MatchString(st.Name) {
		t.Fatalf("Expected Edge Gateway name to match regex, got %s", st.Name)
	}
}

func TestGenerator_T0Name(t *testing.T) {
	type CStruct struct {
		Name string `fake:"{t0_name}"`
	}

	var st CStruct

	err := Struct(&st)
	if err != nil {
		t.Fatalf("Failed to create template: %v", err)
	}

	if !regex.T0NameRegex().MatchString(st.Name) {
		t.Fatalf("Expected T0 name to match regex, got %s", st.Name)
	}
}
