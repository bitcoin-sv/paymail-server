package models

type AliasResponse struct {
	UserID  uint64 `json:"user_id,omitempty" db:"user_id"`
	Paymail string `json:"paymail,omitempty" db:"paymail"`
	Error   string `json:"error,omitempty" db:"-"`
}

type NewAliasDetails struct {
	Paymail     string `json:"paymail"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Avatar      string `json:"avatar,omitempty"`
	Address     string `json:"address,omitempty"`
	PhoneNumber string `json:"phoneNumber,omitempty"`
}
