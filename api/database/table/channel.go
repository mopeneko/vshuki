package table

import "github.com/jinzhu/gorm"

type Channel struct {
	gorm.Model
	YouTubeID string
	Name      string
}
