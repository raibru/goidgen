package blockid_test

import (
	"regexp"
	"testing"

	"github.com/raibru/goidgen/gen/blockid"
)

var wantBlockidRegexLow = regexp.MustCompile(`^[a-zA-Z0-9-_]*[0-9a-f]{1,16}[a-zA-Z0-9-_]*$`)
var wantBlockidRegexHigh = regexp.MustCompile(`^[a-zA-Z0-9-_]*[0-9A-F]{1,16}[a-zA-Z0-9-_]*$`)

func Test_GenerateBlockids_UpperHexPart_Successful(t *testing.T) {
	//t.Fatal("Check RED Test")
	tests := []struct {
		wantByteCount   int
		wantPrefixName  string
		wantPostfixName string
	}{
		{1, "", ""},
		// {2, "", ""},
		// {4, "", ""},
		// {8, "", ""},
		// {16, "", ""},
	}

	for _, tt := range tests {
		idParam := blockid.GenerateParam{
			NumByteCount: tt.wantByteCount,
			PrefixName:   tt.wantPrefixName,
			PostfixName:  tt.wantPostfixName,
			NumBlockids:  1,
		}

		result, err := blockid.GenerateId(&idParam)
		if err != nil {
			t.Errorf("Test failed due err %v", err)
		}
		if len(result) != 1 {
			t.Errorf("Test failed want blockid count: %d but got: %d", 1, len(result))
		}
		for _, blockid := range result {
			if !wantBlockidRegexHigh.MatchString(blockid) {
				t.Errorf("Test failed want blockid pattern: %s\n\tbut got: %s", wantBlockidRegexHigh, blockid)
			}
		}
	}

}
