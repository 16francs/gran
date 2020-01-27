package persistence

import (
	"context"
	"strings"

	"github.com/16francs/gran/api/user/internal/domain"
	"github.com/16francs/gran/api/user/internal/domain/repository"
	"github.com/16francs/gran/api/user/lib/firebase/firestore"
)

type groupPersistence struct {
	firestore *firestore.Firestore
}

// NewGroupPersistence - GroupRepositoryの生成
func NewGroupPersistence(fs *firestore.Firestore) repository.GroupRepository {
	return &groupPersistence{
		firestore: fs,
	}
}

func (gp *groupPersistence) Create(ctx context.Context, u *domain.User, g *domain.Group) error {
	g.UserRefs = append(g.UserRefs, getUserReference(u.ID))

	_, err := gp.firestore.Add(ctx, GroupCollection, g)
	if err != nil {
		return err
	}

	return nil
}

func getUserReference(userID string) string {
	return strings.Join([]string{GroupCollection, userID}, "/")
}
