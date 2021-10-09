package auth

import (
	"hitrix-test/internal/entities"

	"github.com/coretrix/hitrix/service/component/authentication"
	"github.com/coretrix/hitrix/service/component/password"
	"github.com/latolukasz/beeorm"
)

type Service struct {
	ormEngine             *beeorm.Engine
	authenticationService *authentication.Authentication
	passwordService       password.IPassword
}

type RegisterParams struct {
	Email    string
	Password string
}

type LoginParams struct {
	Email    string
	Password string
}

func (s *Service) Register(params RegisterParams) error {
	hashedPassword, err := s.passwordService.HashPassword(params.Password)
	if err != nil {
		return err
	}
	u := &entities.User{
		Email:    params.Email,
		Password: hashedPassword,
	}
	err = s.ormEngine.FlushWithCheck(u)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Login(params LoginParams) (accessToken, refreshToken string, err error) {
	u := &entities.User{}
	accessToken, refreshToken, err = s.authenticationService.Authenticate(s.ormEngine, params.Email, params.Password, u)
	if err != nil {
		return
	}
	return accessToken, refreshToken, nil
}

func New(ormEngine *beeorm.Engine, authenticationService *authentication.Authentication, passwordService password.IPassword) *Service {
	return &Service{
		ormEngine:             ormEngine,
		authenticationService: authenticationService,
		passwordService:       passwordService,
	}
}
