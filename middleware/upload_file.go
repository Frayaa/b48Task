package middleware

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func UploadFile(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		file, err := c.FormFile("image")

		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		src, err := file.Open()

		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		fmt.Println("src:", src)

		defer src.Close()

		tempFile, err := os.CreateTemp("uploads", "image-*.png")

		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		defer tempFile.Close()

		fmt.Println("tempfile:", tempFile)

		writtenCopy, err := io.Copy(tempFile, src)

		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		fmt.Println("writenncopy:", writtenCopy)

		data := tempFile.Name()
		fmt.Println("data file:", data)

		filename := data[8:]

		fmt.Println("filename kepotong:", filename)

		c.Set("dataFile", filename)

		return next(c)


	}
}