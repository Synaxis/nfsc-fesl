package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
)

var (
	General  cfg
	Database MySQL
)

type cfg struct {
	LogLevel string `envconfig:"LOG_LEVEL" default:"INFO"`

	GameSpyIP      string `envconfig:"GAMESPY_IP" default:"127.0.0.1"`
	FeslClientPort int    `envconfig:"FESL_CLIENT_PORT" default:"18270"`
	FeslServerPort int    `envconfig:"FESL_SERVER_PORT" default:"18051"`
	ThtrClientPort int    `envconfig:"THEATER_CLIENT_PORT" default:"18210"`
	ThtrServerPort int    `envconfig:"THEATER_SERVER_PORT" default:"18056"`

	// An address for clients where theater can be found.
	// Preferable: domain name without protocol scheme and port
	ThtrAddr string `envconfig:"THEATER_ADDR" default:"127.0.0.1"`

	// Telemetry
	TelemetryIP    string `envconfig:"TELEMETRY_IP" default:"127.0.0.1"`
	TelemetryPort  int    `envconfig:"TELEMETRY_PORT" default:"13505"`
	TelemetryToken string `envconfig:"TELEMETRY_TOKEN"`
}

type MySQL struct {
	DatabaseHost         string `envconfig:"DATABASE_HOST" default:"127.0.0.1"`
	DatabasePort         int    `envconfig:"DATABASE_PORT" default:"3306"`
	DatabaseUserName     string `envconfig:"DATABASE_USERNAME" default:"root"`
	DatabasePassword     string `envconfig:"DATABASE_PASSWORD" default:"test"`
	DatabaseName         string `envconfig:"DATABASE_NAME" default:"dev"`
	DatabaseMaxIdleConns int    `envconfig:"DATABASE_MAX_IDLE_CONNS" default:"10"`
	DatabaseMaxOpenConns int    `envconfig:"DATABASE_MAX_OPEN_CONNS" default:"40"`
}

func Initialize() {
	if err := envconfig.Process("", &General); err != nil {
		logrus.WithError(err).Fatal("config: Initialize values for General")
	}
	if err := envconfig.Process("", &Database); err != nil {
		logrus.WithError(err).Fatal("config: Initialize values for Database")
	}
}

// LogLevel parses a default log level from a string
func LogLevel() logrus.Level {
	lvl, err := logrus.ParseLevel(General.LogLevel)
	if err != nil {
		logrus.WithError(err).Fatal("config: Parse log level")
	}
	return lvl
}

func bindAddr(addr string, port int) string {
	return fmt.Sprintf("%s:%d", addr, port)
}

func FeslClientAddr() string {
	return bindAddr(General.GameSpyIP, General.FeslClientPort)
}

func FeslServerAddr() string {
	return bindAddr(General.GameSpyIP, General.FeslServerPort)
}

func ThtrClientAddr() string {
	return bindAddr(General.GameSpyIP, General.ThtrClientPort)
}

func ThtrServerAddr() string {
	return bindAddr(General.GameSpyIP, General.ThtrServerPort)
}
