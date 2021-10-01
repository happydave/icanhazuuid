package uuid

import (
	"testing"
)

func TestUUID(t *testing.T) {
	uuid, err := GenerateUUID()
	if err != nil {
		t.Errorf("ERROR: GenerateUUID returned: %s", err.Error())
	}

	if uuid == "" {
		t.Errorf("ERROR: GenerateUUID invalid: %s", uuid)
	}

	t.Logf("INFO: Generated %s", uuid)
}
