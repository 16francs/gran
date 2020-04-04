package application

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/16francs/gran/api/user/internal/application/request"
	"github.com/16francs/gran/api/user/internal/domain"
	mock_validation "github.com/16francs/gran/api/user/mock/application/validation"
	mock_service "github.com/16francs/gran/api/user/mock/domain/service"
)

func TestUserApplication_Create(t *testing.T) {
	testCases := map[string]struct {
		Request *request.CreateUser
	}{
		"ok": {
			Request: &request.CreateUser{
				Email:                "hoge@hoge.com",
				Password:             "12345678",
				PasswordConfirmation: "12345678",
			},
		},
	}

	for result, testCase := range testCases {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		// Defined variables
		ves := make([]*domain.ValidationError, 0)

		u := &domain.User{
			Email:    testCase.Request.Email,
			Password: testCase.Request.Password,
		}

		// Defined mocks
		urvm := mock_validation.NewMockUserRequestValidation(ctrl)
		urvm.EXPECT().CreateUser(testCase.Request).Return(ves)

		usm := mock_service.NewMockUserService(ctrl)
		usm.EXPECT().Create(ctx, u).Return(u, nil)

		// Start test
		t.Run(result, func(t *testing.T) {
			target := NewUserApplication(urvm, usm)

			err := target.Create(ctx, testCase.Request)
			if err != nil {
				t.Fatalf("error: %v", err)
				return
			}
		})
	}
}

func TestUserApplication_UpdateProfile(t *testing.T) {
	testCases := map[string]struct {
		Request *request.UpdateProfile
	}{
		"ok": {
			Request: &request.UpdateProfile{
				Name:        "テストユーザー",
				DisplayName: "ユーザー",
				Email:       "hoge@hoge.com",
				PhoneNumber: "",
				Thumbnail:   "",
				Biography:   "",
			},
		},
	}

	for result, testCase := range testCases {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		// Defined variables
		ves := make([]*domain.ValidationError, 0)

		u := &domain.User{
			Name:        testCase.Request.Name,
			DisplayName: testCase.Request.DisplayName,
			Email:       testCase.Request.Email,
			PhoneNumber: testCase.Request.PhoneNumber,
			Biography:   testCase.Request.Biography,
		}

		// Defined mocks
		urvm := mock_validation.NewMockUserRequestValidation(ctrl)
		urvm.EXPECT().UpdateProfile(testCase.Request).Return(ves)

		usm := mock_service.NewMockUserService(ctrl)
		usm.EXPECT().Authentication(ctx).Return(u, nil)
		usm.EXPECT().Update(ctx, u).Return(u, nil)

		// Start test
		t.Run(result, func(t *testing.T) {
			target := NewUserApplication(urvm, usm)

			err := target.UpdateProfile(ctx, testCase.Request)
			if err != nil {
				t.Fatalf("error: %v", err)
				return
			}
		})
	}
}
