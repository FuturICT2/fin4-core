package models_test

import (
	"testing"

	"github.com/FuturICT2/fin4-core/server/models"
)

func Test_NewUser(t *testing.T) {
	u := models.NewUser()
	if len(u.Salt) < 32 {
		t.Error("User default Salt must be 32 char long")
	}
	if u.AgreeToTerms != models.DbFalse {
		t.Error("Default AgreeToTerms must be false")
	}
	if u.IsDeleted != models.DbFalse {
		t.Error("Default IsDeleted must be false")
	}
}

func Test_NewUserResetToken(t *testing.T) {
	userID := models.ID(44)
	token := models.NewUserResetToken(userID)
	if len(token.Token) < 64 {
		t.Error("token.Token must not be less than 64 char long")
	}
	if token.UserID != userID {
		t.Error("userID must be set correctly")
	}
}
