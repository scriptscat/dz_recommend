package synchronizer

import (
	"context"

	"github.com/codfrm/cago/config"
	cagoMysql "github.com/codfrm/cago/database/mysql"
	"github.com/go-mysql-org/go-mysql/canal"
	mysqlDriver "github.com/go-sql-driver/mysql"
)

func Synchronizer(ctx context.Context, cancel context.CancelFunc, config *config.Config) error {
	mysqlCfg := &cagoMysql.Config{}
	if err := config.Scan("mysql", mysqlCfg); err != nil {
		return err
	}
	cfg := canal.NewDefaultConfig()
	dsn, err := mysqlDriver.ParseDSN(mysqlCfg.Dsn)
	if err != nil {
		return err
	}
	cfg.Addr = dsn.Addr
	cfg.User = dsn.User
	cfg.Password = dsn.Passwd
	cfg.Dump.ExecutionPath = ""
	cfg.IncludeTableRegex = []string{mysqlCfg.Prefix + "forum_thread"}
	c, err := canal.NewCanal(cfg)
	if err != nil {
		return err
	}
	s := &syncToEs{
		ctx:    ctx,
		cancel: cancel,
		config: config,
		canal:  c,
	}
	return s.start()
}
