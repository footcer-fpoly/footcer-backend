package upload

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/labstack/echo"
	uuid "github.com/satori/go.uuid"
)

func Upload(c echo.Context) ([]string, error) {

	images := make([]string, 0)

	// Multipart form
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
			path := "../images/" + folder + "/%s%s"
			fileName := uuid.NewV4()
			filePath := fmt.Sprintf(path, fileName, filepath.Ext(file.Filename))
			// Destination
			dst, err := os.Create(filePath)
			if err != nil {
				return nil, err

			}
			defer dst.Close()

			// Copy
			if _, err = io.Copy(dst, src); err != nil {
				return nil, err

			}

			url := "/" + folder + "/%s%s"
			url = fmt.Sprintf(url, fileName, filepath.Ext(file.Filename))
			images = append(images, url)
		}
	}
	return images, err

}
