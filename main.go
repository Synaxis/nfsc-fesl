package main

import (
	"context"
	"flag"

	"github.com/google/gops/agent"
	"github.com/sirupsen/logrus"
	"github.com/subosito/gotenv"

	"github.com/Synaxis/nfsc-fesl/backend/config"
	"github.com/Synaxis/nfsc-fesl/backend/fesl"
	"github.com/Synaxis/nfsc-fesl/backend/matchmaking"
	"github.com/Synaxis/nfsc-fesl/backend/network"
	"github.com/Synaxis/nfsc-fesl/backend/theater"
)

func main() {
	initConfig()
	initLogger()

	ctx, cancel := context.WithCancel(context.Background())
	defer func() {
		cancel()
	}()
	startServer(ctx)

	// Use "github.com/google/gops/agent"
	if err := agent.Listen(agent.Options{}); err != nil {
		logrus.Fatal(err)
	}

	logrus.Println("==Plasma Online==")
	<-ctx.Done()
}

func initConfig() {
	var (
		configFile string
	)
	flag.StringVar(&configFile, "config", ".env", "Path to configuration file")
	flag.Parse()

	gotenv.Load(configFile)
	config.Initialize()
}

func initLogger() {
	logrus.SetLevel(config.LogLevel())

	logrus.SetFormatter(&logrus.JSONFormatter{
		DisableTimestamp: true,
		//DisableColors: false,
	})
	//  logrus.SetFormatter(new(prefixed.TextFormatter))
	//  logrus.SetFormatter(&prefixed.TextFormatter{
	//  	DisableTimestamp: true,
	//  	DisableColors:    false,
	//  })
}

func startServer(ctx context.Context) {

	network.InitClientData()
	mm := matchmaking.NewPool()

	fesl.New(config.FeslClientAddr(), false, mm).ListenAndServe(ctx)
	fesl.New(config.FeslServerAddr(), true, mm).ListenAndServe(ctx)

	theater.New(config.ThtrClientAddr(), mm).ListenAndServe(ctx)
	theater.New(config.ThtrServerAddr(), mm).ListenAndServe(ctx)
}
