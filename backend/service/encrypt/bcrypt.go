package encrypt

import (
	"backend/service"

	"golang.org/x/crypto/bcrypt"
)

type Encrypt struct {
	cost int
}

func (e *Encrypt) Generate(password string) (string, error) {
	pass, err := bcrypt.GenerateFromPassword([]byte(password), e.cost)
	if err != nil {
		return "", err
	}
	return string(pass), nil
}

func (e *Encrypt) Compare(hash string, password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false, err
	}
	return true, nil
}

func New() service.Encrypt {
	return &Encrypt{
		cost: cost,
	}
}
