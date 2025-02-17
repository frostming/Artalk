package db

import (
	"fmt"

	"github.com/ArtalkJS/Artalk/internal/config"
	"github.com/ArtalkJS/Artalk/internal/db/logger"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func NewDB(conf config.DBConf) (*gorm.DB, error) {
	gormConfig := &gorm.Config{
		Logger: logger.NewGormLogger(),
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: conf.TablePrefix,
		},
	}

	var dsn string
	if conf.Dsn != "" {
		dsn = conf.Dsn
	} else {
		dsn = getDsnByConf(conf)
	}

	switch conf.Type {
	case config.TypeSQLite:
		return OpenSQLite(dsn, gormConfig)
	case config.TypeMySql:
		return OpenMySql(dsn, gormConfig)
	case config.TypePostgreSQL:
		return OpenPostgreSQL(dsn, gormConfig)
	case config.TypeMSSQL:
		return OpenSqlServer(dsn, gormConfig)
	}

	return nil, fmt.Errorf("unsupported database type %s", conf.Type)
}

func NewTestDB() (*gorm.DB, error) {
	return OpenSQLite("file::memory:?cache=shared", &gorm.Config{})
}

func CloseDB(db *gorm.DB) error {
	ddb, err := db.DB()
	if err != nil {
		return err
	}
	return ddb.Close()
}
