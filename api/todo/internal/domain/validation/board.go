package validation

import (
	"context"

	"github.com/calmato/gran/api/todo/internal/domain"
)

// BoardDomainValidation - BoardDomainValidationインターフェース
type BoardDomainValidation interface {
	Board(ctx context.Context, b *domain.Board) []*domain.ValidationError
	BoardList(ctx context.Context, bl *domain.BoardList) []*domain.ValidationError
}
