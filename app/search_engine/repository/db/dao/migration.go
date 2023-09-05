package dao

import (
	"os"

	"github.com/CocaineCong/tangseng/app/search_engine/repository/db/model"
	log "github.com/CocaineCong/tangseng/pkg/logger"
)

// 自动迁移模式
func migration() {
	err := _db.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(
			&model.InputData{},
		)
	if err != nil {
		log.LogrusObj.Infoln("register table fail")
		os.Exit(0)
	}
	log.LogrusObj.Infoln("register table success")
}
