package tokenservice

import (
	"strconv"

	"github.com/FuturICT2/fin4-core/server/datatype"
	"github.com/lytics/logrus"
)

func (db *Service) FindTokenForUser(
	id datatype.ID,
	userID datatype.ID,
) (*datatype.Token, error) {
	token, err := db.findTokenBy("id", strconv.Itoa(int(id)))
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("tokenservice:FindTokenForUser:e1")
		return nil, err
	}
	token.FavouriteCount = db.CountTokenLikes(token.ID)
	token.DidUserLike = db.DidUserLike(userID, token.ID)
	claims, err := db.GetTokenClaims(userID, token.ID)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("tokenservice:FindTokenForUser:e2")
		return nil, err
	}
	token.Claims = claims
	miners, err := db.getTokenMiners(token.ID)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("tokenservice:FindTokenForUser:e3")
		return nil, err
	}
	token.Miners = miners
	likers, err := db.getTokenLikers(token.ID)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("tokenservice:FindTokenForUser:e4")
		return nil, err
	}
	token.Likers = likers
	return token, nil
}
