package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/project-Grok3_Golang/internal/ai"
)

var aiClient *ai.AIClient

func InitAIClient() {
	aiClient = ai.NewAIClient("your-api-key", "https://api.example.com")
}

func Predict(c *gin.Context) {
	var input struct {
		Text string `json:"text"`
	}
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := aiClient.Predict(input.Text)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": result})
}
