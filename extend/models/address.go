package models

import (
	"github.com/dicteam/wallet-base/db"
)

// Address .
type Address struct {
	ID      uint   `gorm:"AUTO_INCREMENT" json:"id"`
	Address string `gorm:"index;unique_index:addr_ver" json:"address"`
	Type    uint   `gorm:"type:tinyint;index" json:"type"`
	Version string `gorm:"size:8;unique_index:addr_ver" json:"version"`
	PubKey  string `gorm:"size:512" json:"pubkey"`
	Status  uint   `gorm:"type:tinyint;default:0;index"`
}

// GetAllAddresses gets all addresses.
func GetAllAddresses() []*Address {
	var (
		addrs []*Address
	)
	db.Default().Find(&addrs)
	return addrs
}

func (a Address) TableName() string { return "address" }
