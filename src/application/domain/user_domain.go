package domain

import (
	"crypto/md5"
	"encoding/hex"
)

type UserDomain struct {
	Id       string
	Email    string
	Password string
	Name     string
	Age      int8
}

func (ud *UserDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.Password))
	ud.Password = hex.EncodeToString(hash.Sum(nil))
}

func (ud *UserDomain) GetId() string {
	return ud.Id
}

func (ud *UserDomain) GetEmail() string {
	return ud.Email
}

func (ud *UserDomain) GetPassword() string {
	return ud.Password
}

func (ud *UserDomain) GetName() string {
	return ud.Name
}

func (ud *UserDomain) GetAge() int8 {
	return ud.Age
}
