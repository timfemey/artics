package comment

import (
	"artics-server/config"
	"context"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

var firestore = config.Firestore()

var auth = config.Auth()

func Comment(fiber *fiber.Ctx) error {
	ctx := context.Background()
	docID := fiber.FormValue("id")
	uid := fiber.FormValue("uid")
	comment := fiber.FormValue("comment")

	if len(docID) < 3 && 3 > len(comment) {
		return fiber.Status(http.StatusForbidden).JSON(map[string]any{
			"status":  false,
			"message": "Article not found, Comment not sent",
		})
	}

	_, err := auth.GetUser(ctx, uid)
	if err != nil {
		return fiber.Status(http.StatusForbidden).JSON(map[string]any{
			"status":  false,
			"message": "Invalid Account, Create an account before uploading a comment",
		})
	}

	_, _, err2 := firestore.Collection("article").Doc(docID).Collection("comments").Add(ctx, map[string]any{
		"uid":     uid,
		"comment": comment,
	})
	if err2 != nil {
		return fiber.Status(http.StatusFailedDependency).JSON(map[string]any{
			"status":  false,
			"message": "Failed to Post Comment",
		})
	}
	return fiber.Status(http.StatusOK).JSON(map[string]any{
		"status":  true,
		"message": "Successfully Uploaded Comment",
	})
}
