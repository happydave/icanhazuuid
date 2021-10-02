package uuid

import (
    "crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"
)

type UUID [16]byte

// Generate returns a new v4 UUID as string
func Generate() (string) {
    uuid := UUID{}

    // rand.Seed(time.Now().UnixNano())

    rb := make([]byte, 16)
    _, err := rand.Read(rb)

    if err != nil {
        return ""
    }

    for i:=0; i<len(uuid);i++ {
        uuid[i] = rb[i]
    }

    uuid[6] = (uuid[6] & 0x0f) | (4 << 4) // set version (4)
    uuid[8] = (uuid[8] & (0xff>>2) | (0x02 << 6)) // set variant (RFC 4122)

    return uuid.String()
}

// FromString creates a UUID from string
func FromString(s string) (UUID, error) {
	uuid := UUID{}

	// strip any '-'
	s = strings.ReplaceAll(s, "-", "")

	// Verify 32 char hex string
	if len(s) != 32 {
		return uuid, errors.New("Invalid length (" + fmt.Sprint(len(s)) + ")")
	}

	u, err := hex.DecodeString(s)

	if err != nil {
		return uuid, err
	}

    for i:=0; i<len(uuid);i++ {
        uuid[i] = u[i]
    }

	return uuid, nil
}

// Version returns UUID version
func (u UUID) Version() byte {
	return u[6] >> 4
}

// String returns string representation of UUID
func (u UUID) String() string {

	s := make([]byte, 36) // 32 + ('-' * 4)

	hex.Encode(s[0:8], u[0:4])
	s[8] = '-'
	hex.Encode(s[9:13], u[4:6])
	s[13] = '-'
	hex.Encode(s[14:18], u[6:8])
	s[18] = '-'
	hex.Encode(s[19:23], u[8:10])
	s[23] = '-'
	hex.Encode(s[24:36], u[10:16])

	return string(s)
}
