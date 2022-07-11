package main

import (
	"context"
	"log"

	"github.com/codfrm/cago"
	"github.com/codfrm/cago/config"
	"github.com/codfrm/cago/config/file"
	"github.com/codfrm/cago/database/mysql"
	"github.com/codfrm/cago/mux"
	"github.com/codfrm/cago/pkg/logger"
	"github.com/codfrm/cago/server"
	"github.com/scriptscat/dz_recommend/go-recommender/es"
	"github.com/scriptscat/dz_recommend/go-recommender/synchronizer"
)

func main() {
	ctx := context.Background()
	source, err := file.NewSource("config.yaml", file.Yaml())
	if err != nil {
		log.Fatal(err)
	}
	cfg, err := config.NewConfig("go-recommend", source)
	if err != nil {
		log.Fatalf("load config err: %v", err)
	}
	err = cago.New(ctx, cfg).
		Registry(cago.FuncComponent(logger.Logger)).
		Registry(cago.FuncComponent(mysql.Mysql)).
		Registry(cago.FuncComponent(es.Elasticsearch)).
		RegistryCancel(cago.FuncComponentCancel(synchronizer.Synchronizer)).
		RegistryCancel(server.Http(func(r *mux.RouterGroup) error {
			r.GET("/recommend", recommend())
			return nil
		})).
		Start()
	if err != nil {
		log.Fatalf("start err: %v", err)
		return
	}
}
