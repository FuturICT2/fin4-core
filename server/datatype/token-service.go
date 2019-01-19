package datatype

//TokenService defines asset service interface
type TokenService interface {
	FindAll(userID ID) ([]Token, error)
	InsertToken(
		userID ID,
		name string,
		symbol string,
		purpose string,
		hardCap string,
		blockchainAddress string,
		txAddress string,
	) (*Token, error)
	ToggleTokenLike(
		userID ID,
		tokenID ID,
	) error
	FindTokenForUser(
		userID ID,
		id ID,
	) (*Token, error)
	InsertClaim(
		userID ID,
		proposal string,
		actionID ID,
		logoPath string,
	) error
	FindClaim(id ID) (*Claim, error)
	ApproveClaim(
		claimID ID,
		approverID ID,
	) error
}
