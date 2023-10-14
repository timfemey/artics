package comment

import (
	"artics-server/config"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"os"

	firestorepkg "cloud.google.com/go/firestore"
	"github.com/gofiber/fiber/v2"
)

var firestore = config.Firestore()

var auth = config.Auth()

func Comment(fiber *fiber.Ctx) error {
	ctx := context.Background()
	docID := fiber.FormValue("article_id")
	uid := fiber.FormValue("uid")
	comment := fiber.FormValue("comment")

	if len(docID) < 3 && 3 > len(comment) {
		return fiber.Status(http.StatusForbidden).JSON(map[string]any{
			"status":  false,
			"message": "Article not found, Comment not sent",
		})
	}

	// url := fmt.Sprintf("https://api.moderatecontent.com/text/?exclude=bun,hell&replace=****&key=%[1]s&msg=%[2]s", os.Getenv("API_KEY"), comment)
	u := &url.URL{
		Scheme:   "https",
		Host:     "api.moderatecontent.com",
		Path:     "/text",
		Fragment: "none",
	}
	q := u.Query()
	q.Set("exclude", "bun,hell")
	q.Add("replace", "****")
	q.Add("key", os.Getenv("API_KEY"))
	q.Add("msg", comment)
	u.RawQuery = q.Encode()
	res, errReq := http.Get(u.String())
	if errReq != nil {
		return fiber.Status(http.StatusFailedDependency).JSON(map[string]any{
			"status":  false,
			"message": "Failed to Validate Comment",
		})
	}

	body, errBody := io.ReadAll(res.Body)
	if errBody != nil {
		return fiber.Status(http.StatusFailedDependency).JSON(map[string]any{
			"status":  false,
			"message": "Failed to Read Comment Validation",
		})
	}
	defer res.Body.Close()
	var reqBody map[string]any

	errReqBody := json.Unmarshal(body, &reqBody)
	if errReqBody != nil {

		return fiber.Status(http.StatusFailedDependency).JSON(map[string]any{
			"status":  false,
			"message": "Failed to Check Comment Body for Explicit Language",
		})
	}
	if len(reqBody["bad_words"].([]any)) > 0 {
		return fiber.Status(http.StatusForbidden).JSON(map[string]any{
			"status":  false,
			"message": "Explicit Comment Detected, Not Allowed!",
		})
	}
	_, errUser := auth.GetUser(ctx, uid)
	if errUser != nil {
		return fiber.Status(http.StatusForbidden).JSON(map[string]any{
			"status":  false,
			"message": "Invalid Account, Create an account before uploading a comment",
		})
	}

	_, _, err2 := firestore.Collection("article").Doc(docID).Collection("comments").Add(ctx, map[string]any{
		"uid":       uid,
		"comment":   comment,
		"timestamp": firestorepkg.ServerTimestamp,
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
