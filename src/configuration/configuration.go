package configuration

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/piLights/dioder"
)

type Configuration struct {
	BindTo              string
	ConfigurationFile   string
	Debug               bool
	IPv4Only            bool
	IPv6Only            bool
	NoAutoconfiguration bool
	Password            string
	PiBlaster           string
	Pins                dioder.Pins
	ServerName          string
	UpdateURL           string
	UseAvahi            bool

	DioderInstance dioder.Dioder `json:"-"`
}

func (config *Configuration) WriteConfigurationToFile(fileName string) error {

	serializedConfiguration, error := json.Marshal(config)
	if error != nil {
		return error
	}

	error = ioutil.WriteFile(fileName, serializedConfiguration, os.ModePerm)

	return error
}

func NewConfiguration(fileName string) (Configuration, error) {
	var DioderConfiguration Configuration
	if fileName == "" {
		DioderConfiguration = Configuration{}
		return DioderConfiguration, nil
	}

	content, error := ioutil.ReadFile(fileName)
	if error != nil {
		return Configuration{}, error
	}

	error = json.Unmarshal(content, &DioderConfiguration)

	return DioderConfiguration, error
}
