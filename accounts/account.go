package accounts

import (
	"github.com/oceanhq/accounts/uuid"
)

type Account struct {
	ID       uuid.UUID
	Username string
}
