package port

import "context"

type IPremiumRepo interface {
	UpdateUserPremium(ctx context.Context, userID string) error
}
