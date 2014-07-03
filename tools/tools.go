/*
Meuh
Copyright (C) 2014 Thomas Silvi

This file is part of Meuh.

GoSimpleConfigLib is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 2 of the License, or
(at your option) any later version.

GoSimpleConfigLib is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with Foobar. If not, see <http://www.gnu.org/licenses/>.
*/

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
