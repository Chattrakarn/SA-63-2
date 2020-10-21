package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/User/app/ent"
	"github.com/User/app/ent/volume"
	"github.com/gin-gonic/gin"
)

// VolumeController defines the struct for the volume controller
type VolumeController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateVolume handles POST requests for adding volume entities
// @Summary Create volume
// @Description Create volume
// @ID create-volume
// @Accept   json
// @Produce  json
// @Param volume body ent.Volume true "Volume entity"
// @Success 200 {object} ent.Volume
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /volumes [post]
func (ctl *VolumeController) CreateVolume(c *gin.Context) {
	obj := ent.Volume{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "volume binding failed",
		})
		return
	}

	v, err := ctl.client.Volume.
		Create().
		SetVolumeType(obj.VolumeType).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, v)
}

// GetVolume handles GET requests to retrieve a volume entity
// @Summary Get a volume entity by ID
// @Description get volume by ID
// @ID get-volume
// @Produce  json
// @Param id path int true "Volume ID"
// @Success 200 {object} ent.Volume
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /volumes/{id} [get]
func (ctl *VolumeController) GetVolume(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	v, err := ctl.client.Volume.
		Query().
		Where(volume.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, v)
}

// ListVolume handles request to get a list of volume entities
// @Summary List volume entities
// @Description list volume entities
// @ID list-volume
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Volume
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /volumes [get]
func (ctl *VolumeController) ListVolume(c *gin.Context) {
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

	volumes, err := ctl.client.Volume.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, volumes)
}

// DeleteVolume handles DELETE requests to delete a volume entity
// @Summary Delete a volume entity by ID
// @Description get volume by ID
// @ID delete-volume
// @Produce  json
// @Param id path int true "Volume ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /volumes/{id} [delete]
func (ctl *VolumeController) DeleteVolume(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Volume.
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

// UpdateVolume handles PUT requests to update a volume entity
// @Summary Update a volume entity by ID
// @Description update volume by ID
// @ID update-volume
// @Accept   json
// @Produce  json
// @Param id path int true "Volume ID"
// @Param volume body ent.Volume true "Volume entity"
// @Success 200 {object} ent.Volume
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /volumes/{id} [put]
func (ctl *VolumeController) UpdateVolume(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Volume{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "volume binding failed",
		})
		return
	}
	obj.ID = int(id)
	v, err := ctl.client.Volume.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, v)
}

// NewVolumeController creates and registers handles for the volume controller
func NewVolumeController(router gin.IRouter, client *ent.Client) *VolumeController {
	uc := &VolumeController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
}

// InitVolumeController registers routes to the main engine
func (ctl *VolumeController) register() {
	volumes := ctl.router.Group("/volumes")

	volumes.GET("", ctl.ListVolume)

	// CRUD
	volumes.POST("", ctl.CreateVolume)
	volumes.GET(":id", ctl.GetVolume)
	volumes.PUT(":id", ctl.UpdateVolume)
	volumes.DELETE(":id", ctl.DeleteVolume)
}


