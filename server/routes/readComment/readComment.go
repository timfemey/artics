package readcomment

import (
	"artics-server/config"
	"context"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

var firestore = config.Firestore()

var auth = config.Auth()

func ReadComment(fiber *fiber.Ctx) error {
	ctx := context.Background()
	articleId := fiber.FormValue("article_id")

	if len(articleId) < 4 {
		return fiber.Status(http.StatusForbidden).JSON(map[string]any{
			"message": "Invalid Article ID",
			"status":  false,
		})
	}

	docs, err := firestore.Collection("article").Doc(articleId).Collection("comments").Where("timestamp", ">=", time.Now().AddDate(0, 0, -21)).Limit(30).Documents(ctx).GetAll()
	if err != nil {
		return fiber.Status(http.StatusExpectationFailed).JSON(map[string]any{
			"message": "Failed to Fetch Article Comments from Server",
			"status":  false,
		})
	}

	var data []map[string]any
	for i := 0; i < len(docs); i++ {
		doc := docs[i].Data()
		user, err := auth.GetUser(ctx, doc["uid"].(string))
		if err != nil {
			return fiber.Status(http.StatusFailedDependency).JSON(map[string]any{
				"status":  false,
				"message": "Failed to Fetch Article Comments, Cant get Users",
			})
		}
		data = append(data, map[string]any{"username": user.DisplayName, "dp": user.PhotoURL, "comment": doc["comment"], "timestamp": doc["timestamp"]})
	}

	return fiber.Status(http.StatusOK).JSON(map[string]any{
		"status": true,
		"data":   data,
	})
}
