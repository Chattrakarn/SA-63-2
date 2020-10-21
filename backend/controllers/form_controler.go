package controllers

import (
	"context"
	
	"strconv"

	"github.com/User/app/ent"
	"github.com/User/app/ent/form"
	"github.com/gin-gonic/gin"
)

// FormController defines the struct for the form controller
type FormController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateForm handles POST requests for adding form entities
// @Summary Create form
// @Description Create form
// @ID create-form
// @Accept   json
// @Produce  json
// @Param form body ent.Form true "Form entity"
// @Success 200 {object} ent.Form
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /forms [post]
func (ctl *FormController) CreateForm(c *gin.Context) {
	obj := ent.Form{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Form binding failed",
		})
		return
	}

	f, err := ctl.client.Form.
		Create().
		SetFormType(obj.FormType).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, f)
}

// GetForm handles GET requests to retrieve a form entity
// @Summary Get a form entity by ID
// @Description get form by ID
// @ID get-form
// @Produce  json
// @Param id path int true "Form ID"
// @Success 200 {object} ent.Form
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /forms/{id} [get]
func (ctl *FormController) GetForm(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	f, err := ctl.client.Form.
		Query().
		Where(form.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, f)
}

// ListForm handles request to get a list of form entities
// @Summary List form entities
// @Description list form entities
// @ID list-form
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Form
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /forms [get]
func (ctl *FormController) ListForm(c *gin.Context) {
	limitQuery := c.Query("limit")
	limit := 10
	if limitQuery != "" {
		limit64, err := strconv.ParseInt(limitQuery, 10, 64)
		if err == nil {
			limit = int(limit64)
		}
	}

	offsetQuery := c.Query("offset")
	offset := 0
	if offsetQuery != "" {
		offset64, err := strconv.ParseInt(offsetQuery, 10, 64)
		if err == nil {
			offset = int(offset64)
		}
	}

	forms, err := ctl.client.Form.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, forms)
}

// // DeleteForm handles DELETE requests to delete a form entity
// // @Summary Delete a form entity by ID
// // @Description get form by ID
// // @ID delete-form
// // @Produce  json
// // @Param id path int true "Form ID"
// // @Success 200 {object} gin.H
// // @Failure 400 {object} gin.H
// // @Failure 404 {object} gin.H
// // @Failure 500 {object} gin.H
// // @Router /forms/{id} [delete]
// func (ctl *FormController) DeleteForm(c *gin.Context) {
// 	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
// 	if err != nil {
// 		c.JSON(400, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}

// 	err = ctl.client.Form.
// 		DeleteOneID(int(id)).
// 		Exec(context.Background())
// 	if err != nil {
// 		c.JSON(404, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}

// 	c.JSON(200, gin.H{"result": fmt.Sprintf("ok deleted %v", id)})
// }

// // UpdateForm handles PUT requests to update a form entity
// // @Summary Update a form entity by ID
// // @Description update form by ID
// // @ID update-form
// // @Accept   json
// // @Produce  json
// // @Param id path int true "Form ID"
// // @Param form body ent.Form true "Form entity"
// // @Success 200 {object} ent.Form
// // @Failure 400 {object} gin.H
// // @Failure 500 {object} gin.H
// // @Router /forms/{id} [put]
// func (ctl *FormController) UpdateForm(c *gin.Context) {
// 	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
// 	if err != nil {
// 		c.JSON(400, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}

// 	obj := ent.Form{}
// 	if err := c.ShouldBind(&obj); err != nil {
// 		c.JSON(400, gin.H{
// 			"error": "form binding failed",
// 		})
// 		return
// 	}
// 	obj.ID = int(id)
// 	f, err := ctl.client.Form.
// 		UpdateOne(&obj).
// 		Save(context.Background())
// 	if err != nil {
// 		c.JSON(400, gin.H{"error": "update failed"})
// 		return
// 	}

// 	c.JSON(200, f)
// }

// NewFormController creates and registers handles for the form controller
func NewFormController(router gin.IRouter, client *ent.Client) *FormController {
	uc := &FormController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
}

// InitFormController registers routes to the main engine
func (ctl *FormController) register() {
	forms := ctl.router.Group("/forms")

	forms.GET("", ctl.ListForm)

	// CRUD
	forms.POST("", ctl.CreateForm)
	forms.GET(":id", ctl.GetForm)
	// forms.PUT(":id", ctl.UpdateForm)
	// forms.DELETE(":id", ctl.DeleteForm)
}

