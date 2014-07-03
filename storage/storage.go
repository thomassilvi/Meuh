package storage

import (
	"github.com/gorilla/sessions"
	"github.com/thomassilvi/Meuh/configuration"
	"labix.org/v2/mgo"
)

var StoreFS *sessions.FilesystemStore

var MDB *mgo.Database = nil

//-------------------------------------------------------------------------------------------------

func InitStorage() (err error) {
	StoreFS = sessions.NewFilesystemStore(configuration.AppConfiguration.Cookie.StorePath,
		[]byte(configuration.AppConfiguration.Cookie.StoreSecret))

	session, err := mgo.Dial("localhost")
	if err != nil {
		return err
	}
	session.SetMode(mgo.Monotonic, true)
	MDB = session.DB("mce")

	return nil
}

//-------------------------------------------------------------------------------------------------

func GetDB() *mgo.Database {
	return MDB
}

//-------------------------------------------------------------------------------------------------
