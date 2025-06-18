package serialnumid_test

import (
	"testing"

	"github.com/raibru/goidgen/gen/serialnumid"
)

func Test_GenerateSerialNumberId_Successful(t *testing.T) {
	//t.Fatal("Check RED Test")

	tests := []struct {
		wantIdCounter      int
		wantStartNumberId  int
		expectCountedId    int
		expectSerialNumIds []string
	}{
		{1, 1, 1, []string{"1"}},
		{2, 1, 2, []string{"1", "2"}},
		{2, 5, 2, []string{"5", "6"}},
		{5, 10, 5, []string{"10", "11", "12", "13", "14"}},
	}

	for _, tt := range tests {
		idParam := serialnumid.GenerateParam{
			NumCount: tt.wantIdCounter,
			StartNum: tt.wantStartNumberId,
			UsingId:  tt.wantStartNumberId,
		}

		result, err := serialnumid.GenerateId(&idParam)
		if err != nil {
			t.Errorf("Test failed due err %v", err)
		}
		if len(result) != tt.expectCountedId {
			t.Errorf("Test failed want serial number count: %d but got: %d", 1, len(result))
		}
		if len(result) != len(tt.expectSerialNumIds) {
			t.Errorf("Test failed want expected serial number list: %v but got: %v", tt.expectSerialNumIds, result)
		}
		for i := range result {
			if tt.expectSerialNumIds[i] != result[i] {
				t.Errorf("Test failed want expected serial number list: %v but got: %v", tt.expectSerialNumIds, result)
			}
		}
	}
}
