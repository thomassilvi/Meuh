package tools

import (
	"bytes"
	"code.google.com/p/go.crypto/pbkdf2"
	"crypto/sha256"
	"errors"
	"regexp"
	"time"
)

var EmailRegExp = regexp.MustCompile(`(?i)[A-Z0-9._%+-]+@(?:[A-Z0-9-]+\.)+[A-Z]{2,6}`)

//-------------------------------------------------------------------------------------------------

func CheckEmailValidity(e string) error {
	if !EmailRegExp.MatchString(e) {
		return errors.New("email badly formed")
	}
	return nil
}

//-------------------------------------------------------------------------------------------------

func CheckPasswordValidity(p string, pc string) error {
	// TODO
	if p != pc {
		return errors.New("password and password confirmation do not match")
	}
	if len(p) < 9 {
		return errors.New("password is too short")
	}
	return nil
}

//-------------------------------------------------------------------------------------------------

func GetSalt() []byte {
	// TODO : create random 200 bytes 
	tmpSalt := time.Now().Format(time.StampNano)
	result := []byte(tmpSalt)
	return result
}

//-------------------------------------------------------------------------------------------------

func CryptPassword(password string) ([]byte, []byte) {
	salt := GetSalt()
	pass := []byte(password)
	encrypted_password := pbkdf2.Key(pass, salt, 4096, sha256.Size, sha256.New)
	return encrypted_password, salt
}

//-------------------------------------------------------------------------------------------------

func ComparePassword(password string, encrypted_password []byte, salt []byte) bool {
	pass := []byte(password)
	new_encrypted_password := pbkdf2.Key(pass, salt, 4096, sha256.Size, sha256.New)
	return (bytes.Compare(new_encrypted_password, encrypted_password) == 0)
}

//-------------------------------------------------------------------------------------------------
