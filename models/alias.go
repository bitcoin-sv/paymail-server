package models

type AliasDetails struct {
	UserID  uint64 `json:"user_id,omitempty" db:"user_id"`
	Paymail string `json:"paymail,omitempty" db:"paymail"`
	Error   error  `json:"error,omitempty" db:"-"`
}
