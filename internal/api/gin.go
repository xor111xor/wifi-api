package api

import (
	"github.com/gin-gonic/gin"
	"github.com/xor111xor/wifi-api/internal/repository"
	"net/http"
)

func RunApi(repo *repository.InMemoryRepo) {
	// Create Gin router
	router := gin.Default()

	// Instantiate Handler
	recipesHandler := NewRecipesHandler(repo)

	// Register Routes
	router.GET("/api/:date", recipesHandler.GetMetrics)

	// Start the server
	err := router.Run()
	if err != nil {
		panic(err)
	}
}

type RecipesHandler struct {
	store *repository.InMemoryRepo
}

func NewRecipesHandler(s *repository.InMemoryRepo) *RecipesHandler {
	return &RecipesHandler{
		store: s,
	}
}

func (h RecipesHandler) GetMetrics(c *gin.Context) {
	date := c.Param("date")

	recipe, err := h.store.GetMetricsFromString(date)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}

	c.JSON(200, recipe)
}
