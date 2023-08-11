package gin

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Login Page
func (h *Handler) HomeHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"topics": []struct {
			Name        string
			Tags        string
			Comment     string
			Description string
		}{
			{
				Name:        "Study of Linked List",
				Tags:        "computer science, data structures, algorithms",
				Description: "Flashcards and resumes about linked lists.",
			},
			{
				Name:        "Study of Hash Table",
				Tags:        "computer science, data structures, algorithms",
				Description: "Flashcards and resumes about Hash Table.",
			},
			{
				Name:        "Study of Binary Tree",
				Tags:        "computer science, data structures, algorithms",
				Description: "Flashcards and resumes about Binary Tree.",
			},
			{
				Name:        "Study of Bubble Sort",
				Tags:        "computer science, data structures, algorithms",
				Description: "Flashcards and resumes about Bubble Sort.",
			},
			{
				Name:        "Study of Linked List",
				Tags:        "computer science, data structures, algorithms",
				Description: "Flashcards and resumes about linked lists.",
			},
			{
				Name:        "Study of Hash Table",
				Tags:        "computer science, data structures, algorithms",
				Description: "Flashcards and resumes about Hash Table.",
			},
			{
				Name:        "Study of Binary Tree",
				Tags:        "computer science, data structures, algorithms",
				Description: "Flashcards and resumes about Binary Tree.",
			},
			{
				Name:        "Study of Bubble Sort",
				Tags:        "computer science, data structures, algorithms",
				Description: "Flashcards and resumes about Bubble Sort.",
			},
			{
				Name:        "Study of Linked List",
				Tags:        "computer science, data structures, algorithms",
				Description: "Flashcards and resumes about linked lists.",
			},
			{
				Name:        "Study of Hash Table",
				Tags:        "computer science, data structures, algorithms",
				Description: "Flashcards and resumes about Hash Table.",
			},
			{
				Name:        "Study of Binary Tree",
				Tags:        "computer science, data structures, algorithms",
				Description: "Flashcards and resumes about Binary Tree.",
			},
			{
				Name:        "Study of Bubble Sort",
				Tags:        "computer science, data structures, algorithms",
				Description: "Flashcards and resumes about Bubble Sort.",
			},
		},
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

	studyCase, err := h.StudyCaseService.GetStudyCase(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.HTML(http.StatusOK, "study_case.html", gin.H{
		"study_case": studyCase,
	})
}


