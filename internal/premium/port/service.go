package port

import "context"

type IPremiumService interface {
	UpdateUserPremium(ctx context.Context, userEmail string) error
}
