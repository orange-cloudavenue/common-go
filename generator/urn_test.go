package generator

import (
	"testing"

	"github.com/orange-cloudavenue/common-go/urn"
)

func TestGenerator_URN(t *testing.T) {
	type CStruct struct {
		ID string `fake:"{urn:vm}"`
	}

	var st CStruct

	err := Struct(&st)
	if err != nil {
		t.Fatalf("Failed to create template: %v", err)
	}

	if !urn.IsVM(st.ID) {
		t.Fatalf("Expected URN to be of type VM, got %s", st.ID)
	}
}

func TestGenerator_URN_NoURNType(t *testing.T) {
	type CStruct struct {
		ID string `fake:"{urn}"`
	}

	var st CStruct

	err := Struct(&st)
	if err == nil {
		t.Fatal("Expected error when no URN type is provided, but got none")
	}
}
