package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/User/app/ent"
	"github.com/User/app/ent/unit"
	"github.com/gin-gonic/gin"
)

// UnitController defines the struct for the unit controller
type UnitController struct {
	client *ent.Client
	router gin.IRouter
}
  
// CreateUnit handles POST requests for adding unit entities
// @Summary Create unit
// @Description Create unit
// @ID create-unit
// @Accept   json
// @Produce  json
// @Param unit body ent.Unit true "Unit entity"
// @Success 200 {object} ent.Unit
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /units [post]
func (ctl *UnitController) CreateUnit(c *gin.Context) {
	obj := ent.Unit{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "unit binding failed",
		})
		return
	}

	un, err := ctl.client.Unit.
		Create().
		SetUnitType(obj.UnitType).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, un)
}

// GetUnit handles GET requests to retrieve a unit entity
// @Summary Get a unit entity by ID
// @Description get unit by ID
// @ID get-unit
// @Produce  json
// @Param id path int true "Unit ID"
// @Success 200 {object} ent.Unit
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /units/{id} [get]
func (ctl *UnitController) GetUnit(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	un, err := ctl.client.Unit.
		Query().
		Where(unit.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, un)
}

// ListUnit handles request to get a list of unit entities
// @Summary List unit entities
// @Description list unit entities
// @ID list-unit
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Unit
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /units [get]
func (ctl *UnitController) ListUnit(c *gin.Context) {
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

	units, err := ctl.client.Unit.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
 
	c.JSON(200, units)
}

// DeleteUnit handles DELETE requests to delete a unit entity
// @Summary Delete a unit entity by ID
// @Description get unit by ID
// @ID delete-unit
// @Produce  json
// @Param id path int true "Unit ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /units/{id} [delete]
func (ctl *UnitController) DeleteUnit(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Unit.
		DeleteOneID(int(id)).
		Exec(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"result": fmt.Sprintf("ok deleted %v", id)})
}

// UpdateUnit handles PUT requests to update a unit entity
// @Summary Update a unit entity by ID
// @Description update unit by ID
// @ID update-unit
// @Accept   json
// @Produce  json
// @Param id path int true "Unit ID"
// @Param unit body ent.Unit true "Unit entity"
// @Success 200 {object} ent.Unit
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /units/{id} [put]
func (ctl *UnitController) UpdateUnit(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Unit{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Unit binding failed",
		})
		return
	}
	obj.ID = int(id)
	un, err := ctl.client.Unit.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, un)
}

// NewUnitController creates and registers handles for the unit controller
func NewUnitController(router gin.IRouter, client *ent.Client) *UnitController {
	uc := &UnitController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
}

// InitUnitController registers routes to the main engine
func (ctl *UnitController) register() {
	units := ctl.router.Group("/units")

	units.GET("", ctl.ListUnit)

	// CRUD
	units.POST("", ctl.CreateUnit)
	units.GET(":id", ctl.GetUnit)
	units.PUT(":id", ctl.UpdateUnit)
	units.DELETE(":id", ctl.DeleteUnit)
}

