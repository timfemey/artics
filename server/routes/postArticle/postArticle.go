package postarticle

import (
	"artics-server/config"
	"context"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

var firestore = config.Firestore()

var auth = config.Auth()

func NewArticle(fiber *fiber.Ctx) error {

	uid := fiber.FormValue("uid")
	article_name := fiber.FormValue("article-name")
	article := fiber.FormValue("article")
	ctx := context.Background()

	user, err := auth.GetUser(ctx, uid)
	if err != nil {
		return fiber.Status(http.StatusForbidden).JSON(map[string]any{
			"status":  false,
			"message": "User not found, Create an account before uploading an article",
		})
	}

	if len(article_name) > 100 {
		return fiber.Status(http.StatusForbidden).JSON(map[string]any{
			"status":  false,
			"message": "Article Name Length should not be more than 100 Characters",
		})
	}

	if len(article) > 7000 {
		return fiber.Status(http.StatusForbidden).JSON(map[string]any{
			"status":  false,
			"message": "Article Length should not be more than 7000 Characters",
		})
	}

	firestore.Collection("article").Add(ctx, map[string]any{
		"username":     user.DisplayName,
		"article_name": article_name,
		"article":      article,
	})

	return nil
}
