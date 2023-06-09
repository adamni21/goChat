package crypto

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"
)

func TestParseArgonString(t *testing.T) {
	hash := []byte{1, 2, 3, 4}
	salt := []byte{1, 2, 3, 4}

	params := argonParams{
		time:    10,
		memory:  64,
		threads: 2,
		keyLen:  uint32(len(hash)),
	}

	argonString := fmt.Sprintf("$argon2id$v=19$m=%d,t=%d,p=%d$%x$%x", params.memory, params.time, params.threads, salt, hash)

	parsedParams, parsedHash, parsedSalt, err := parseArgonString(argonString)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(params, *parsedParams) {
		t.Errorf("got %+v want %+v", *parsedParams, params)
	}
	if !bytes.Equal(hash, parsedHash) {
		t.Errorf("got %v want %v", parsedHash, hash)
	}
	if !bytes.Equal(salt, parsedSalt) {
		t.Errorf("got %v want %v", parsedSalt, salt)
	}
}

func TestArgonString(t *testing.T) {
	hash := []byte{1, 2, 3, 4}
	salt := []byte{1, 2, 3, 4}

	params := argonParams{
		time:    10,
		memory:  64,
		threads: 2,
		keyLen:  uint32(len(hash)),
	}

	argonString := params.buildArgonString(hash, salt)

	parsedParams, parsedHash, parsedSalt, err := parseArgonString(argonString)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(params, *parsedParams) {
		t.Errorf("got %+v want %+v", *parsedParams, params)
	}
	if !bytes.Equal(hash, parsedHash) {
		t.Errorf("got %v want %v", parsedHash, hash)
	}
	if !bytes.Equal(salt, parsedSalt) {
		t.Errorf("got %v want %v", parsedSalt, salt)
	}
}
