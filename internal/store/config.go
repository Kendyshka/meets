package store

import (
	"errors"
)

type Config struct {
	NameDB 		string 		`yaml:"nameDB"`
	Port 		int 		`yaml:"port"`
	Host 		string 		`yaml:"host"`
	User 		string 		`yaml:"user"`
	Password 	string 		`yaml:"password"`

}

func (c *Config) Validate() (err error) {
	if c == nil {return errors.New("config store is empty")}
	if c.NameDB == "" {return errors.New("Ошибка с названием БД")}
	if c.Port > 65535 && c.Port < 0 {return errors.New("Ошибка с номером порта")}
	if c.Host == "" {return errors.New("Ошибка с номером хоста")}
	if c.User == "" {return errors.New("Ошибка с названием пользователя")}
	if c.Password == "" {return errors.New("Ошибка с паролем пользователя")}
	return
}