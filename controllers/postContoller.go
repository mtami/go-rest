package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/mtami/go-crud/apiHelpers"
	"github.com/mtami/go-crud/models"
)

// GetPosts ... Get all posts
//
//	@Summary		Get all posts
//	@Description	get all posts
//	@Tags			Posts
//	@Success		200	{array}		models.Post
//	@Failure		404	{object}	object
//	@Router			/v1/posts [get]
func ListPost(c *gin.Context) {
	var post []models.Post
	err := models.GetAllPost(&post)
	if err != nil {
		apiHelpers.RespondJSON(c, 404, post)
	} else {
		apiHelpers.RespondJSON(c, 200, post)
	}
}

// CreatePost ... Create Post
//
//	@Summary		Create new post based on paramters
//	@Description	Create new post
//	@Tags			Posts
//	@Accept			json
//	@Param			post	body		models.Post	true	"Post Data"
//	@Success		200		{object}	object
//	@Failure		400,500	{object}	object
//	@Router			/v1/posts [post]
func AddNewPost(c *gin.Context) {
	var post models.Post
	c.BindJSON(&post)
	err := models.AddNewPost(&post)
	if err != nil {
		apiHelpers.RespondJSON(c, 404, post)
	} else {
		apiHelpers.RespondJSON(c, 200, post)
	}
}

// GetPostByID ... Get the post by id
//
//	@Summary		Get one post
//	@Description	get post by ID
//	@Tags			Posts
//	@Param			id		path		string	true	"Post ID"
//	@Success		200		{object}	models.Post
//	@Failure		400,404	{object}	object
//	@Router			/v1/posts/{id} [get]
func GetOnePost(c *gin.Context) {
	id := c.Params.ByName("id")
	var post models.Post
	err := models.GetOnePost(&post, id)
	if err != nil {
		apiHelpers.RespondJSON(c, 404, post)
	} else {
		apiHelpers.RespondJSON(c, 200, post)
	}
}

// UpdatePost ... Update Post
//
//	@Summary		Update post based on paramters
//	@Description	Update post
//	@Tags			Posts
//	@Accept			json
//	@Param			post	body		models.Post	true	"Post Data"
//	@Success		200		{object}	object
//	@Failure		400,500	{object}	object
//	@Router			/v1/posts [put]
func PutOnePost(c *gin.Context) {
	var post models.Post
	id := c.Params.ByName("id")
	err := models.GetOnePost(&post, id)
	if err != nil {
		apiHelpers.RespondJSON(c, 404, post)
	}
	c.BindJSON(&post)
	err = models.PutOnePost(&post, id)
	if err != nil {
		apiHelpers.RespondJSON(c, 404, post)
	} else {
		apiHelpers.RespondJSON(c, 200, post)
	}
}

// DeletePostByID ... Delete post by id
//
//	@Summary		Delete one post
//	@Description	Delete post by ID
//	@Tags			Posts
//	@Param			id	path string	true "Post ID"
//	@Success		200		{object}	models.Post
//	@Failure		400,404	{object}	object
//	@Router			/v1/posts/{id} [delete]
func DeletePost(c *gin.Context) {
	var post models.Post
	id := c.Params.ByName("id")
	err := models.DeletePost(&post, id)
	if err != nil {
		apiHelpers.RespondJSON(c, 404, post)
	} else {
		apiHelpers.RespondJSON(c, 200, post)
	}
}
