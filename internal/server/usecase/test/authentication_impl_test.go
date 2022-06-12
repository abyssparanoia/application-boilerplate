package test

import (
	"context"
	"testing"
	"time"

	"github.com/abyssparanoia/application-boilerplate/internal/pkg/glueerr"
	"github.com/abyssparanoia/application-boilerplate/internal/pkg/now"
	"github.com/abyssparanoia/application-boilerplate/internal/server/domain/model"
	"github.com/abyssparanoia/application-boilerplate/internal/server/domain/model/testdata"
	mock_repository "github.com/abyssparanoia/application-boilerplate/internal/server/domain/repository/mock"
	"github.com/abyssparanoia/application-boilerplate/internal/server/usecase"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUsecaseAuthentication_VerifyIDToken(t *testing.T) {
	now.FakeNow(time.Date(2020, 05, 01, 10, 0, 0, 0, time.UTC))

	td := testdata.NewDomainModel()
	authUserInfo := td.AuthUserInfo

	type args struct {
		idToken string
	}

	type want struct {
		authUserInfo   *model.AuthUserInfo
		expectedResult error
	}

	tests := map[string]struct {
		args    args
		usecase func(ctx context.Context, ctrl *gomock.Controller) usecase.Authentication
		want    want
	}{
		"success": {
			args: args{
				idToken: "idToken",
			},
			usecase: func(ctx context.Context, ctrl *gomock.Controller) usecase.Authentication {
				mockAuthRepo := mock_repository.NewMockAuthentication(ctrl)
				mockAuthRepo.
					EXPECT().
					VerifyIDToken(ctx, "idToken").
					Return(authUserInfo, nil)

				return usecase.NewAuthentication(
					mockAuthRepo,
				)
			},
			want: want{
				authUserInfo: authUserInfo,
			},
		},
		"failed": {
			args: args{
				idToken: "idToken",
			},
			usecase: func(ctx context.Context, ctrl *gomock.Controller) usecase.Authentication {
				mockAuthRepo := mock_repository.NewMockAuthentication(ctrl)
				mockAuthRepo.
					EXPECT().
					VerifyIDToken(ctx, "idToken").
					Return(nil, glueerr.InternalErr.New())

				return usecase.NewAuthentication(
					mockAuthRepo,
				)
			},
			want: want{
				expectedResult: glueerr.InternalErr.New(),
			},
		},
	}

	for testName, arg := range tests {
		t.Run(testName, func(t *testing.T) {
			ctx := context.Background()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			u := arg.usecase(ctx, ctrl)

			got, err := u.Verify(ctx, arg.args.idToken)
			if arg.want.expectedResult == nil {
				assert.NoError(t, err)
				assert.Equal(t, authUserInfo, got)
			} else {
				assert.ErrorContains(t, err, arg.want.expectedResult.Error())
			}
		})
	}
}
