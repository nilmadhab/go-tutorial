package codeheim

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type User struct {
	gorm.Model
	ID         uint64 `gorm:"primaryKey"`
	Username   string `gorm:"size:64"`
	Password   string `gorm:"size:255"`
	Notes      []Note
	CreditCard *CreditCard
}

type Note struct {
	gorm.Model
	ID      uint64 `gorm:"primaryKey"`
	Name    string `gorm:"size:255"`
	Content string `gorm:"type:text"`
	UserID  uint64 `gorm:"index"`
	User    user
}

type CreditCard struct {
	gorm.Model
	Number string
	UserID uint64
	User   User
}

var DB *gorm.DB

func connectDB() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,        // Disable color
		},
	)

	database, err := gorm.Open(mysql.Open("codeheim:tmp_pwd@tcp(127.0.0.1:3306)/gorm_belongs_to?charset=utf8&parseTime=true"), &gorm.Config{Logger: newLogger})

	if err != nil {
		panic("Failed to connect to databse!")
	}

	DB = database
}

func dbMigrate(db *gorm.DB) {
	DB.AutoMigrate(&User{}, &Note{}, &CreditCard{})
}

func main() {
	connectDB()
	dbMigrate()

}
