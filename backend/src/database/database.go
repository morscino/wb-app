package database

import (
	"github.com/MastoCred-Inc/web-app/utility/environment"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

const PackageName = "storage"

// Storage object
type Storage struct {
	Logger zerolog.Logger
	Env    *environment.Env
	DB     *gorm.DB
}

// Close securely closes the connection to the storage/database
func (d *Storage) Close() {
	sqlDD, _ := d.DB.DB()
	_ = sqlDD.Close()
}
