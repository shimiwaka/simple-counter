package main

import (
	"encoding/json"
	"fmt"
	_ "embed"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Settings struct {
	DB_username string
	DB_pass     string
	DB_host     string
	DB_port     int
	DB_name     string
}

//go:embed settings/settings.json
var settingsRawData []byte

func connectDB() *gorm.DB {
	s := Settings{}

	err := json.Unmarshal(settingsRawData, &s)
	if err != nil {
		return nil
	}

	connectQuery := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		s.DB_username, s.DB_pass, s.DB_host, s.DB_port, s.DB_name)

	db, err := gorm.Open("mysql", connectQuery)

	if err != nil {
		return nil
	}

	return db
}
