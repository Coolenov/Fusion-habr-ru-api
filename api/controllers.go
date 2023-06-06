package api

import (
	"Habr-ru-api/api/internal"
	"github.com/Coolenov/Fusion-library/lib"
	"github.com/gin-gonic/gin"
)

func ReturnAllPosts(c *gin.Context) {

	var posts []lib.Post

	posts = internal.HabrScraper()
	c.JSON(200, posts)
}
