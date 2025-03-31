package service

import (
	"gin-starter/core/dto"
	"reflect"
	"testing"
)

func Test_authService_Login(t *testing.T) {
	type args struct {
		request *dto.LoginRequest
	}
	tests := []struct {
		name    string
		a       *authService
		args    args
		want    *dto.TokenResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.a.Login(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("authService.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("authService.Login() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_authService_SignUp(t *testing.T) {
	type args struct {
		request *dto.SignUpRequest
	}
	tests := []struct {
		name    string
		a       *authService
		args    args
		want    *dto.TokenResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.a.SignUp(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("authService.SignUp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("authService.SignUp() = %v, want %v", got, tt.want)
			}
		})
	}
}
