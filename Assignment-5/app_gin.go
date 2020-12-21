package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"heyyy": "kloudfolks",
		})
	})

	// Routes

	router.GET("/get", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"method": "get",
		})
	})

	router.POST("/post", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"method": "post",
		})
	})

	router.PUT("/put", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"method": "put",
		})
	})

	router.DELETE("/delete", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"method": "delete",
		})
	})

	router.PATCH("/patch", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"method": "patch",
		})
	})

	// Parameters in Path
	router.GET("/heyyy/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(200, gin.H{
			"heyyy": name,
		})
	})

	router.GET("/hello/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		c.JSON(200, gin.H{
			"heyyy":  name,
			"action": action,
		})
	})

	// query string parameters
	router.GET("/thename", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Akash")
		lastname := c.DefaultQuery("lastname", "Verma")

		c.JSON(200, gin.H{
			"firstname": firstname,
			"lastname":  lastname,
		})
	})

	// encoded form
	router.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("message")
		name := c.DefaultPostForm("name", "anonymous")

		c.JSON(200, gin.H{
			"status":  "posted",
			"message": message,
			"name":    name,
		})
	})

	// json rendering
	router.GET("/someJSON", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	router.GET("/moreJSON", func(c *gin.Context) {
		var msg struct {
			Name    string `json:"user"`
			Message string
			Number  int
		}
		msg.Name = "vailhard"
		msg.Message = "hey"
		msg.Number = 432
		c.JSON(http.StatusOK, msg)
	})

	// HTML Rendering
	router.LoadHTMLGlob("templates/*")
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "HTMl Render",
		})
	})

	router.Run()
}
