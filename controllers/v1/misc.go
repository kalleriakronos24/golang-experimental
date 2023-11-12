package v1

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kalleriakronos24/mygoapp2nd/dto"
)

func Pong(c *gin.Context) {
	c.JSON(http.StatusOK, dto.Response{Data: "pong"})
}

func UploadFileMultiple(c *gin.Context) {
	// Multipart form
	form, _ := c.MultipartForm()
	files := form.File["upload[]"]

	for _, file := range files {
		log.Println(file.Filename)
		// Upload the file to specific dst.
		c.SaveUploadedFile(file, ".")
	}
	c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
}

func UploadFileSingle(c *gin.Context) {
	// Multipart form
	file, _ := c.FormFile("file")
	log.Println(file.Filename)

	// Upload the file to specific dst.
	c.SaveUploadedFile(file, ".")

	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}
