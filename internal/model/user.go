package model

import "golang.org/x/crypto/bcrypt"

type User struct {
	ID                int     `json:"id"`
	Email             string  `json:"email"`
	Password          string  `json:"password,omitempty"`
	EncryptedPassword string  `json:"-"`
	Weight            float32 `json:"weight"`
}

func (u *User) BeforeCreate() error {
	if len(u.Password) > 0 {
		enc, err := encryptstring(u.Password)
		if err != nil {
			return err
		}
		u.EncryptedPassword = enc
	}
	return nil
}

func encryptstring(pass string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
