package articlefeeds

import (
	"artics-server/config"
	"context"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

var firestore = config.Firestore()

var auth = config.Auth()

func Articlefeeds(fiber *fiber.Ctx) error {
	ctx := context.Background()
	docs, err := firestore.Collection("article").Where("timestamp", ">=", time.Now().AddDate(0, 0, -21)).Limit(40).Documents(ctx).GetAll()
	if err != nil {
		return fiber.Status(http.StatusFailedDependency).JSON(map[string]any{
			"status":  false,
			"message": "Failed to Fetch Article Feed",
		})
	}
	var data []map[string]any
	for i := 0; i < len(docs); i++ {
		doc := docs[i].Data()
		user, err := auth.GetUser(ctx, doc["uid"].(string))
		if err != nil {
			return fiber.Status(http.StatusFailedDependency).JSON(map[string]any{
				"status":  false,
				"message": "Failed to Fetch Article Feed, Cant get Users",
			})
		}
		data = append(data, map[string]any{"username": user.DisplayName, "dp": user.PhotoURL, "article_name": doc["article_name"], "article": doc["article"], "timestamp": doc["timestamp"]})
	}
	return fiber.Status(http.StatusOK).JSON(map[string]any{
		"status": true,
		"data":   data,
	})
}
