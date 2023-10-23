package pwhash

import (
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

const (
	bcryptRounds = 12
)

func Hash(plain []byte) ([]byte, error) {
	b, err := bcrypt.GenerateFromPassword(plain, bcryptRounds)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return b, nil
}

func Compare(hashed, plain []byte) error {
	if err := bcrypt.CompareHashAndPassword(hashed, plain); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
