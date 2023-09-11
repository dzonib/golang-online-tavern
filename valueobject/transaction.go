package valueobject

import (
	"time"

	"github.com/google/uuid"
)

// Transaction is a value object because it has no identifier and is immutable
type Transaction struct {
	// we are using lowercase, we don't want other subdomains to reach and change the values
	amount int
	// money goes from (user)
	from uuid.UUID
	// money goes to (user)
	to        uuid.UUID
	createdAt time.Time
}
