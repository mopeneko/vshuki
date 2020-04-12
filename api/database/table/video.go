package table

import "github.com/jinzhu/gorm"

type Video struct {
	gorm.Model
	YouTubeID string
	Title     string
	Channel   Channel
	ChannelID uint
}
