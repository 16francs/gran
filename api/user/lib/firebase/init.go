package firebase

import (
	"context"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

// Firebase - Firebaseの構造体
type Firebase struct {
	App *firebase.App
}

// InitializeApp - Cloud Firebase SDKの初期化
func InitializeApp(ctx context.Context, config *firebase.Config, opts ...option.ClientOption) (*Firebase, error) {
	app, err := firebase.NewApp(ctx, config, opts...)
	if err != nil {
		return nil, err
	}

	return &Firebase{app}, nil
}
