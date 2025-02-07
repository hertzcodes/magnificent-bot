package domain

import (
	"time"
)

type (
	UserID    uint64
	DiscordID uint64
	RobloxID  uint64
	Role      uint64
)

const (
	NormalUser Role = iota
	ValueChanger
	Admin
	Owner
)

type User struct {
	ID              UserID
	CreatedAt       time.Time
	DeletedAt       time.Time
	DiscordID       DiscordID
	DiscordUsername string
	RobloxID        RobloxID
	RobloxUsername  string
	Role            Role
	IsPremium       bool
	// InventoryID     uuid.UUID
}
