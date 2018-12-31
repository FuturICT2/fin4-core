package models

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/FuturICT2/fin4-core/server/decimaldt"
	"github.com/sirupsen/logrus"
)

// User user type
type User struct {
	ID              ID        `json:"id"`
	Name            string    `json:"name"`
	EthereumAddress string    `json:"ethereumAddress"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}

type ActionProposal struct {
	ID          ID             `json:"id"`
	Description string         `json:"description"`
	DoerID      ID             `json:"doerId"`
	DoerName    string         `json:"doerName"`
	LogoFile    sql.NullString `json:"logoFile"`
	Approved    bool           `json:"approved"`
	CreatedAt   time.Time      `json:"createdAt"`
	IsOwner     bool           `json:"isOwner"`
}

type ActionSupporter struct {
	Amount    decimaldt.Decimal `json:"amount"`
	TokenId   ID                `json:"tokenId"`
	TokenName string            `json:"tokenName"`
	UserId    ID                `json:"userId"`
	UserName  string            `json:"userName"`
	Status    int               `json:"status"`
}

type Action struct {
	ID          ID                `json:"id"`
	Description string            `json:"description"`
	CreatorID   ID                `json:"creatorID"`
	CreatorName string            `json:"creatorName"`
	Status      int               `json:"status"`
	LogoFile    sql.NullString    `json:"logoFile"`
	StartsAt    time.Time         `json:"startsAt"`
	EndsAt      time.Time         `json:"endsAt"`
	CreatedAt   time.Time         `json:"createdAt"`
	Proposals   []ActionProposal  `json:"proposals"`
	Supporters  []ActionSupporter `json: "supporters"`
}

// NewUser creates a new user
func NewUser() *User {
	return &User{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

const (
	tableUser = "user"
)

// UserStore interface for user store
type UserStore interface {
	Register(name string, address string) (*User, error)
	FindByID(ID) (*User, error)
	GetUserBalances(userId ID) ([]Balance, error)
	InsertBalance(
		userID ID,
		tokenID ID,
		balance string,
	) error
	InsertToken(
		userID ID,
		name string,
		symbol string,
		purpose string,
		totalSupply string,
		blockchainAddress string,
		txAddress string,
		logo string,
	) (*Token, error)
	FindByName(string) (*User, error)
	IsNameRegistered(string) bool
	FindTokens(userID ID) ([]Token, error)
	FindUsers() ([]User, error)
	DoLike(userID ID, tokenID ID, state bool) error
	InsertAction(
		userID ID,
		description string,
		startsAt time.Time,
		endsAt time.Time,
	) error
	ReserveRewardsForAction(
		userID ID,
		tokenID ID,
		actionID ID,
		amount decimaldt.Decimal,
	) error
	FindActions(userid ID) ([]Action, error)
	AddActionProposal(userID ID, proposal string, actionID ID) error
	AprroveProposal(claimID ID, approverID ID) error
}

func (db *UserModel) FindClaim(claimID ID) (*Claim, error) {
	row := db.QueryRow(`SELECT
			c.id,
			c.tokenId,
			c.userId,
			u.name,
			c.text,
			c.logoFile,
			c.isApproved
		FROM claim c
		LEFT JOIN
			user u ON u.id = c.userId
		WHERE c.id=?`,
		claimID,
	)
	entry := Claim{}
	err := row.Scan(
		&entry.ID,
		&entry.TokenID,
		&entry.UserID,
		&entry.UserName,
		&entry.Text,
		&entry.LogoFile,
		&entry.IsApproved,
	)
	if err != nil {
		return nil, err
	}
	return &entry, nil
}

func (db *UserModel) AprroveProposal(claimID ID, approverID ID) error {
	claim, err := db.FindClaim(claimID)
	if err != nil {
		logrus.WithFields(
			logrus.Fields{"e": err.Error()},
		).Error("user:AprroveProposal:0")
		return ErrServerError
	}
	token := db.FindToken(claim.TokenID)
	if token != nil && token.CreatorID != approverID {
		return errors.New("Token creator only can approve")
	}
	if claim.IsApproved == true {
		return errors.New("Claim already approved")
	}
	tx, err := db.Begin()
	if err != nil {
		logrus.WithFields(
			logrus.Fields{"e": err.Error()},
		).Error("user:AprroveProposal:0")
		return ErrServerError
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
		).Error("user:AprroveProposal:1")
		return ErrServerError
	}
	_, err = tx.Exec(`
		INSERT INTO user_holding (userId, tokenId, balance, reserved)
		VALUES (?, ?, 1, 0)
		ON DUPLICATE KEY UPDATE
			balance = balance + 1`,
		claim.UserID,
		claim.TokenID,
	)
	if err != nil {
		logrus.WithFields(
			logrus.Fields{"e": err.Error()},
		).Error("user:AprroveProposal:1.1")
		return ErrServerError
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
		).Error("user:AprroveProposal:1.4")
		return ErrServerError
	}
	err = tx.Commit()
	if err != nil {
		logrus.WithFields(
			logrus.Fields{"e": err.Error()},
		).Error("user:AprroveProposal:3")
		return ErrServerError
	}
	return nil
}

func (db *UserModel) AddActionProposal(
	userID ID,
	proposal string,
	actionID ID,
) error {
	res, err := db.Exec(
		`INSERT INTO claim SET
			tokenId = ?,
			userId = ?,
			text = ?,
			isApproved = 0,
			logoFile = ""`,
		actionID,
		userID,
		proposal,
	)
	if err != nil {
		logrus.WithFields(
			logrus.Fields{"e": err.Error()},
		).Error("user:AddActionProposal:0")
		return ErrServerError
	}
	_, err = res.LastInsertId()
	if err != nil {
		logrus.WithFields(
			logrus.Fields{"e": err.Error()},
		).Error("user:AddActionProposal:1")
		return ErrServerError
	}
	return nil
}

func (db *UserModel) FindActions(userID ID) ([]Action, error) {
	result := []Action{}
	rows, err := db.Query(
		fmt.Sprintf(`
			SELECT
				a.id,
				a.description,
				a.status,
				a.creatorId,
				u.name,
				a.startsAt,
				a.endsAt,
				a.logoFile,
				a.createdAt
			FROM action a
			LEFT JOIN
				user u ON a.creatorId = u.id
			ORDER BY createdAt DESC`),
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var a Action
		err := rows.Scan(
			&a.ID,
			&a.Description,
			&a.Status,
			&a.CreatorID,
			&a.CreatorName,
			&a.StartsAt,
			&a.EndsAt,
			&a.LogoFile,
			&a.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		supporters, err := db.GetActionSupporters(a.ID)
		if err != nil {
			return nil, err
		}
		a.Supporters = supporters
		proposals, err := db.GetActionProposals(a.ID, userID)
		if err != nil {
			return nil, err
		}
		a.Proposals = proposals
		result = append(result, a)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return result, nil
}

func (db *UserModel) GetActionProposals(actionID ID, userID ID) ([]ActionProposal, error) {
	result := []ActionProposal{}
	// ID          ID             `json:"id"`
	// Description string         `json:"description"`
	// DoerID      ID             `json:"doerId"`
	// DoerName    string         `json:"doerName"`
	// LogoFile    sql.NullString `json:"logoFile"`
	// Approved    bool           `json:"approved"`
	// CreatedAt   time.Time      `json:"createdAt"`
	rows, err := db.Query(`
		SELECT
			ap.id,
			ap.proposalTxt,
			ap.isApproved,
			ap.createdAt,
			ap.userId,
			u.name
		FROM action_proposal ap
		LEFT JOIN
			user u ON ap.userId = u.id
		WHERE actionId=?`,
		actionID,
	)
	if err != nil {
		logrus.WithFields(
			logrus.Fields{"e": err.Error()},
		).Error("user:GetActionProposals:0")
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var a ActionProposal
		err := rows.Scan(
			&a.ID,
			&a.Description,
			&a.Approved,
			&a.CreatedAt,
			&a.DoerID,
			&a.DoerName,
		)
		if err != nil {
			logrus.WithFields(
				logrus.Fields{"e": err.Error()},
			).Error("user:GetActionProposals:2")
			return nil, err
		}
		a.IsOwner = a.DoerID == userID
		result = append(result, a)
	}
	if err := rows.Err(); err != nil {
		logrus.WithFields(
			logrus.Fields{"e": err.Error()},
		).Error("user:GetActionProposals:3")
		return nil, err
	}
	return result, nil
}

func (db *UserModel) GetActionSupporters(actionId ID) ([]ActionSupporter, error) {
	result := []ActionSupporter{}
	// Amount
	// TokenId
	// TokenName
	// UserId
	// UserName
	rows, err := db.Query(`
		SELECT
			ar.amount,
			ar.tokenId,
			t.name,
			ar.userId,
			u.name,
			ar.status
		FROM action_reward ar
		LEFT JOIN
			user u ON ar.userId = u.id
		LEFT JOIN
			token t ON ar.tokenId = t.id
		WHERE actionId=?`,
		actionId,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var a ActionSupporter
		err := rows.Scan(
			&a.Amount,
			&a.TokenId,
			&a.TokenName,
			&a.UserId,
			&a.UserName,
			&a.Status,
		)
		if err != nil {
			return nil, err
		}
		result = append(result, a)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return result, nil
}

func (db *UserModel) InsertAction(
	userID ID,
	description string,
	startsAt time.Time,
	endsAt time.Time,
) error {
	res, err := db.Exec(
		`INSERT INTO action SET
			description = ?,
			status = 0,
			creatorId = ?,
			startsAt = ?,
			endsAt = ?,
			createdAt = ?`,
		description,
		userID,
		startsAt,
		endsAt,
		time.Now(),
	)
	if err != nil {
		return ErrServerError
	}
	id, err := res.LastInsertId()
	if err != nil {
		return ErrServerError
	}
	// by default the app demo adds rewards for doing this actions in GDBG token
	err = db.ReserveRewardsForAction(userID, ID(1), ID(id), decimaldt.NewFromFloat(1))
	return err
}

func (db *UserModel) ReserveRewardsForAction(
	userID ID,
	tokenID ID,
	actionID ID,
	amount decimaldt.Decimal,
) error {
	// check if user has balance
	var balance decimaldt.Decimal
	err := db.QueryRow(
		"SELECT balance FROM user_holding WHERE userId=? and tokenId= ?",
		userID,
		tokenID,
	).Scan(&balance)
	if amount.IsBiggerThan(balance) {
		return errors.New("You don't hanve enough balance")
	}
	// reduce amount from balance
	// add amount to reserved
	// insert an entry in the action_rwards table
	/*** start db transaction ***/
	tx, err := db.Begin()
	if err != nil {
		logrus.WithFields(
			logrus.Fields{"e": err.Error()},
		).Error("user:ReserveRewardsForAction:0")
		return ErrServerError
	}
	defer tx.Rollback()
	_, err = tx.Exec(`
		UPDATE user_holding
		SET
			balance = balance - ?,
			reserved = reserved + ?
		WHERE userId = ? AND tokenId = ?`,
		amount,
		amount,
		userID,
		tokenID,
	)
	if err != nil {
		logrus.WithFields(
			logrus.Fields{"e": err.Error()},
		).Error("user:ReserveRewardsForAction:1")
		return ErrServerError
	}
	_, err = tx.Exec(`
		INSERT INTO action_reward SET
			userId = ?,
			tokenId = ?,
			actionId = ?,
			amount = ?,
			status = 0
		ON DUPLICATE KEY UPDATE
			amount = amount + ?`,
		userID,
		tokenID,
		actionID,
		amount,
		amount,
	)
	if err != nil {
		logrus.WithFields(
			logrus.Fields{"e": err.Error()},
		).Error("user:ReserveRewardsForAction:2")
		return ErrServerError
	}
	err = tx.Commit()
	if err != nil {
		logrus.WithFields(
			logrus.Fields{"e": err.Error()},
		).Error("user:ReserveRewardsForAction:3")
		return ErrServerError
	}
	return nil
}

//DoLike Registers a new user
func (db *UserModel) DoLike(userID ID, tokenID ID, state bool) error {
	if state == false {
		_, err := db.Exec(
			`delete from token_like where
				userId = ? and  tokenId = ?`,
			userID,
			tokenID,
		)
		if err != nil {
			return ErrServerError
		}
		return nil
	} else {
		_, err := db.Exec(
			`INSERT INTO token_like SET
				userId = ?,
				tokenId = ?`,
			userID,
			tokenID,
		)
		if err != nil {
			return ErrServerError
		}
		return nil
	}

}

// FindUsers finds all users
func (db *UserModel) FindUsers() ([]User, error) {
	result := []User{}
	rows, err := db.Query(
		fmt.Sprintf(`SELECT id, name, ethereumAddress FROM user`),
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var c User
		err := rows.Scan(
			&c.ID,
			&c.Name,
			&c.EthereumAddress,
		)
		if err != nil {
			return nil, err
		}
		result = append(result, c)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return result, nil
}

//Register Registers a new user
func (db *UserModel) Register(name string, address string) (*User, error) {
	var err error
	u, err := db.FindByName(name)
	if err == nil {
		return u, nil
	}
	user := NewUser()
	user.Name = name
	res, err := db.Exec(
		`INSERT INTO user SET
			name = ?,
			ethereumAddress = ?,
			createdAt = ?,
			updatedAt = ?`,
		user.Name,
		address,
		user.CreatedAt,
		user.UpdatedAt,
	)
	if err != nil {
		return nil, ErrServerError
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, ErrServerError
	}
	db.InsertBalance(ID(id), ID(1), "1000")
	return db.FindByID(ID(id))
}

// FindByID finds a user by id
func (db *UserModel) FindByID(id ID) (*User, error) {
	row := db.QueryRow(
		fmt.Sprintf("SELECT %s FROM user WHERE id=?", getUserCols()),
		id,
	)
	return scanUser(row)
}

//FindByName finds a user by email
func (db *UserModel) FindByName(name string) (*User, error) {
	row := db.QueryRow(
		fmt.Sprintf("SELECT %s FROM user WHERE name=?", getUserCols()),
		name,
	)

	return scanUser(row)
}

//IsNameRegistered checks whether email is already registered or not
func (db *UserModel) IsNameRegistered(name string) bool {
	var count int32
	err := db.QueryRow(
		"SELECT count(*) as counter FROM user WHERE name=?",
		name,
	).Scan(&count)
	return err == nil && count > 0
}

func getUserCols() string {
	return "id, name, ethereumAddress, createdAt, updatedAt"
}

func scanUser(row *sql.Row) (*User, error) {
	var user User
	err := row.Scan(
		&user.ID,
		&user.Name,
		&user.EthereumAddress,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, errors.New("Not found")
	} else if err != nil {
		logrus.WithFields(
			logrus.Fields{"e": err.Error()},
		).Error("user:scanUser:1")
		return nil, err
	}
	return &user, nil
}
