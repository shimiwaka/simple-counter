package main

import (
	"time"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Access struct {
	gorm.Model
	Target	   string
	Date       time.Time
}

func main() {
  db, err := gorm.Open("mysql", "testuser:testpass@tcp(localhost:3309)/testdb?charset=utf8mb4&parseTime=True&loc=Local")

  if err != nil {
	fmt.Printf("err : %s", err)
    panic("failed to connect database")
  }

  db.AutoMigrate(&Access{})
  db.Exec("DELETE FROM accesses")
}