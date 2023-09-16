package storage

import (
	"fmt"
	"testing"

	bolt "go.etcd.io/bbolt"

	"github.com/CocaineCong/tangseng/app/search_engine/analyzer"
	"github.com/CocaineCong/tangseng/config"
	log "github.com/CocaineCong/tangseng/pkg/logger"
)

func TestMain(m *testing.M) {
	// 这个文件相对于config.yaml的位置
	re := config.ConfigReader{FileName: "../../../../config/config.yaml"}
	config.InitConfigForTest(&re)
	log.InitLog()
	analyzer.InitSeg()
	fmt.Println("Write tests on values: ", config.Conf)
	m.Run()
}

func TestForwardDBRead(t *testing.T) {
	a := config.Conf.SeConfig.StoragePath + "1.forward"
	forward, err := NewForwardDB(a)
	if err != nil {
		fmt.Println("err", err)
	}
	count, err := forward.ForwardCount()
	if err != nil {
		fmt.Println("Err", err)
	}
	fmt.Println(count)
	r, err := forward.GetForward(1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(r))
}

func TestBoltDB(t *testing.T) {
	dbName := config.Conf.SeConfig.StoragePath + "0.forward"
	fmt.Println(dbName)
	db, err := bolt.Open(dbName, 0600, nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(db)
}
