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

func TestUsecaseUser_Get(t *testing.T) {
	now.FakeNow(time.Date(2020, 05, 01, 10, 0, 0, 0, time.UTC))

	td := testdata.NewDomainModel()
	user := td.User

	type args struct {
		id string
	}

	type want struct {
		user           *model.User
		expectedResult error
	}

	tests := map[string]struct {
		args    args
		usecase func(ctx context.Context, ctrl *gomock.Controller) usecase.User
		want    want
	}{
		"success": {
			args: args{
				id: user.ID,
			},
			usecase: func(ctx context.Context, ctrl *gomock.Controller) usecase.User {
				mockUserRepo := mock_repository.NewMockUser(ctrl)
				mockUserRepo.
					EXPECT().
					Get(ctx, user.ID).
					Return(user, nil)

				return usecase.NewUser(
					mockUserRepo,
				)
			},
			want: want{
				user: user,
			},
		},
		"not found": {
			args: args{
				id: user.ID,
			},
			usecase: func(ctx context.Context, ctrl *gomock.Controller) usecase.User {
				mockUserRepo := mock_repository.NewMockUser(ctrl)
				mockUserRepo.
					EXPECT().
					Get(ctx, user.ID).
					Return(nil, glueerr.NotFoundErr.New())

				return usecase.NewUser(
					mockUserRepo,
				)
			},
			want: want{
				expectedResult: glueerr.NotFoundErr.New(),
			},
		},
	}

	for testName, arg := range tests {
		t.Run(testName, func(t *testing.T) {
			ctx := context.Background()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			u := arg.usecase(ctx, ctrl)

			got, err := u.Get(ctx, arg.args.id)
			if arg.want.expectedResult == nil {
				assert.NoError(t, err)
				assert.Equal(t, user, got)
			} else {
				assert.ErrorContains(t, err, arg.want.expectedResult.Error())
			}
		})
	}
}
