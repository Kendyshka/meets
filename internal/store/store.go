package store

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Store struct {
	db *gorm.DB
}


func ConnectDB(config *Config) (s *Store, err error) {
	var db *gorm.DB
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.User, config.Password, config.Host,config.Port, config.NameDB)
	if db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {return}

	s = &Store{db: db}
	return
}

func (s *Store) GetMeets() (meets []*Meet, err error) {
	var result *gorm.DB
	if result = s.db.Find(&meets); result.Error != nil {return}
	return
}

func (s *Store) GetMeetById(id int) (meet *Meet, err error) {
	var result *gorm.DB
	if result = s.db.First(&meet, id); result.Error != nil {return}
	return
}