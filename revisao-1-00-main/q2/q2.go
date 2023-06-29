package q2

import (
	"errors"
	"strings"
)

func AverageLettersPerWord(text string) (float64, error) {
	// Implemente sua solução aqui
	if len(text) == 0 || strings.TrimSpace(text) == "" {
		return 0, errors.New("texto vazio")
	}

	palavras := strings.Fields(text)
	wordCount := len(palavras)
	contagem := 0

	for _, word := range palavras {
		contagem += len(word)
	}

	average := float64(contagem) / float64(wordCount)
	return average, nil
}
