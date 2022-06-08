package database

import (
	"github.com/rs/zerolog"
	"gitlab.com/mastocred/web-app/utility/environment"
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
