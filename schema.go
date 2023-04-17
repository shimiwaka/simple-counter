package main

import (
	"time"
	"github.com/jinzhu/gorm"
)

type Access struct {
	gorm.Model
	Target	   string
	Date       time.Time
}