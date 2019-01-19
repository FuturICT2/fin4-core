package tokenservice

import (
	"errors"

	"github.com/FuturICT2/fin4-core/server/datatype"
	"github.com/lytics/logrus"
)

// ApproveClaim approves a claim
func (db *Service) ApproveClaim(
	claimID datatype.ID,
	approverID datatype.ID,
) error {
	claim, err := db.FindClaim(claimID)
	if err != nil {
		logrus.WithFields(
			logrus.Fields{"e": err.Error()},
		).Error("user:ApproveActionClaim:0")
		return datatype.ErrServerError
	}
	if claim.IsApproved == true {
		return errors.New("Claim already approved")
	}
	tx, err := db.Begin()
	if err != nil {
		logrus.WithFields(
			logrus.Fields{"e": err.Error()},
		).Error("user:ApproveActionClaim:0")
		return datatype.ErrServerError
	}
	defer tx.Rollback()
	_, err = tx.Exec(`
			UPDATE token
			SET
				totalSupply = totalSupply + 1
			WHERE id=?`,
		claim.TokenID,
	)
	if err != nil {
		logrus.WithFields(
			logrus.Fields{"e": err.Error()},
		).Error("user:ApproveActionClaim:1")
		return datatype.ErrServerError
	}
	_, err = tx.Exec(`
		INSERT INTO user_balance (userId, tokenId, balance, reserved)
		VALUES (?, ?, 1, 0)
		ON DUPLICATE KEY UPDATE
			balance = balance + 1`,
		claim.UserID,
		claim.TokenID,
	)
	if err != nil {
		logrus.WithFields(
			logrus.Fields{"e": err.Error()},
		).Error("user:ApproveActionClaim:1.1")
		return datatype.ErrServerError
	}
	_, err = tx.Exec(`
			UPDATE claim
			SET
				isApproved = true
			WHERE id=?`,
		claim.ID,
	)
	if err != nil {
		logrus.WithFields(
			logrus.Fields{"e": err.Error()},
		).Error("user:ApproveActionClaim:1.4")
		return datatype.ErrServerError
	}
	err = tx.Commit()
	if err != nil {
		logrus.WithFields(
			logrus.Fields{"e": err.Error()},
		).Error("user:ApproveActionClaim:3")
		return datatype.ErrServerError
	}
	return nil
}
