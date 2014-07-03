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

package configuration

import (
	"github.com/thomassilvi/GoSimpleConfigLib"
	"log"
)

type Configuration struct {
	Web struct {
		Port                    uint16
		TemplatesPath           string
		PublicAssetsPath        string
		DisableUserRegistration bool
	}
	Database struct {
		Host     string
		Port     uint16
		Dbname   string
		User     string
		Password string
	}
	Cookie struct {
		StorePath   string
		StoreSecret string
	}
}

var AppConfiguration Configuration

//-------------------------------------------------------------------------------------------------

func InitConfiguration(filename string) (err error) {
	AppConfiguration = Configuration{}

	AppConfiguration.Web.DisableUserRegistration = false

	err = simple_config.ReadConfig(filename, &AppConfiguration)
	if err != nil {
		return err
	}

	if AppConfiguration.Web.TemplatesPath == "" {
		AppConfiguration.Web.TemplatesPath = "templates"
	}

	log.Println("INFO:Configuration:", AppConfiguration)

	return nil
}

//-------------------------------------------------------------------------------------------------
