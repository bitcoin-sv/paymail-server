package main

import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/libsv/p4-server/cmd/internal"
	"github.com/libsv/p4-server/config"
	"github.com/libsv/p4-server/log"
)

const appname = "payment-protocol-rest-server"
const banner = `
====================================================================
         _         _       _            _            _     
        /\ \      /\ \    /\ \        /\ \          _\ \   
       /  \ \    /  \ \   \_\ \      /  \ \        /\__ \  
      / /\ \ \  / /\ \ \  /\__ \    / /\ \ \      / /_ \_\ 
     / / /\ \_\/ / /\ \_\/ /_ \ \  / / /\ \ \    / / /\/_/ 
    / / /_/ / / / /_/ / / / /\ \ \/ / /  \ \_\  / / /      
   / / /__\/ / / /__\/ / / /  \/_/ / /    \/_/ / / /       
  / / /_____/ / /_____/ / /     / / /         / / / ____   
 / / /     / / /     / / /     / / /________ / /_/_/ ___/\ 
/ / /     / / /     /_/ /     / / /_________/_______/\__\/ 
\/_/      \/_/      \_\/      \/____________\_______\/     

====================================================================
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
		WithDeployment(appname).
		WithLog().
		WithPayD().
		WithSockets().
		WithTransports().
		Load()
	log := log.NewZero(cfg.Logging)
	log.Infof("\n------Environment: %#v -----\n", cfg.Server)
	if err := cfg.Validate(); err != nil {
		log.Fatal(err, "config error")
	}

	e := internal.SetupEcho(log)

	if cfg.Server.SwaggerEnabled {
		internal.SetupSwagger(*cfg.Server, e)
	}

	// setup transports
	switch cfg.Transports.Mode {
	case config.TransportModeHTTP:
		internal.SetupHTTPEndpoints(internal.SetupDeps(*cfg, log), e)
	case config.TransportModeSocket:
		s := internal.SetupSockets(*cfg.Sockets, e)
		internal.SetupSocketMetrics(s)
		defer s.Close()
	case config.TransportModeHybrid:
		s := internal.SetupHybrid(*cfg, log, e)
		internal.SetupSocketMetrics(s)
		defer s.Close()
	}
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
