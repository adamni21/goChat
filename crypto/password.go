package crypto

import (
	"crypto/rand"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

type PasswordHasher interface {
	Generate(password string) (string, error)
	Verify(password, hash string) (bool, error)
}

const Mega_Byte = 1024

type argonParams struct {
	time    uint32
	memory  uint32
	threads uint8
	keyLen  uint32
}

type argon2Hasher struct {
	*argonParams
	saltLen uint32
}

func NewArgon2Hasher() *argon2Hasher {
	return &argon2Hasher{
		argonParams: &argonParams{
			time:    10,
			memory:  64 * Mega_Byte,
			threads: 2,
			keyLen:  64,
		},
		saltLen: 16,
	}
}

func (a *argon2Hasher) Generate(password []byte) (string, error) {
	salt, err := generateRandomBytes(a.saltLen)
	if err != nil {
		return "", fmt.Errorf("generating random salt: %w", err)
	}

	hash, err := a.hash(password, salt)
	if err != nil {
		return "", fmt.Errorf("hashing password '%s' salt '%s': %w", password, salt, err)
	}

	return buildArgonString(hash, salt, a.time, a.memory, a.threads), nil
}

func (a *argon2Hasher) Verify(password, hash []byte) (bool, error) {
	salt, err := generateRandomBytes(16)
	if err != nil {
		return false, err
	}

	argon2.IDKey(password, salt, a.time, a.memory, a.threads, a.keyLen)

	return true, nil
}

func (p *argonParams) hash(password, salt []byte) ([]byte, error) {
	return argon2.IDKey(password, salt, p.time, p.memory, p.threads, p.keyLen), nil
}

func buildArgonString(hash, salt []byte, time, memory uint32, threads uint8) string {
	return fmt.Sprintf("$%s$v=%d$m=%d,t=%d,p=%d$%s$%s", "argon2id", argon2.Version, memory, time, threads, salt, hash)
}

func parseArgonString(argonString string) (params *argonParams, hash, salt []byte, err error) {
	parts := strings.Split(argonString, "$")

	hash = []byte(parts[5])
	salt = []byte(parts[4])

	var memory uint32
	var time uint32
	var threads uint8
	count, err := fmt.Sscanf(parts[3], "m=%d,t=%d,p=%d", &memory, &time, &threads)
	if err != nil {
		return nil, nil, nil, err
	}
	if count != 3 {
		return nil, nil, nil, fmt.Errorf("didn't parse all params from argonString: '%s'", argonString)
	}

	params = &argonParams{
		time:    time,
		memory:  memory,
		threads: threads,
		keyLen:  uint32(len(hash)),
	}

	return
}

func generateRandomBytes(n uint32) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}