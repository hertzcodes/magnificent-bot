package entities

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	DiscordID       uint64
	DiscordUsername string
	RobloxID        uint64
	RobloxUsername  string
	Role            uint64
	IsPremium       bool
	// InventoryID uuid.UUID
}
