package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
)

// curl -X POST http://localhost:8080/upload \
//  -F "upload[]=@/Users/aaronmegs/test1.zip" \
//  -F "upload[]=@/Users/aaronmegs/test2.zip" \
//  -H "Content-Type: multipart/form-data"
func main() {
	router := gin.Default()
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.POST("/upload", func(c *gin.Context) {
		// Multipart form
		form, err := c.MultipartForm()
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("get multipart form err: %s", err.Error()))
			return
		}
		files := form.File["upload[]"]

		for _, file := range files {
			filename := filepath.Base(file.Filename)
			if err := c.SaveUploadedFile(file, filename); err != nil {
				c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
				return
			}
		}

		c.String(http.StatusOK, fmt.Sprintf("uploaded successfully %d files", len(files)))
	})

	router.Run(":8080")
}
