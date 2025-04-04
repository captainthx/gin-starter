package service

import (
	errs "gin-starter/common/err"
	"gin-starter/core/domain"
	"gin-starter/core/dto"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type mockAuthRepository struct {
	mock.Mock
}

// Mock สำหรับ FindByEmail
func (m *mockAuthRepository) FindByEmail(email string) (*domain.User, error) {
	args := m.Called(email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.User), args.Error(1)
}

// Mock สำหรับ CreateUser
func (m *mockAuthRepository) CreateUser(user *domain.User) (*domain.User, error) {
	args := m.Called(user)
	return args.Get(0).(*domain.User), args.Error(1)
}

// Mock สำหรับ ExistByEmail
func (m *mockAuthRepository) ExistByEmail(email string) bool {
	args := m.Called(email)
	return args.Bool(0)
}

func Test_authService_Login(t *testing.T) {

	hashPassword, _ := bcrypt.GenerateFromPassword([]byte("12345678"), bcrypt.DefaultCost)

	type args struct {
		request *dto.LoginRequest
	}

	// mock user
	mockUser := &domain.User{
		Email:    "user@example.com",
		Password: string(hashPassword),
	}

	tests := []struct {
		name      string
		mockSetup func(repo *mockAuthRepository)
		args      args
		want      *dto.TokenResponse
		wantErr   error
	}{
		{
			name: "login success",
			args: args{
				&dto.LoginRequest{
					Email:    "user@example.com",
					Password: "12345678",
				},
			},
			mockSetup: func(repo *mockAuthRepository) {
				repo.On("FindByEmail", "user@example.com").Return(mockUser, nil)
			},
			want:    &dto.TokenResponse{AccessToken: "mock-token", RefreshToken: "mock-token"},
			wantErr: nil,
		},
		{
			name: "login fail not found user",
			args: args{
				&dto.LoginRequest{
					Email:    "user1@example.com",
					Password: "12345678",
				},
			},
			mockSetup: func(repo *mockAuthRepository) {
				repo.On("FindByEmail", "user1@example.com").Return(nil, gorm.ErrRecordNotFound)
			},
			want:    nil,
			wantErr: errs.NewNotFoundError("invalid credentials"),
		},
		{
			name: "login fail invalid password",
			args: args{
				&dto.LoginRequest{
					Email:    "user@example.com",
					Password: "1234567",
				},
			},
			mockSetup: func(repo *mockAuthRepository) {
				repo.On("FindByEmail", "user@example.com").Return(mockUser, nil)
			},
			want:    nil,
			wantErr: errs.NewBadRequestError("invalid credentials"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(mockAuthRepository)

			if tt.mockSetup != nil {
				tt.mockSetup(mockRepo)
			}

			service := NewAuthService(mockRepo)
			got, err := service.Login(tt.args.request)

			if tt.wantErr != nil {
				assert.EqualError(t, err, "invalid credentials")
				assert.ErrorIs(t, err, tt.wantErr)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, got)
			}
			defer mockRepo.AssertExpectations(t)
		})
	}
}

func Test_authService_SignUp(t *testing.T) {
	type args struct {
		request *dto.SignUpRequest
	}

	tests := []struct {
		name      string
		mockSetup func(repo *mockAuthRepository)
		args      args
		want      *dto.TokenResponse
		wantErr   error
	}{
		{
			name: "signUp success",
			args: args{
				&dto.SignUpRequest{
					Email:    "test@example.com",
					Password: "12345678",
				},
			},
			mockSetup: func(repo *mockAuthRepository) {
				repo.On("ExistByEmail", "test@example.com").Return(false)
				repo.On("CreateUser", mock.MatchedBy(func(user *domain.User) bool {
					return user.Email == "test@example.com" && bcrypt.CompareHashAndPassword([]byte(user.Password), []byte("12345678")) == nil
				})).Return(&domain.User{Email: "test@example.com", Password: "hashed-password"}, nil)
			},
			want:    &dto.TokenResponse{AccessToken: "mock-token", RefreshToken: "mock-token"},
			wantErr: nil,
		},
		{
			name: "signUp fail to create user",
			args: args{
				&dto.SignUpRequest{
					Email:    "test2@example.com",
					Password: "12345678",
				},
			},
			mockSetup: func(repo *mockAuthRepository) {
				repo.On("ExistByEmail", "test2@example.com").Return(false)
				repo.On("CreateUser", mock.Anything).Return((*domain.User)(nil), errs.NewUnexpectedError("unknow error"))
			},
			want:    nil,
			wantErr: errs.NewUnexpectedError("unknow error"),
		},
		{
			name: "signUp fail email already exists",
			args: args{
				&dto.SignUpRequest{
					Email:    "test@example.com",
					Password: "12345678",
				},
			},
			mockSetup: func(repo *mockAuthRepository) {
				repo.On("ExistByEmail", "test@example.com").Return(true)
			},
			want:    nil,
			wantErr: errs.NewBadRequestError("email already exists"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(mockAuthRepository)

			if tt.mockSetup != nil {
				tt.mockSetup(mockRepo)
			}

			service := NewAuthService(mockRepo)

			got, err := service.SignUp(tt.args.request)

			if tt.wantErr != nil {
				assert.ErrorIs(t, err, tt.wantErr)
				assert.Nil(t, got)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, got)
			}
			defer mockRepo.AssertExpectations(t)
		})
	}
}
