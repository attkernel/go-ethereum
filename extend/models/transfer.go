package models

import (
	"fmt"

	"github.com/dicteam/wallet-base/db"
	log "github.com/maiiz/coinlib/log"
	"github.com/shopspring/decimal"
)

type Transfer struct {
	SequnceID   string          `gorm:"primary_key;size:32" json:"sequnce_id"`
	Hash        string          `gorm:"column:txid;size:90;index" json:"hash"`
	AddressFrom string          `gorm:"size:256;index" json:"addressfrom"`
	AddressTo   string          `gorm:"size:256;index" json:"addressto"`
	Height      uint64          `gorm:"type:bigint;default:0;index" json:"height"`
	Amount      decimal.Decimal `gorm:"type:decimal(32,20);default:0" json:"amount"`
}

func (tx *Transfer) Insert() (err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("insert transaction %s failed, %v",
				tx.Hash, err)
			log.Error(err)
		}
	}()

	if TxExistedBySeqID(tx.SequnceID) {
		return nil
	}

	err = db.Default().FirstOrCreate(tx, "sequnce_id = ?", tx.SequnceID).Error
	if err != nil {
		return err
	}

	return nil
}

func TxExistedBySeqID(sequenceID string) bool {
	var tx Transfer
	db.Default().Where("sequnce_id = ?", sequenceID).First(&tx)
	if tx.SequnceID == "" {
		return false
	}
	return tx.SequnceID == sequenceID
}

func (tx Transfer) TableName() string { return "transfer" }
