package gin

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateFlashCardHandler(c *gin.Context) {
	id := c.Param("id")

	var flashCardDTO struct {
		Front string `json:"front"`
		Back  string `json:"back"`
		// status string `json:"status"`
	}

	err := c.ShouldBindJSON(&flashCardDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if flashCardDTO.Back == "" || flashCardDTO.Front == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "front and back are required"})
		return
	}

	currentTime := time.Now()

	flashCard, err := h.StudyCaseService.CreateFlashCard(flashCardDTO.Front, flashCardDTO.Back, currentTime.Add(24*time.Hour), id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"flashcard": flashCard})
}
