package controllers

import (
	"github.com/lavender-snow/get-words/app/models"
)

// GetWord DBからランダムに単語を返却
func GetWord() (string, error) {
	word, err := models.SelectWord()
	if err != nil {
		return "", err
	}
	return word, nil
}

func GetPokemon() (string, error) {
	pokemon, err := models.SelectPokemon()
	if err != nil {
		return "", err
	}
	return pokemon, nil
}
