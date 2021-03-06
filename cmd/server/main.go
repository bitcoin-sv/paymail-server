package main

import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/bitcoin-sv/paymail/cmd/internal"
	"github.com/bitcoin-sv/paymail/config"
	"github.com/bitcoin-sv/paymail/config/databases"
	"github.com/bitcoin-sv/paymail/log"
)

const appname = "paymail-server"
const banner = `
=========================================================================================================================================

███████████                                                  ███  ████      █████████                                                   
░░███░░░░░███                                                ░░░  ░░███     ███░░░░░███                                                  
 ░███    ░███  ██████   █████ ████ █████████████    ██████   ████  ░███    ░███    ░░░   ██████  ████████  █████ █████  ██████  ████████ 
 ░██████████  ░░░░░███ ░░███ ░███ ░░███░░███░░███  ░░░░░███ ░░███  ░███    ░░█████████  ███░░███░░███░░███░░███ ░░███  ███░░███░░███░░███
 ░███░░░░░░    ███████  ░███ ░███  ░███ ░███ ░███   ███████  ░███  ░███     ░░░░░░░░███░███████  ░███ ░░░  ░███  ░███ ░███████  ░███ ░░░ 
 ░███         ███░░███  ░███ ░███  ░███ ░███ ░███  ███░░███  ░███  ░███     ███    ░███░███░░░   ░███      ░░███ ███  ░███░░░   ░███     
 █████       ░░████████ ░░███████  █████░███ █████░░████████ █████ █████   ░░█████████ ░░██████  █████      ░░█████   ░░██████  █████    
░░░░░         ░░░░░░░░   ░░░░░███ ░░░░░ ░░░ ░░░░░  ░░░░░░░░ ░░░░░ ░░░░░     ░░░░░░░░░   ░░░░░░  ░░░░░        ░░░░░     ░░░░░░  ░░░░░     
                         ███ ░███                                                                                                        
                        ░░██████                                                                                                         
                         ░░░░░░      																
						 
=========================================================================================================================================
`

// main is the entry point of the application.
// @title Payment Protocol Server
// @version 0.0.1
// @description Payment Protocol Server is an implementation of a Bip-270 payment flow.
// @termsOfService https://github.com/libsv/go-payment_protocol/blob/master/CODE_STANDARDS.md
// @license.name ISC
// @license.url https://github.com/libsv/go-payment_protocol/blob/master/LICENSE
// @host localhost:8445
// @schemes:
//	- http
//	- https
func main() {
	println("\033[32m" + banner + "\033[0m")
	config.SetupDefaults()
	cfg := config.NewViperConfig(appname).
		WithServer().
		WithDb().
		WithDeployment(appname).
		WithLog().
		WithPayD().
		WithDPP().
		WithPaymail().
		WithTransports().
		Load()
	log := log.NewZero(cfg.Logging)
	log.Infof("\n------Environment: %#v -----\n", cfg.Server)
	if err := cfg.Validate(); err != nil {
		log.Fatal(err, "config error")
	}

	if err := internal.GenerateCapabilitiesDocument(cfg.Paymail); err != nil {
		log.Fatal(err, "error generating capabilities document")
	}
	if err := internal.GenerateCapabilitiesDocumentV1(cfg.Paymail); err != nil {
		log.Fatal(err, "error generating capabilities document v1")
	}

	db, err := databases.NewDbSetup().SetupDb(log, cfg.Db)
	if err != nil {
		log.Fatal(err, "failed to setup database")
	}
	// nolint:errcheck // dont care about error.
	defer db.Close()

	e := internal.SetupEcho(log)

	if cfg.Server.SwaggerEnabled {
		internal.SetupSwagger(*cfg.Server, e)
	}

	// setup transports
	internal.SetupHTTPEndpoints(internal.SetupDeps(*cfg, log, db), e)

	if cfg.Deployment.IsDev() {
		internal.PrintDev(e)
	}
	go func() {
		log.Error(e.Start(cfg.Server.Port), "echo server failed")
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		log.Error(err, "")
	}

}
