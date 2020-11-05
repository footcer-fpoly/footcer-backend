package upload

import (
	"cloud.google.com/go/storage"
	"github.com/labstack/echo"
	uuid "github.com/satori/go.uuid"
	"google.golang.org/api/option"
	"google.golang.org/appengine"
	"io"
	"net/url"
)

var (
	storageClient *storage.Client
)

func Upload(c echo.Context) ([]string, error) {
	bucket := "footcer" //your bucket name

	var err error

	ctx := appengine.NewContext(c.Request())

	storageClient, err = storage.NewClient(ctx, option.WithCredentialsFile("../../security/pro/key_bucket.json"))
	if err != nil {
		return nil, err
	}
	images := make([]string, 0)
	//	// Multipart form
	form, err := c.MultipartForm()
	if err != nil {
		return nil, err
	}
	folder := form.Value["folder"][0]
	if folder != "" {
		files := form.File["files"]
		for _, file := range files {
			// Source
			src, err := file.Open()
			if err != nil {
				return nil, err

			}
			defer src.Close()
			fileName := folder + "/" + uuid.NewV1().String()
			sw := storageClient.Bucket(bucket).Object(fileName).NewWriter(ctx)

			if _, err := io.Copy(sw, src); err != nil {

				return nil, err
			}

			if err := sw.Close(); err != nil {
				return nil, err
			}
			u, err := url.Parse("https://storage.googleapis.com/" + bucket + "/" + sw.Attrs().Name)
			images = append(images, u.String())
		}
	}
	return images, err
}

func UploadForKey(c echo.Context, keyParam string) ([]string, error) {
	bucket := "footcer" //your bucket name

	var err error

	ctx := appengine.NewContext(c.Request())

	storageClient, err = storage.NewClient(ctx, option.WithCredentialsFile("../../security/pro/key_bucket.json"))
	if err != nil {
		return nil, err
	}
	images := make([]string, 0)
	//	// Multipart form
	form, err := c.MultipartForm()
	if err != nil {
		return nil, err
	}
	folder := form.Value["folder"][0]
	if folder != "" {
		files := form.File[keyParam]
		if files != nil {
			for _, file := range files {
				// Source
				src, err := file.Open()
				if err != nil {
					return nil, err

				}
				defer src.Close()
				fileName := folder + "/" + uuid.NewV1().String()
				sw := storageClient.Bucket(bucket).Object(fileName).NewWriter(ctx)

				if _, err := io.Copy(sw, src); err != nil {

					return nil, err
				}

				if err := sw.Close(); err != nil {
					return nil, err
				}
				u, err := url.Parse("https://storage.googleapis.com/" + bucket + "/" + sw.Attrs().Name)
				images = append(images, u.String())
			}
		} else {
			images = append(images, "")

		}

	}
	return images, err
}
