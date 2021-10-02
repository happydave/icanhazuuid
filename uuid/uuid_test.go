package uuid

import (
	"testing"
)

func TestGenerate(t *testing.T) {
	uuid := Generate()

	if uuid == "" {
		t.Errorf("ERROR: Generate invalid: %s", uuid)
	}

	uuid2 := Generate()

	if uuid == uuid2 {
		t.Errorf("ERROR: Random failed: %s %s", uuid, uuid2)
	}

	t.Logf("INFO: Generated %s", uuid)
	t.Logf("INFO: Generated %s", uuid2)
}

func TestFromString(t *testing.T) {
	_, err := FromString("00000000-0000-0000-0000-000000000000")
	if err != nil {
		t.Errorf("ERROR: UUID 1 FromString failed: %s", err)
	}

	_, err = FromString("1bfabfc2-9676-4a75-85fe-76df2dbcbdec")
	if err != nil {
		t.Errorf("ERROR: UUID 2 FromString failed: %s", err)
	}

	_, err = FromString("invalid")
	if err == nil {
		t.Errorf("ERROR: UUID 3 FromString failed to fail")
	}

	const s = "1bfabfc2-9676-4a75-85fe-76df2dbcbdez"
	_, err = FromString(s)
	if err == nil {
		t.Errorf("ERROR: UUID 4 FromString failed to fail")
	}
}

func TestVersion(t *testing.T) {
	u1 := UUID{}

	// 00000000-0000-0000-0000-000000000000
	v1 := u1.Version()
	if v1 != 0 {
		t.Errorf("ERROR: UUID 1 version (%d) expected (0)", v1)
	}

	// 1bfabfc2-9676-4a75-85fe-76df2dbcbdec
	u2, err := FromString("1bfabfc2-9676-4a75-85fe-76df2dbcbdec")
	v2 := u2.Version()

	if err != nil {
		t.Errorf("ERROR: UUID 2 version error: %s", err.Error())
	}

	if v2 != 4 {
		t.Errorf("ERROR: UUID 2 version (%d) expected (4)", v2)
	}
}

func TestString(t *testing.T) {
	u := UUID{}

	// 00000000-0000-0000-0000-000000000000
	s := u.String()
	if s != "00000000-0000-0000-0000-000000000000" {
		t.Errorf("ERROR: Empty UUID string expected (00000000-0000-0000-0000-0000000000) found (%s)", s)
	}

	// 1bfabfc2-9676-4a75-85fe-76df2dbcbdec
	u2, err := FromString("1bfabfc2-9676-4a75-85fe-76df2dbcbdec")

	if err != nil {
		t.Errorf("ERROR: UUID 2 version error: %s", err.Error())
	}

	if u2.String() != "1bfabfc2-9676-4a75-85fe-76df2dbcbdec" {
		t.Errorf("ERROR: UUID 2 String() expected (%s) returned (%s)", "1bfabfc2-9676-4a75-85fe-76df2dbcbdec", u2.String())
	}

	// t.Logf("INFO: String() on empty uuid: %s", s)
}
