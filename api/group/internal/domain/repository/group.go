package repository

import (
	"context"

	"github.com/16francs/gran/api/group/internal/domain"
)

// GroupRepository - GroupRepositoryインターフェース
type GroupRepository interface {
	Index(ctx context.Context, u *domain.User) ([]*domain.Group, error)
	Show(ctx context.Context, groupID string) (*domain.Group, error)
	Create(ctx context.Context, u *domain.User, g *domain.Group) error
	Update(ctx context.Context, g *domain.Group) error
	InviteUser(ctx context.Context, u *domain.Group, g *domain.Group) error
	ExistUserIDInUserRefs(ctx context.Context, userID string, g *domain.Group) bool
	ExistUserIDInInvitedUserRefs(ctx context.Context, userID string, g *domain.Group) bool
}
