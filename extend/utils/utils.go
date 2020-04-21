package utils

import (
	"fmt"

	"github.com/maiiz/coinlib/crypto"
	"github.com/maiiz/coinlib/utils"
)

func GenSequenceID(txid, to string, index int) string {
	return utils.BytesToHex(crypto.Hash160([]byte(txid + to + fmt.Sprintf("%d", index))))[:32]
}
