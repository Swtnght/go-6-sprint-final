package service

import (
	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
	"strings"
)

func DetectAndConvert(input string) (string, error) {
	input = strings.TrimSpace(input)
	if input == "" {
		return "", nil
	}

	cleaned := strings.Trim(input, ".- ")
	if cleaned == "" {
		return morse.ToText(input), nil
	}

	return morse.ToMorse(input), nil
}
