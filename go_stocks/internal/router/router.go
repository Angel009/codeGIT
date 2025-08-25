package router

import (
	"database/sql"
	"stocks/internal/handlers"
	"stocks/internal/repository"
	"stocks/internal/services"

	"github.com/gin-gonic/gin"
)

func Setup(db *sql.DB, apiURL string) *gin.Engine {
	r := gin.Default()

	// CORS
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	apiDataRepo := repository.NewApiDataRepository(db)
	analysisService := services.NewAnalysisService()

	api := r.Group("/api")
	{
		api.GET("/external", handlers.GetApiData(apiURL, apiDataRepo, analysisService))
	}

	return r
}
