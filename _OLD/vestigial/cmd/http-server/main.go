package main

import (
	"fmt"
	"log"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/nch-bowstave/paymail"
	"github.com/nch-bowstave/paymail/config"
)

const (
	appname = "paymail"
	banner  = `
========================================================================================
  _____  ____  __    _ ____    __  ____    ____  ____    
 |     ||    \ \ \  //|    \  /  ||    \  |    ||    |   
 |    _||     \ \ \// |     \/   ||     \ |    ||    |_  
 |___|  |__|\__\/__/  |__/\__/|__||__|\__\|____||______| 
========================================================================================
`
)

func main() {

	cfg := config.NewViperConfig(appname).
		WithServer().
		WithDb().
		WithCapability().
		WithLog()
	if err := cfg.Validate(); err != nil {
		log.Fatalf("%s", err)
	}

	fmt.Println(banner)
	e := echo.New()
	// g := e.Group("")
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.RequestID())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	err := paymail.GenerateCapabilitiesDocument()
	if err != nil {
		panic(err)
	}

	// log.Printf("setting up %s db connection \n", cfg.Db.Type)
	// db, err := databases.NewDbSetup().SetupDb(cfg.Db)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer func() {
	// 	_ = db.Close()
	// }()
	// log.Println("db connection setup")

	// paymailStore := sql.NewPaymailDb(db, cfg.Db.Type)
	// paymailService := service.NewPaymailService(paymailStore, cfg.Paymail.Domain)

	// web.NewAccount(paymailService).RegisterRoutes(g)
	// web.NewBsvAlias(paymailService).RegisterRoutes(g)
	// web.NewPKI(paymailService).RegisterRoutes(g)

	e.Logger.Fatal(e.Start(cfg.Server.Port))
}
