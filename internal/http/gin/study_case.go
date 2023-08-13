package gin

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/caiostarke/hermes/study_case"
	"github.com/gin-gonic/gin"
)

// Login Page
func (h *Handler) HomeHandler(c *gin.Context) {
	studyCases, err := h.StudyCaseService.ListStudyCases()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.HTML(http.StatusOK, "index.html", gin.H{
		"data": studyCases,
	})
}

type studyCaseJson struct {
	Name        string          `json:"name"`
	Tags        json.RawMessage `json:"tags"`
	Comment     string          `json:"comment"`
	Description string          `json:"description"`
}

func (h *Handler) CreateStudyCaseHandler(c *gin.Context) {
	sc := studyCaseJson{}

	err := c.ShouldBindJSON(&sc)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if sc.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "name is required"})
		return
	}

	studyCase, err := h.StudyCaseService.CreateStudyCase(sc.Name, sc.Tags, sc.Comment, sc.Description)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": studyCase})
}

func (h *Handler) GetStudyCaseHandler(c *gin.Context) {
	id := c.Param("id")
	data := study_case.StudyCaseMetadata{}

	studyCase, err := h.StudyCaseService.GetStudyCase(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	flashcards, err := h.StudyCaseService.ListFlashCards(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(flashcards)

	data.FlashCards = flashcards
	data.Name = studyCase.Name
	data.Comments = studyCase.Comment.String
	if studyCase.Tags.Valid {
		if err = json.Unmarshal(studyCase.Tags.RawMessage, &data.Tags); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	c.HTML(http.StatusOK, "study_case.html", gin.H{
		"data": data,
		"id":   id,
	})
}
