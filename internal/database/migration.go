package database

import (
	"github.com/anfelo/comments-api/internal/comment"
	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) error {
	if err := db.AutoMigrate(&comment.Comment{}); err != nil {
		return err
	}
	return nil
}
