package handlers

import (
	"strings"

	"github.com/caiostarke/hermes/text-ingestor/internal/entity"
	"github.com/caiostarke/hermes/text-ingestor/internal/infra/database"
)

var (
	FrequencyValue = 3
)

type TextHandler struct {
	Repo database.TextInterface
}

func NewTextHandler(repo database.TextInterface) *TextHandler {
	return &TextHandler{Repo: repo}
}

func (h *TextHandler) ReturnFilteredTexts() ([]entity.Text, error) {
	texts, err := h.Repo.FindAll()
	if err != nil {
		return nil, err
	}

	// stripe the texts retrieved from the database into tokens
	// define a frequency a token appears
	// mark tokens as true based on frequency they appear
	// get the text that contains the token

	var tokens []string
	frequencyMap := make(map[string]int)
	markedTokens := make(map[string]bool)

	// stripe the texts retrieved from the database into tokens
	for _, text := range texts {
		tokens = append(tokens, strings.Split(strings.TrimSpace(text.Text), " ")...)
	}

	// define a frequency a token appears
	for _, v := range tokens {
		frequencyMap[v]++
	}

	for value, frequency := range frequencyMap {
		if frequency >= FrequencyValue {
			markedTokens[value] = true
		}
	}

	filteredTokens := []string{}
	for _, token := range tokens {
		if markedTokens[token] {
			filteredTokens = append(filteredTokens, token)
			markedTokens[token] = false
		}
	}

	filteredTexts := []entity.Text{}
	for _, v := range filteredTokens {
		for _, text := range texts {
			if strings.Contains(text.Text, v) {
				if !ContainText(filteredTexts, text.Text) {
					filteredTexts = append(filteredTexts, text)
				}
			}
		}
	}

	return filteredTexts, nil
}

// return true if exists a text into texts
func ContainText(texts []entity.Text, text string) bool {
	for _, v := range texts {
		if strings.TrimSpace(v.Text) == strings.TrimSpace(text) {
			return true
		}
	}

	return false
}
