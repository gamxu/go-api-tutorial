package controllers

import (
	"github/com/gamxu/go-api-tutorial/initializers"
	"github/com/gamxu/go-api-tutorial/models"

	"github.com/gin-gonic/gin"
)

func PostCreate(c *gin.Context) {
	// Get data from req body
	var body struct{
		Title string
		Body string
	}
	c.Bind(&body)

	// Create post
	post := models.Post{ Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)

	// Return on Error
	if result.Error != nil {
		c.Status(400)
		return
	}

	// Return on success
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostIndex(c *gin.Context){
	// Get post
	var posts []models.Post
	initializers.DB.Find(&posts) //find all post and assign to posts

	// Response with them
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostShow(c *gin.Context){
	// Get id from endpoint
	id := c.Param("id")

	// Get post
	var post models.Post
	initializers.DB.First(&post, id) //find post from id and assign to post

	// Response with them
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostUpdate(c *gin.Context){
	// Get id from endpoint
	id := c.Param("id")

	// Get data from req body
	var body struct{
		Title string
		Body string
	}
	c.Bind(&body)

	// Find post
	var post models.Post
	initializers.DB.First(&post, id) //find post from id and assign to post

	// Update post
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body: body.Body,
	})

	// Response
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostDelete(c *gin.Context){
	// Get data from id
	id := c.Param("id")

	// Find and delete data
	initializers.DB.Delete(&models.Post{}, id)

	// Response
	c.Status(200)
}