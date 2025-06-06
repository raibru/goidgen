package uuid_test

import (
	"regexp"
	"testing"

	"github.com/raibru/goidgen/gen/uuid"
)

var wantUuidRegexLow = regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`)
var wantUuidRegexHigh = regexp.MustCompile(`^[0-9A-F]{8}-[0-9A-F]{4}-[0-9A-F]{4}-[0-9A-F]{4}-[0-9A-F]{12}$`)

func Test_GenerateUuids_UuidVersion_4_Successful(t *testing.T) {
	//t.Fatal("Check RED Test")
	tests := []struct {
		uuidCount     int
		uuidIsUpper   bool
		wantUuidCount int
	}{
		{1, false, 1},
		{1, true, 1},
	}

	for _, tt := range tests {
		uuidParam := uuid.GenerateParam{
			UuidVersionFlag: string("4"),
			NumUUIDs:        tt.uuidCount,
			ToUppercases:    tt.uuidIsUpper,
		}

		result, err := uuid.GenerateId(&uuidParam)
		if err != nil {
			t.Errorf("Test failed due err %v", err)
		}
		if len(result) != tt.wantUuidCount {
			t.Errorf("Test failed want uuid count: %d but got: %d", tt.uuidCount, len(result))
		}
		for _, uuid := range result {
			if tt.uuidIsUpper {
				if !wantUuidRegexHigh.MatchString(uuid) {
					t.Errorf("Test failed want uuid pattern: %s\n\tbut got: %s", wantUuidRegexHigh, uuid)
				}
			} else {
				if !wantUuidRegexLow.MatchString(uuid) {
					t.Errorf("Test failed want uuid pattern: %s\n\tbut got: %s", wantUuidRegexHigh, uuid)
				}
			}
		}

	}

	//for _, tt := range tests {
	//	t.Logf("Test TC: [%s]", tt.input)
	//	data, err := hex.DecodeString(tt.input)
	//	if err != nil {
	//		t.Errorf("can not decode hex string [%s]=%v", tt.input, err)
	//	}
	//	tc := MakeTelecommand()
	//	if err := tc.Parse(data); err != nil {
	//		t.Errorf("can not parse to TC packet [%s]=%v", tt.input, err)
	//	}

	//	if err := tc.Validate(); err != nil {
	//		t.Errorf("parsed PUS primary header is not valid: %s", err)
	//	}
	//	if packet.ToPacketErrorControlHexString(tc.PacketErrorControl) != tt.expectedCrc {
	//		t.Errorf("expect parsed packet error control %s but got %04x", tt.expectedCrc, tc.PacketErrorControl)
	//	}
	//	//t.Logf("\n%s\n", tc.Dump())
	//}
}
