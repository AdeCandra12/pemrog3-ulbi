package main

import (
	"log"

	"github.com/AdeCandra12/pemrog3-ulbi/config"

	"github.com/aiteung/musik"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/whatsauth/whatsauth"

	"github.com/AdeCandra12/pemrog3-ulbi/url"

	"github.com/gofiber/fiber/v2"

	_ "github.com/AdeCandra12/pemrog3-ulbi/docs"
	// @title TES SWAG
	// @version 1.0
	// @description This is a sample swagger for Fiber
	// @contact.name API Support
	// @license.url http://github.com/AdeCandra12
	// @contact.email adecandra1500@gmail.com
	// @host pemrog3-ulbi.herokuapp.com
	// @BasePath /
	// @aschemes https http
)

func main() {
	go whatsauth.RunHub()
	site := fiber.New(config.Iteung)
	site.Use(cors.New(config.Cors))
	url.Web(site)
	log.Fatal(site.Listen(musik.Dangdut()))
}
