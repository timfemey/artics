package readarticle

import (
	"artics-server/config"
	"context"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

var firestore = config.Firestore()

var auth = config.Auth()

func ReadArticle(fiber *fiber.Ctx) error {
	ctx := context.Background()

	articleId := fiber.FormValue("article_id")

	if len(articleId) < 4 {
		return fiber.Status(http.StatusForbidden).JSON(map[string]any{
			"message": "Invalid Article ID",
			"status":  false,
		})
	}

	doc, err := firestore.Collection("article").Doc(articleId).Get(ctx)
	if err != nil {
		return fiber.Status(http.StatusExpectationFailed).JSON(map[string]any{
			"message": "Failed to Fetch Article from Server",
			"status":  false,
		})
	}

	data := doc.Data()
	user, err := auth.GetUser(ctx, data["uid"].(string))

	return fiber.Status(http.StatusOK).JSON(map[string]any{
		"article":      data["article"],
		"username":     user.DisplayName,
		"dp":           user.PhotoURL,
		"article_name": data["article_name"],
		"status":       true,
	})

}
