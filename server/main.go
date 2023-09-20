package main

import (
	"artics-server/routes/comment"
	postarticle "artics-server/routes/postArticle"
	readarticle "artics-server/routes/readArticle"
	"artics-server/routes/viewComments"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		Prefork: true,
	})

	app.Post("/newarticle", postarticle.NewArticle)

	app.Post("/comment", comment.Comment)

	app.Get("/article", readarticle.ReadArticle)

	app.Get("/viewcomms", viewComments.ViewComments)

	log.Fatal(app.Listen(":5000"))
}
