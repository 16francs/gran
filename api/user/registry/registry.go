package registry

import (
	"github.com/16francs/gran/api/user/internal/application"
	rv "github.com/16francs/gran/api/user/internal/application/validation"
	"github.com/16francs/gran/api/user/internal/domain/service"
	"github.com/16francs/gran/api/user/internal/infrastructure/persistence"
	dv "github.com/16francs/gran/api/user/internal/infrastructure/validation"
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
	v1User := v1UserInjection(fa, fs)

	return &Registry{
		V1User:   v1User,
		V1Health: v1Health,
	}
}

func v1HealthInjection() handler.APIV1HealthHandler {
	hh := handler.NewAPIV1HealthHandler()

	return hh
}

func v1UserInjection(fa *authentication.Auth, fs *firestore.Firestore) v1.APIV1UserHandler {
	up := persistence.NewUserPersistence(fa, fs)
	udv := dv.NewUserDomainValidation(up)
	us := service.NewUserService(udv, up)
	urv := rv.NewUserRequestValidation()
	uu := application.NewUserApplication(urv, us)
	uh := v1.NewAPIV1UserHandler(uu)

	return uh
}
