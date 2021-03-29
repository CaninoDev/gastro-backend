package ginHTTP

import (
	"log"
	"net/http"

	"github.com/CaninoDev/gastro/server/api/account"
	"github.com/CaninoDev/gastro/server/api/authentication"

	"github.com/CaninoDev/gastro/server/api/menu"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type menuHandler struct {
	menuSvc menu.Service
	authSvc authentication.Service
}

// RegisterRoutes sets up menu API endpoint using Gin has the delivery.
func RegisterRoutes(svc menu.Service, authSvc authentication.Service, r *gin.Engine, authMiddleWare gin.HandlerFunc) {
	h := menuHandler{svc, authSvc}
	publicRoutes(r, &h)
	privateRoutes(r, &h, authMiddleWare)
}

func publicRoutes(r *gin.Engine, h *menuHandler) {
	menuGroup := r.Group("/api/v1")
	menuViewGroup := menuGroup.Group("")
	menuViewGroup.GET("/sections", h.listSections)
	menuViewGroup.GET("/sections/:id", h.findSectionByID)
	menuViewGroup.GET("/items", h.listItems)
	menuViewGroup.GET("/items/:id", h.findItemByID)
}

func privateRoutes(r *gin.Engine, h *menuHandler, authMiddleWare gin.HandlerFunc) {
	menuEditGroup := r.Group("", authMiddleWare)
	menuEditGroup.POST("/sections", h.createSection)
	menuEditGroup.PATCH("/sections/:id", h.updateSection)
	menuEditGroup.DELETE("/sections/:id", h.deleteSection)
	menuEditGroup.POST("/items", h.createItem)
	menuEditGroup.PATCH("/items/:id", h.updateItem)
	menuEditGroup.DELETE("/items/:id", h.deleteItem)
}


// --- Sections --- //
func (h *menuHandler) listSections(ctx *gin.Context) {
	sections, err := h.menuSvc.Sections(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": sections})

}

func (h *menuHandler) findSectionByID(ctx *gin.Context) {
	rawID := ctx.Param("id")
	log.Printf("ID: %s", rawID)
	section, err := h.menuSvc.SectionByID(ctx, rawID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": section})

}

// createSection creates a new section.
func (h *menuHandler) createSection(ctx *gin.Context) {
	role, exists := ctx.Get("role")
	if !exists || role != account.Admin {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	var section menu.Section

	if err := ctx.ShouldBindJSON(&section); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.menuSvc.NewSection(ctx, &section); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": section})
}


// updateSection update section's data.
func (h *menuHandler) updateSection(ctx *gin.Context) {
	role, exists := ctx.Get("role")
	if !exists || role != account.Admin {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	var section menu.Section

	if err := ctx.ShouldBindJSON(&section); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rawID := ctx.Param("id")
	id, err := uuid.Parse(rawID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	section.ID = id
	if err := h.menuSvc.UpdateSectionData(ctx, &section); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": section})
}

func (h *menuHandler) deleteSection(ctx *gin.Context) {
	role, exists := ctx.Get("role")
	if !exists || role != account.Admin {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	rawID := ctx.Param("id")
	if err := h.menuSvc.DeleteSection(ctx, rawID); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "section deleted"})

}


// ---  Item  --- //
func (h *menuHandler) listItems(ctx *gin.Context) {
	items, err := h.menuSvc.Items(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": items})

}

// createSection creates a new section.
func (h *menuHandler) createItem(ctx *gin.Context) {
	role, exists := ctx.Get("role")
	if !exists || role != account.Admin {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	var item menu.Item

	if err := ctx.ShouldBindJSON(&item); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.menuSvc.NewItem(ctx, &item); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": item})
}

// updateSection creates a new section.
func (h *menuHandler) updateItem(ctx *gin.Context) {
	role, exists := ctx.Get("role")
	if !exists || role != account.Admin {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	rawID := ctx.Param("id")
	id, err := uuid.Parse(rawID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	var item menu.Item

	if err := ctx.ShouldBindJSON(&item); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item.ID = id
	if err := h.menuSvc.UpdateItemData(ctx, &item); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": item})
}
func (h *menuHandler) findItemByID(ctx *gin.Context) {
	rawID := ctx.Param("id")
	item, err := h.menuSvc.ItemByID(ctx, rawID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": item})

}

func (h *menuHandler) deleteItem(ctx *gin.Context) {
	role, exists := ctx.Get("role")
	if !exists || role != account.Admin {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	rawID := ctx.Param("id")

	if err := h.menuSvc.DeleteItem(ctx, rawID); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "item deleted"})

}