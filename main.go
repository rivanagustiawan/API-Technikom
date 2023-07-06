package main

import (
	"log"
	"technikom/handler"
	"technikom/klien"
	"technikom/paket"
	"technikom/portofolio"
	"technikom/user"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	dsn := "root:root@tcp(127.0.0.1:8889)/apicms?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// db.AutoMigrate()

	// db.AutoMigrate(paket.Paket{}, klien.Klien{}, portofolio.Portofolio{}, user.User{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	paketRepository := paket.NewRepository(db)
	klienRepository := klien.NewRepository(db)
	portoRepository := portofolio.NewRepository(db)

	userServise := user.NewService(userRepository)
	paketService := paket.NewService(paketRepository)
	klienService := klien.NewService(klienRepository)
	portoService := portofolio.NewService(portoRepository)

	userHandler := handler.NewUserHandler(userServise)
	paketHandler := handler.NewPaketHandler(paketService)
	klienHandler := handler.NewKlienHandler(klienService)
	portofolioHandler := handler.NewPortoHandler(portoService)

	router := gin.Default()

	api := router.Group("/api/v1")
	api.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://google.com"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	api.GET("/users", userHandler.GetUser)

	api.GET("/paket", paketHandler.GetPakets)
	api.GET("/paket/:id", paketHandler.GetPaket)
	api.GET("/klien", klienHandler.GetKliens)
	api.GET("/klien/:id", klienHandler.GetKlien)
	api.GET("/portofolio", portofolioHandler.GetPortofolios)
	api.GET("/portofolio/:id", portofolioHandler.GetPortofolio)

	api.POST("/paket", paketHandler.CreatePaket)
	api.POST("/paket/deskripsi", paketHandler.CreateDeskPaket)
	api.POST("/klien", klienHandler.CreateKlien)
	api.POST("/portofolio", portofolioHandler.CreatePorto)

	api.PUT("/paket/:id", paketHandler.UpdatePaket)
	api.PUT("/klien/:id", klienHandler.UpdateKlien)
	api.PUT("/portofolio/:id", portofolioHandler.UpdatePortofolio)

	api.DELETE("/paket/:id", paketHandler.DeletePaket)
	api.DELETE("/paket/deskripsi/:id", paketHandler.DeleteDesk)
	api.DELETE("/klien/:id", klienHandler.DeleteKlien)
	api.DELETE("/portofolio/:id", portofolioHandler.DetelePortofolio)

	router.Run()
}
