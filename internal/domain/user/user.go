package user

import (
	"crypto/md5"
	"encoding/hex"
	"errors"

	"github.com/google/uuid"
)

type StoredUser struct {
	Username string    `json:"username"`
	Password string    `json:"password"`
	ID       uuid.UUID `json:"uuid"`
}

type User struct {
	id           uuid.UUID
	username     string
	passwordHash string
}

func NewUser(id uuid.UUID, username string, password string) *User {
	return &User{
		id:           id,
		username:     username,
		passwordHash: password,
	}
}

func CreateUser(username string, password string) (*User, error) {
	hasher := md5.New()
	hasher.Write([]byte(password))
	hash := hex.EncodeToString(hasher.Sum(nil))

	if username == "" || password == "" {
		return nil, errors.New("empty name or password")
	}

	return NewUser(uuid.New(), username, hash), nil
}

func (u *User) ID() uuid.UUID {
	return u.id
}

func (u *User) Username() string {
	return u.username
}

func (u *User) PasswordHash() string {
	return u.passwordHash
}
