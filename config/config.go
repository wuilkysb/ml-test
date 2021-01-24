package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/labstack/gommon/log"
	"sync"
	"time"
)

type Config struct {
	ServerPort        string        `required:"true" split_words:"true"`
	ServerHost        string        `required:"true" split_words:"true"`
	Postfix           string        `required:"true"`
	PgHost            string        `required:"true" split_words:"true"`
	PgPort            int           `required:"true" split_words:"true"`
	PgName            string        `required:"true" split_words:"true"`
	PgUser            string        `required:"true" split_words:"true"`
	PgPassword        string        `required:"true" split_words:"true"`
	PgTimeout         time.Duration `required:"true" split_words:"true"`
	PgPoolSize        int           `required:"true" split_words:"true"`
	MigrationsCommand string        `required:"true" split_words:"true"`
}

var once sync.Once
var c Config

func Environments() Config {
	once.Do(func() {
		if err := envconfig.Process("", &c); err != nil {
			log.Errorf("Error parsing environment vars %#v", err)
		}
	})

	return c
}
