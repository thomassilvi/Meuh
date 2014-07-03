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
