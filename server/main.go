package main

import (
	articlefeeds "artics-server/routes/articleFeeds"
	"artics-server/routes/comment"
	postarticle "artics-server/routes/postArticle"
	readarticle "artics-server/routes/readArticle"
	readcomment "artics-server/routes/readComment"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

}

func main() {
	app := fiber.New(fiber.Config{
		Prefork: true,
	})

	app.Get("/articleFeed", articlefeeds.Articlefeeds)

	app.Post("/newarticle", postarticle.NewArticle)

	app.Post("/comment", comment.Comment)

	app.Post("/article", readarticle.ReadArticle)

	app.Post("/viewcomms", readcomment.ReadComment)

	log.Fatal(app.Listen(":5000"))
}
