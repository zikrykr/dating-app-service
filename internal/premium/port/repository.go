package port

import "context"

type IPremiumRepo interface {
	UpdateUserPremium(ctx context.Context, userEmail string) error
}
