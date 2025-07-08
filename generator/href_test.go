package generator

import (
	"testing"

	"github.com/orange-cloudavenue/common-go/internal/regex"
)

func TestGenerator_HREF_UUID(t *testing.T) {
	type CStruct struct {
		Href string `fake:"{href_uuid}"`
	}

	var st CStruct

	err := Struct(&st)
	if err != nil {
		t.Fatalf("Failed to create template: %v", err)
	}

	if !regex.UUID4Regex().MatchString(st.Href) {
		t.Fatalf("Expected HREF to have a valid UUID, got %s", st.Href)
	}
}
