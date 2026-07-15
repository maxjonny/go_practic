package service

import (
	"context"
	"errors"
	"testing"

	"main/internal/models"
	"main/internal/service/mocks"

	"go.uber.org/mock/gomock"
)

//~/go/bin/mockgen -source=pkg/BasicRepository.go -destination=internal/service/mocks/mock_basic_repo.go -package=mocks

func TestGetUserCount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockRepo(ctrl)

	tests := []struct {
		name      string
		device    string
		setupMock func()
		wantCount int
		wantErr   bool
	}{
		{
			name:   "нет узлов — возвращаем 0",
			device: "empty_device",
			setupMock: func() {
				mockRepo.Device.EXPECT().
					GetActiveNode(gomock.Any(), "empty_device").
					Return([]string{}, nil)
			},
			wantCount: 0,
			wantErr:   false,
		},
		{
			name:   "ошибка при получении узлов",
			device: "bad_device",
			setupMock: func() {
				mockRepo.Device.EXPECT().
					GetActiveNode(gomock.Any(), "bad_device").
					Return(nil, errors.New("db connection refused"))
			},
			wantCount: 0,
			wantErr:   true,
		},
		{
			name:   "узлы есть, но пользователей нет",
			device: "Z5",
			setupMock: func() {
				mockRepo.Device.EXPECT().
					GetActiveNode(gomock.Any(), "Z5").
					Return([]string{"10", "20"}, nil)

				mockRepo.User.EXPECT().
					GetUsersByNodes(gomock.Any(), []string{"10", "20"}).
					Return([]models.UserCard{}, nil)
			},
			wantCount: 0,
			wantErr:   false,
		},
		{
			name:   "пользователи найдены — кеш обновлён",
			device: "Z5",
			setupMock: func() {
				mockRepo.Device.EXPECT().
					GetActiveNode(gomock.Any(), "Z5").
					Return([]string{"10"}, nil)

				mockRepo.User.EXPECT().
					GetUsersByNodes(gomock.Any(), []string{"10"}).
					Return([]models.UserCard{
						{GID: "gid1", Name: "Иван"},
						{GID: "gid2", Name: "Мария"},
					}, nil)

				mockRepo.User.EXPECT().
					DropCache(gomock.Any(), "Z5")

				mockRepo.User.EXPECT().
					CreateCache(gomock.Any(), "Z5", gomock.Any())
			},
			wantCount: 2,
			wantErr:   false,
		},
		{
			name:   "ошибка при получении пользователей",
			device: "Z5",
			setupMock: func() {
				mockRepo.Device.EXPECT().
					GetActiveNode(gomock.Any(), "Z5").
					Return([]string{"10"}, nil)

				mockRepo.User.EXPECT().
					GetUsersByNodes(gomock.Any(), []string{"10"}).
					Return(nil, errors.New("query timeout"))
			},
			wantCount: 0,
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMock()

			svc := InitService(&mockRepo)

			count, err := svc.GetUserCount(context.Background(), tt.device)

			if tt.wantErr && err == nil {
				t.Error("ожидалась ошибка, но её нет")
			}
			if !tt.wantErr && err != nil {
				t.Errorf("ошибка не ожидалась, но получена: %v", err)
			}
			if count != tt.wantCount {
				t.Errorf("count = %d, want %d", count, tt.wantCount)
			}
		})
	}
}
