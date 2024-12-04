package config

import "time"

type ENV string

const (
	EnvLocal       ENV = "local"
	EnvDevelopment ENV = "development"
	EnvStaging     ENV = "staging"
	EnvProduction  ENV = "production"
)

// config はアプリケーションの設定を表す構造体です。基本的には環境変数から読み込みます。
type config struct {
	Application struct {
		Name      string `env:"APP_NAME" envDefault:"submarine-backend"`
		Env       ENV    `env:"APP_ENV" envDefault:"local"`
		JWTSecret string `env:"JWT_SECRET"`
	}

	Server struct {
		Host            string        `env:"HOST" envDefault:"localhost"`
		Port            int           `env:"PORT" envDefault:"8080"`
		ShutdownTimeout time.Duration `env:"SHUTDOWN_TIMEOUT" envDefault:"10s"`
		AllowOrigin     string        `env:"APP_ALLOW_ORIGIN" envDefault:"http://localhost:3000"`
	}

	GitHub struct {
		ClientID     string `env:"GITHUB_CLIENT_ID"`
		ClientSecret string `env:"GITHUB_CLIENT_SECRET"`
		RedirectURI  string `env:"GITHUB_REDIRECT_URI"`
	}

	DB struct {
		User     string `env:"DB_USER"`
		Password string `env:"DB_PASSWORD"`
		Host     string `env:"DB_HOST"`
		Port     int    `env:"DB_PORT"`
		DBName   string `env:"DB_NAME"`
		SSLMode  string `env:"DB_SSLMODE" envDefault:"disable"`
	}
}

// Config は読み込まれた設定を保持します。
var Config *config
