package gin

import (
	"log"
	"net/http"
	"os"

	"github.com/caiostarke/hermes/study_case"

	"github.com/gin-gonic/gin"
)

func NewHandler(svc *study_case.Service) *Handler {
	return &Handler{StudyCaseService: *svc}
}

type Handler struct {
	StudyCaseService study_case.Service
}

func SetupRoutes(h *Handler) {
	port := os.Getenv("PORT")
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "./static")

	r.GET("/", h.HomeHandler)
	r.GET("/topic/:id", h.GetStudyCaseHandler)
	r.POST("/api/v1/study_case", h.CreateStudyCaseHandler)

	topic := r.Group("/topic/:id")
	{
		topic.POST("/flashcard", h.CreateFlashCardHandler)
	}


	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
}
