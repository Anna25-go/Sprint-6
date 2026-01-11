package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func MorseCode(inputData string) bool {
	for _, v := range inputData {
		if v != '.' && v != '-' && v != ' ' {
			return false
		}
	}
	return true
}

func Translate(inputData string) (outputData string, err error) {
	inputData = strings.TrimSpace(inputData)
	if inputData == "" {
		return "", errors.New("the input data cannot be empty")
	}
	if MorseCode(inputData) {
		return morse.ToText(inputData), err
	} else {
		return morse.ToMorse(inputData), err
	}
}
