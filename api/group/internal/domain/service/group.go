package service

import (
	"context"

	"golang.org/x/xerrors"

	"github.com/16francs/gran/api/group/internal/domain"
	"github.com/16francs/gran/api/group/internal/domain/repository"
	"github.com/16francs/gran/api/group/internal/domain/validation"
)

// GroupService - GroupServiceインターフェース
type GroupService interface {
	Index(ctx context.Context, u *domain.User) ([]*domain.Group, error)
	Show(ctx context.Context, groupID string) (*domain.Group, error)
	Create(ctx context.Context, u *domain.User, g *domain.Group) error
}

type groupService struct {
	GroupDomainValidation validation.GroupDomainValidation
	GroupRepository       repository.GroupRepository
}

// NewGroupService - GroupServiceの生成
func NewGroupService(gdv validation.GroupDomainValidation, gr repository.GroupRepository) GroupService {
	return &groupService{
		GroupDomainValidation: gdv,
		GroupRepository:       gr,
	}
}

func (gs *groupService) Index(ctx context.Context, u *domain.User) ([]*domain.Group, error) {
	g, err := gs.GroupRepository.Index(ctx, u)
	if err != nil {
		err = xerrors.Errorf("Failed to Domain/Repository: %w", err)
		return nil, domain.Unauthorized.New(err)
	}

	return g, nil
}

func (gs *groupService) Show(ctx context.Context, groupID string) (*domain.Group, error) {
	g, err := gs.GroupRepository.Show(ctx, groupID)
	if err != nil {
		err = xerrors.Errorf("Failed to Domain/Repository: %w", err)
		return nil, domain.Unauthorized.New(err)
	}

	return g, nil
}

func (gs *groupService) Create(ctx context.Context, u *domain.User, g *domain.Group) error {
	if ves := gs.GroupDomainValidation.Group(ctx, g); len(ves) > 0 {
		err := xerrors.New("Failed to Domain/DomainValidation")
		return domain.InvalidDomainValidation.New(err, ves...)
	}

	if err := gs.GroupRepository.Create(ctx, u, g); err != nil {
		err = xerrors.Errorf("Failed to Domain/Repository: %w", err)
		return domain.ErrorInDatastore.New(err)
	}

	return nil
}
