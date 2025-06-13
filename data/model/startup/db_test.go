package startup

import (
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func Test_StartupOnChain(t *testing.T) {
	db, err := gorm.Open(mysql.Open("root:st123456@tcp(127.0.0.1:3306)/metaland?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}

	db.Debug().Where(Startup{
		TxHash:  "11",
		ChainID: 1,
		ComerID: 1,
	}).Where("on_chain", false).Updates(Startup{
		OnChain: true,
	})
}
