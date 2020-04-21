package models

import (
	"fmt"
	"os"
	"testing"

	"github.com/dicteam/wallet-base/db"
)

func TestInitTable(t *testing.T) {
	dbInstance, err := db.New("root:123456@tcp(127.0.0.1:3306)/testDb?charset=utf8&parseTime=True&loc=Local", "")
	if err != nil {
		fmt.Println("Execute() error!")
		os.Exit(1)
	}
	defer dbInstance.Close()
	db.Default().AutoMigrate(
		&Transfer{},
		&Address{},
	)
}
