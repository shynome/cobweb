package main

import (
	"log"
	"os"
	"os/signal"
	"runtime"

	_ "github.com/GoAdminGroup/go-admin/adapter/echo"              // web framework adapter
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/sqlite" // sql driver
	_ "github.com/GoAdminGroup/themes/adminlte"                    // ui theme

	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/labstack/echo/v4"

	"github.com/shynome/cobweb/models"
	"github.com/shynome/cobweb/pages"
	"github.com/shynome/cobweb/tables"
	"github.com/shynome/cobweb/v2ray"
)

func main() {
	startServer()
}

func startServer() {
	e := echo.New()

	eng := engine.Default().AddConfigFromYAML("./config.yaml")
	models.Init(eng.SqliteConnection())

	orm := models.GetORM()
	var users []models.V2rayUser
	if err := orm.Model(models.V2rayUser{}).Find(&users).Error; err != nil {
		log.Fatal(err)
	}
	v2, err := v2ray.New(users)
	if err != nil {
		log.Fatal(err)
	}

	e.Use(v2ray.InjectV2rayMiddleware(v2))
	if err := eng.
		AddGenerators(tables.Generators).
		Use(e); err != nil {
		panic(err)
	}
	eng.HTML("GET", "/admin", pages.GetDashBoard)

	go func() { e.Logger.Fatal(e.Start(":3006")) }()
	go func() {
		if err := v2.Server.Start(); err != nil {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	// go-admin
	log.Print("closing database connection")
	eng.SqliteConnection().Close()
	// v2ray
	defer v2.Server.Close()
	runtime.GC()
}
