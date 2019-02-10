package assetservice

import (
	"errors"
	"fmt"
	"time"

	"github.com/FuturICT2/fin4-core/server/apperrors"
	"github.com/FuturICT2/fin4-core/server/datatype"
	"github.com/FuturICT2/fin4-core/server/helpers"
	"mvdan.cc/xurls"
)

// CreateAssetBlock insert asset block
func (db *Service) CreateAssetBlock(
	userID datatype.ID,
	assetID datatype.ID,
	blockText string,
	images []string,
) (*datatype.Block, error) {
	isEmptyText := len(blockText) < 1
	isEmptyImages := len(images) < 1
	if isEmptyText && isEmptyImages {
		return nil, errors.New("Post should have text or images")
	}
	if !isEmptyText && len(blockText) > 512 {
		return nil, errors.New("Post text is limited to 512 characters")
	}
	asset, err := db.FindByID(assetID)
	if err != nil {
		apperrors.Critical("assetservice:CreateAssetBlock:1", err)
		return nil, datatype.ErrServerError
	}
	if asset == nil {
		return nil, errors.New("Asset doesn't exist")
	}
	urls := xurls.Relaxed().FindAllString(blockText, -1)
	var ytVideoID string
	for _, url := range urls {
		videoID := helpers.FindVideoID(url)
		if len(videoID) > 0 {
			ytVideoID = videoID
			break
		}
	}
	// Start db transaction
	tx, err := db.Begin()
	if err != nil {
		apperrors.Critical("assetservice:CreateAssetBlock:2", err)
		return nil, err
	}
	defer tx.Rollback()
	res, err := tx.Exec(
		`INSERT INTO asset_block SET
	        assetId = ?,
					userId = ?,
					text = ?,
					status = ?,
					videoID = ?,
					createdAt = ?
	      `,
		assetID,
		userID,
		blockText,
		datatype.BlockUnverified,
		ytVideoID,
		time.Now(),
	)
	if err != nil {
		apperrors.Critical("assetservice:CreateAssetBlock:3", err)
		return nil, datatype.ErrServerError
	}
	blockID, err := res.LastInsertId()
	if err != nil {
		apperrors.Critical("assetservice:CreateAssetBlock:4", err)
		return nil, datatype.ErrServerError
	}
	// insert images if there is any
	if len(images) > 0 {
		imgsStmt := `
			INSERT INTO asset_block_image
	    	(blockId, filepath)
			VALUES `
		for idx, image := range images {
			imgsStmt += fmt.Sprintf(" (%d, '%s')", blockID, image)
			if idx+1 != len(images) {
				imgsStmt += ", "
			}
		}
		res, err = tx.Exec(imgsStmt)
		if err != nil {
			apperrors.Critical("assetservice:CreateAssetBlock:5", err)
			return nil, datatype.ErrServerError
		}
	}
	/*** Commit db transaction ***/
	err = tx.Commit()
	if err != nil {
		apperrors.Critical("assetservice:CreateAssetBlock:6", err)
		return nil, datatype.ErrServerError
	}
	return nil, nil
}
