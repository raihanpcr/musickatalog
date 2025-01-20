package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/raihanpcr/musickatalog/internal/configs"
	"github.com/raihanpcr/musickatalog/pkg/internalsql"
)

func main(){
	
	var (
		cfg *configs.Config
	)

	err := configs.Init(
		configs.WithConfigFolder(
			[]string{"./internal/configs/"},
		),
		configs.WithConfigFile("config"),
		configs.WithConfigType("yaml"),
	)
	if err != nil {
		log.Fatal("Gagal inisiasi config", err)
	}
	cfg = configs.Get()

	db, err := internalsql.Connect(cfg.Database.DataSourceName)
	if err != nil {
		log.Fatal("failed to connect to database, err: %+v", err)
	}

	r := gin.Default()

	r.Run(cfg.Service.Port)
}