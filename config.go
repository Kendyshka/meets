package untitled5

import (
	"errors"
	"gopkg.in/yaml.v3"
	"os"
	"untitled5/internal/store"
)

func ParseConfig(pathFile string) (config *Config, err error) {
	var rawData []byte
	if rawData, err = os.ReadFile(pathFile); err != nil {return}

	if err = yaml.Unmarshal(rawData, &config); err != nil {return}

	if err = config.Validate(); err != nil {return}
	return
}

type Config struct {
	LogLvl 	string 			`yaml:"log_lvl"`
	Store 	*store.Config 	`yaml:"store"`

}

func (c *Config) Validate() (err error) {
	if c == nil {return errors.New("config is empty")}
	if c.LogLvl == "" {return errors.New("LogLVL ошибка")}
	if err = c.Store.Validate(); err != nil {return}
	return
}