package controllers

import (
	"github.com/adit4git/go-crud/initializers"
	"github.com/adit4git/go-crud/models"
	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
//get data off request body

var body struct{
	Body string
	Title string
}

c.Bind(&body)
//create a post

post := models.Post{Title: body.Title, Body: body.Body}

result := initializers.DB.Create(&post) // pass pointer of data to Create



if result.Error != nil {
	c.Status(400)
	return 
}

//return it

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostIndex(c *gin.Context) {

	//get the post

	var posts []models.Post
	// Get all records
	initializers.DB.Find(&posts)
	// return them

	c.JSON(200, gin.H{
		"posts": posts,
	})

}

func PostShow(c *gin.Context) {

	//retrieve id from url
	id := c.Param("id")
	
	var post models.Post
	// Get specific post
	initializers.DB.First(&post,id)

	
	// return them

	c.JSON(200, gin.H{
		"post": post,
	})

}

func PostUpdate(c *gin.Context) {

	//retrieve id from url
	id := c.Param("id")
	//get post details from body of request

	var body struct{
		Body string
		Title string
	}
	
	c.Bind(&body)
	//create a post
	
	//post := models.Post{Title: body.Title, Body: body.Body}
	

	//find post from db
	var post models.Post
	// Get specific post
	initializers.DB.First(&post,id)

	// update it
	// Update attributes with `struct`, will only update non-zero fields
	initializers.DB.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Body,})	
    
	
	// return updated

	c.JSON(200, gin.H{
		"post": post,
	})

}

func PostDelete(c *gin.Context) {

	//retrieve id from url
	id := c.Param("id")
	
	
	// Get specific post
	initializers.DB.Delete(&models.Post{},id)

	// return them
	c.Status(200)

}