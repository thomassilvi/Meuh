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
