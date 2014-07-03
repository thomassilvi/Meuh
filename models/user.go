package models

import (
	"errors"
	"github.com/thomassilvi/Meuh/storage"
	"github.com/thomassilvi/Meuh/tools"
	"labix.org/v2/mgo/bson"
	"log"
	"time"
)

type BasicUser struct {
	Id    string
	Login string
	Email string
}

type UserPage struct {
	BasicPage
	User     BasicUser
	IsLogged bool
}

type UserDB struct {
	Id                bson.ObjectId `bson:"_id,omitempty" json:"-"`
	Login             string        `bson:"login" json:"login"`
	Email             string        `bson:"email" json:"email"`
	EncryptedPassword []byte        `bson:"encrypted_password" json:"encrypted_password"`
	Salt              []byte        `bson:"salt" json:"salt"`
	CreatedAt         time.Time     `bson:"created_at" json:"-"`
	UpdatedAt         time.Time     `bson:"updated_at" json:"-"`
}

var (
	ErrInvalidLoginPassword = errors.New("The username or password you entered is not valid")
	ErrUserNotFound         = errors.New("User not found")
	ErrUserSave             = errors.New("Error occured when saving user changes")
	ErrUserPasswordSave     = errors.New("Error occured when saving password modification")
)

//-------------------------------------------------------------------------------------------------

func (u *BasicUser) GetByLogin() error {
	db := storage.GetDB()
	c := db.C("users")
	utmp := UserDB{}
	err := c.Find(bson.M{"login": u.Login}).One(&utmp)

	if err != nil {
		if err.Error() == "not found" {
			return ErrUserNotFound
		}
		return err
	}

	u.Id = utmp.Id.Hex()
	u.Email = utmp.Email

	return err
}

//-------------------------------------------------------------------------------------------------

func (u *BasicUser) GetById() error {
	if u.Id == "" {
		return ErrUserNotFound
	}

	db := storage.GetDB()
	c := db.C("users")
	utmp := UserDB{}
	err := c.FindId(bson.ObjectIdHex(u.Id)).One(&utmp)

	if err != nil {
		if err.Error() == "not found" {
			return ErrUserNotFound
		}
		return err
	}

	u.Login = utmp.Login
	u.Email = utmp.Email

	return err
}

//-------------------------------------------------------------------------------------------------

func (u *BasicUser) Authentificate(password string) error {
	db := storage.GetDB()
	c := db.C("users")
	utmp := UserDB{}
	err := c.Find(bson.M{"login": u.Login}).One(&utmp)

	if err != nil {
		return ErrInvalidLoginPassword
	}

	if !tools.ComparePassword(password, utmp.EncryptedPassword, utmp.Salt) {
		return ErrInvalidLoginPassword
	}
	u.Id = utmp.Id.Hex()

	return err
}

//-------------------------------------------------------------------------------------------------

func (u *BasicUser) Save() error {
	db := storage.GetDB()
	c := db.C("users")

	change := bson.M{
		"$set": bson.M{"email": u.Email},
	}

	err := c.UpdateId(bson.ObjectIdHex(u.Id), change)

	if err != nil {
		log.Println("ERROR:User:Save:" + err.Error())
		return ErrUserSave
	}

	return nil
}

//-------------------------------------------------------------------------------------------------

func (u *BasicUser) ChangePassword(password string) error {
	db := storage.GetDB()
	c := db.C("users")

	encryptedPassword, salt := tools.CryptPassword(password)

	change := bson.M{
		"$set": bson.M{"salt": salt, "encrypted_password": encryptedPassword},
	}

	err := c.UpdateId(bson.ObjectIdHex(u.Id), change)

	if err != nil {
		log.Println("ERROR:User:ChangePassword:" + err.Error())
		return ErrUserPasswordSave
	}

	return nil
}

//-------------------------------------------------------------------------------------------------

func (u *BasicUser) Create(password string) error {
	db := storage.GetDB()
	utmp := UserDB{}
	utmp.Login = u.Login
	utmp.Email = u.Email
	utmp.EncryptedPassword, utmp.Salt = tools.CryptPassword(password)
	utmp.CreatedAt = time.Now()
	utmp.UpdatedAt = time.Now()
	c := db.C("users")
	err := c.Insert(&utmp)
	return err
}

//-------------------------------------------------------------------------------------------------
