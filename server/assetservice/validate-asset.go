package assetservice

import (
	"errors"
	"strings"
	"unicode"
)

// ValidateAsset checks whether a user is a miner of a token or not
func (db *Service) ValidateAsset(
	name string,
	symbol string,
	description string,
) error {
	if len(name) > 65 || len(name) < 3 {
		return errors.New("Topic name has to be between 3 and 64 characters")
	}
	if len(symbol) < 3 || len(symbol) > 8 {
		return errors.New("Topic symbol has to be between 3 and 8 characters")
	}
	if len(strings.Fields(symbol)) > 1 || !isLetter(symbol) {
		return errors.New("Topic symbol should have have letters only and no whitespaces")
	}
	if len(description) < 3 || len(description) > 1024 {
		return errors.New("Topic description has to be between 3 and 1024 characters")
	}
	return nil
}

func isLetter(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}
