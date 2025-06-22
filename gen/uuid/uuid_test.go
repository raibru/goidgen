package uuid_test

import (
	"regexp"
	"testing"

	"github.com/raibru/goidgen/gen/uuid"
)

var wantUuidRegexLow = regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`)
var wantUuidRegexHigh = regexp.MustCompile(`^[0-9A-F]{8}-[0-9A-F]{4}-[0-9A-F]{4}-[0-9A-F]{4}-[0-9A-F]{12}$`)

func Test_GenerateUuids_UuidVersion_V4_Successful(t *testing.T) {
	//t.Fatal("Check RED Test")
	tests := []struct {
		uuidCount     int
		uuidIsUpper   bool
		wantUuidCount int
	}{
		{1, false, 1},
		{1, true, 1},
		{2, true, 2},
		{10, true, 10},
		{1000, true, 1000},
	}

	for _, tt := range tests {
		uuidParam := uuid.GenerateParam{
			UuidVersionFlag: string("4"),
			NumUUIDs:        tt.uuidCount,
			ToUppercases:    tt.uuidIsUpper,
		}

		result, err := uuid.GenerateId(&uuidParam)
		if err != nil {
			t.Errorf("Expected no error for V1, got %v", err)
		}
		if len(result) != tt.wantUuidCount {
			t.Errorf("Expected uuid count: %d but got: %d", tt.uuidCount, len(result))
		}
		for _, uuid := range result {
			if tt.uuidIsUpper {
				if !wantUuidRegexHigh.MatchString(uuid) {
					t.Errorf("Generated ID %s does not match UUID format pattern:\n\t%s", uuid, wantUuidRegexHigh)
				}
			} else {
				if !wantUuidRegexLow.MatchString(uuid) {
					t.Errorf("Generated ID %s does not match UUID format pattern:\n\t%s", uuid, wantUuidRegexLow)
				}
			}
		}
	}
}

}
