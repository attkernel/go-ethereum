package env

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/dicteam/wallet-base/db"
	boom "github.com/tylertreat/BoomFilters"
)

var Filter *boom.StableBloomFilter

func Init() {
	// initial db
	_, err := db.New("root:123456@tcp(127.0.0.1:3306)/testDb?charset=utf8&parseTime=True&loc=Local", "")
	if err != nil {
		panic(err)
	}
	fmt.Println("init DB successed!")

	//load compressed bits
	compressedBits, _ := ioutil.ReadFile("/data/result")
	var out bytes.Buffer
	newBuffer := bytes.NewBuffer(compressedBits)
	r, _ := zlib.NewReader(newBuffer)
	io.Copy(&out, r)

	// initial BoomFilters
	sbf := boom.NewDefaultStableBloomFilter(500000000, 0.01)
	_ = sbf.GobDecode(out.Bytes())
	Filter = sbf
	fmt.Println("init BoomFilters successed!")
}

func Default() *boom.StableBloomFilter {
	return Filter
}
