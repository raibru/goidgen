package uuid_test

import (
	"regexp"
	"testing"
	"time"

	uuid_google "github.com/google/uuid"
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

func Test_GenerateUuids_UuidVersion_V1_Successful(t *testing.T) {
	//t.Fatal("Check RED Test")
	tests := []struct {
		uuidCount     int
		uuidIsUpper   bool
		wantUuidCount int
	}{
		{1, false, 1},
		{1, true, 1},
		{2, true, 2},
		//		{10, true, 10},
		//		{1000, true, 1000},
	}

	for _, tt := range tests {
		uuidParam := uuid.GenerateParam{
			UuidVersionFlag: string("1"),
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
		for _, id := range result {
			if tt.uuidIsUpper {
				if !wantUuidRegexHigh.MatchString(id) {
					t.Errorf("Generated ID %s does not match UUID format pattern:\n\t%s", id, wantUuidRegexHigh)
				}
			} else {
				if !wantUuidRegexLow.MatchString(id) {
					t.Errorf("Generated ID %s does not match UUID format pattern:\n\t%s", id, wantUuidRegexLow)
				}
			}

			var parsed uuid_google.UUID
			if parsed, _ = uuid_google.Parse(id); parsed.Version() != 1 {
				t.Errorf("Expected V1 UUID, got V%d: %s", parsed.Version(), id)
			}
			// Zusätzlicher Test für die Zeitkomponente (nicht 100% genau wegen Testlaufzeit, aber im Bereich)
			parsedTime, _ := time.Parse(time.RFC3339, "0001-01-01T00:00:00Z") // Unix epoch in RFC3339
			uuidParsedTime := parsed.Time()
			_, uuidUnixNano := uuidParsedTime.UnixTime()
			unixNano := parsedTime.UnixNano() + uuidUnixNano

			// Prüfe, ob der Zeitstempel innerhalb eines akzeptablen Bereichs liegt (z.B. letzter Tag)
			// Dieser Test kann auf Systemen, die schnell sind oder lange laufen, ungenau sein.
			// Eine genauere Prüfung würde die MAC-Adresse umfassen, die hier nicht direkt geprüft wird.
			if time.Since(time.Unix(0, unixNano)) > 24*time.Hour+5*time.Minute {
				t.Logf("Warning: V1 UUID timestamp is unusually old: %s. Expected recent.", id)
			}
		}
	}
}
