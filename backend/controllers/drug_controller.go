package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/User/app/ent"
	"github.com/User/app/ent/drug"
	"github.com/User/app/ent/form"
	"github.com/User/app/ent/user"
	"github.com/User/app/ent/unit"
	"github.com/User/app/ent/volume"
	"github.com/gin-gonic/gin"
)

// DrugController defines the struct for the drug controller
type DrugController struct {
	client *ent.Client
	router gin.IRouter
}

type Drug struct {
	User   int
	Unit   int
	Form   int
	Volume   int
	Strength   int
	Information   string	
	DrugType string
}

// CreateDrug handles POST requests for adding drug entities
// @Summary Create drug
// @Description Create drug
// @ID create-drug
// @Accept   json
// @Produce  json
// @Param drug body ent.Drug true "Drug entity"
// @Success 200 {object} ent.Drug
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /drugs [post]
func (ctl *DrugController) CreateDrug(c *gin.Context) {
	obj := Drug{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "drug binding failed",
		})
		return
	}

	un, err := ctl.client.Unit.
		Query().
		Where(unit.IDEQ(int(obj.Unit))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": "unit not found",
		})
		return
	}

	u, err := ctl.client.User.
		Query().
		Where(user.IDEQ(int(obj.User))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": "user not found",
		})
		return
	}

	f, err := ctl.client.Form.
		Query().
		Where(form.IDEQ(int(obj.Form))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": "form not found",
		})
		return
	}

	v, err := ctl.client.Volume.
		Query().
		Where(volume.IDEQ(int(obj.Volume))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": "volume not found",
		})
		return
	}


	dg, err := ctl.client.Drug.
		Create().
		SetUnit(un).
		SetUser(u).
		SetForm(f).
		SetVolume(v).
		SetDrugType(obj.DrugType).
		SetStrength(obj.Strength).
		SetInformation(obj.Information).
		Save(context.Background())
		
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, gin.H{
		"status": true,
		"data": dg,
	})
}

// GetDrug handles GET requests to retrieve a drug entity
// @Summary Get a drug entity by ID
// @Description get drug by ID
// @ID get-drug
// @Produce  json
// @Param id path int true "Drug ID"
// @Success 200 {object} ent.Drug
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /drugs/{id} [get]
func (ctl *DrugController) GetDrug(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	dg, err := ctl.client.Drug.
		Query().
		Where(drug.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, dg)
}

// ListDrug handles request to get a list of drug entities
// @Summary List drug entities
// @Description list drug entities
// @ID list-drug
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Drug
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /drugs [get]
func (ctl *DrugController) ListDrug(c *gin.Context) {
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

	drugs, err := ctl.client.Drug.
		Query().
		Limit(limit).
		Offset(offset).
		WithUser().
		WithForm().
		WithVolume().
		WithUnit().
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, drugs)
}

// DeleteDrug handles DELETE requests to delete a drug entity
// @Summary Delete a drug entity by ID
// @Description get drug by ID
// @ID delete-drug
// @Produce  json
// @Param id path int true "Drug ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /drugs/{id} [delete]
func (ctl *DrugController) DeleteDrug(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Drug.
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

// UpdateDrug handles PUT requests to update a drug entity
// @Summary Update a drug entity by ID
// @Description update drug by ID
// @ID update-drug
// @Accept   json
// @Produce  json
// @Param id path int true "Drug ID"
// @Param drug body ent.Drug true "Drug entity"
// @Success 200 {object} ent.Drug
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /drugs/{id} [put]
func (ctl *DrugController) UpdateDrug(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Drug{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "drug binding failed",
		})
		return
	}
	obj.ID = int(id)
	dg, err := ctl.client.Drug.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, dg)
}

// NewDrugController creates and registers handles for the drug controller
func NewDrugController(router gin.IRouter, client *ent.Client) *DrugController {
	dc := &DrugController{
		client: client,
		router: router,
	}
	dc.register()
	return dc
}

// InitDrugController registers routes to the main engine
func (ctl *DrugController) register() {
	drugs := ctl.router.Group("/drugs")

	drugs.GET("", ctl.ListDrug)

	// CRUD
	drugs.POST("", ctl.CreateDrug)
	drugs.GET(":id", ctl.GetDrug)
	// drugs.PUT(":id", ctl.UpdateDrug)
	// drugs.DELETE(":id", ctl.DeleteDrug)
}

