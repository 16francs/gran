package registry

import (
	"github.com/16francs/gran/api/user/internal/application"
	"github.com/16francs/gran/api/user/internal/application/validation"
	"github.com/16francs/gran/api/user/internal/infrastructure/persistence"
	"github.com/16francs/gran/api/user/internal/interface/handler"
	v1 "github.com/16francs/gran/api/user/internal/interface/handler/v1"
	"github.com/16francs/gran/api/user/lib/firebase/authentication"
	"github.com/16francs/gran/api/user/lib/firebase/firestore"
)

// Registry - DIコンテナ
type Registry struct {
	V1Health handler.APIV1HealthHandler
	V1User   v1.APIV1UserHandler
}

// NewRegistry - internalディレクトリのファイルを読み込み
func NewRegistry(fa *authentication.Auth, fs *firestore.Firestore) *Registry {
	v1Health := v1HealthInjection()
	v1User := v1UserInjection(*fa, *fs)

	return &Registry{
		V1User:   v1User,
		V1Health: v1Health,
	}
}

func v1HealthInjection() handler.APIV1HealthHandler {
	hh := handler.NewAPIV1HealthHandler()

	return hh
}

func v1UserInjection(fa authentication.Auth, fs firestore.Firestore) v1.APIV1UserHandler {
	uv := validation.NewUserValidation()
	up := persistence.NewUserPersistence(fa, fs)
	uu := application.NewUserApplication(uv, up)
	uh := v1.NewAPIV1UserHandler(uu)

	return uh
}
