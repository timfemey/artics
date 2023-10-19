package postarticle

import (
	"artics-server/config"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"

	firestorepkg "cloud.google.com/go/firestore"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

var firestore = config.Firestore()

var auth = config.Auth()

var storage = config.Storage()

func NewArticle(fiber *fiber.Ctx) error {
	uuid := uuid.New()
	uid := fiber.FormValue("uid")
	article_name := fiber.FormValue("article-name")
	article := fiber.FormValue("article")
	banner, bannerFormErr := fiber.FormFile("banner")
	if bannerFormErr != nil {
		return fiber.Status(http.StatusForbidden).JSON(map[string]any{
			"status":  false,
			"message": "Error getting Banner file",
		})
	}
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
	content := banner.Header.Get("Content-Type")
	if content != "image/png" && content != "image/jpeg" && content != "image/svg" && content != "image/webp" && content != "image/jpg" {
		return fiber.Status(http.StatusForbidden).JSON(map[string]any{
			"status":  false,
			"message": "Image Format not Supported, Only PNG, JPEG, SVG and WEBP formats allowed",
		})
	}

	if banner.Size > int64(3*1024*1024) {
		return fiber.Status(http.StatusForbidden).JSON(map[string]any{
			"status":  false,
			"message": "Image should not be greater than 2MB",
		})
	}

	file, err := banner.Open()
	if err != nil {
		return fiber.Status(http.StatusInternalServerError).JSON(map[string]any{
			"status":  false,
			"message": "Could not Read File",
		})
	}
	fileBuf := bytes.NewBuffer(nil)
	_, errCopy := io.Copy(fileBuf, file)
	if errCopy != nil {
		return fiber.Status(http.StatusInternalServerError).JSON(map[string]any{
			"status":  false,
			"message": "Could not Read File Copy",
		})
	}
	wc := storage.Object(user.UID + "-" + time.Now().Format("20060102150405")).NewWriter(ctx)
	wc.ChunkSize = 0 // retries are not supported for chunk size 0.

	if _, err = io.Copy(wc, fileBuf); err != nil {
		return fiber.Status(http.StatusInternalServerError).JSON(map[string]any{
			"status":  false,
			"message": "Could not Read File Copy to Database",
		})
	}

	// Data can continue to be added to the file until the writer is closed.
	if err := wc.Close(); err != nil {
		return fiber.Status(http.StatusInternalServerError).JSON(map[string]any{
			"status":  false,
			"message": " Read File Error to Database",
		})
	}

	u := &url.URL{
		Scheme:   "https",
		Host:     "api.moderatecontent.com",
		Path:     "/text",
		Fragment: "none",
	}
	q := u.Query()
	q.Set("key", os.Getenv("API_KEY"))
	q.Add("url", wc.MediaLink)
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

	if reqBody["rating_letter"] != "e" {
		err := storage.Object(wc.Name).Delete(ctx)
		fmt.Println(err)

		return fiber.Status(http.StatusForbidden).JSON(map[string]any{
			"status":  false,
			"message": "Explicit Image Content not Allowed",
		})
	}

	_, err2 := firestore.Collection("article").Doc(uuid.String()).Set(ctx, map[string]any{
		"uid":          user.UID,
		"article_name": article_name,
		"article":      article,
		"id":           uuid.String(),
		"timestamp":    firestorepkg.ServerTimestamp,
		"image":        wc.MediaLink,
	})

	if err2 != nil {
		return fiber.Status(http.StatusFailedDependency).JSON(map[string]any{
			"status":  false,
			"message": "Failed to Upload Article",
		})
	}

	return fiber.Status(http.StatusOK).JSON(map[string]any{
		"status":  true,
		"message": "Article Uploaded Successfully",
	})
}
