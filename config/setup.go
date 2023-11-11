package config

import (
	app "github.com/hexcraft-biz/envmod-app"
	mysql "github.com/hexcraft-biz/envmod-mysql"
)

// ================================================================
// Config
// ================================================================
type Config struct {
	*app.App
	*mysql.Mysql
}

func Load() (*Config, error) {
	emApp, err := app.New()
	if err != nil {
		return nil, err
	}

	emMysql, err := mysql.New()
	if err != nil {
		return nil, err
	}

	return &Config{
		App:   emApp,
		Mysql: emMysql,
	}, nil
}
