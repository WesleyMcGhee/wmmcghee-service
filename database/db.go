package database

import (
  "gorm.io/gorm"
  "gorm.io/driver/sqlite"
)

func DB() *gorm.DB {
  db, err := gorm.Open(sqlite.Open("site.db"), &gorm.Config{})
  if err != nil {
    panic("failed to connect database")
  }

  return db
}
