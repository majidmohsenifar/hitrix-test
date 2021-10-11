package auth

import (
	"context"
	"hitrix-test/internal/entities"
	"strings"

	"github.com/coretrix/hitrix/service/component/authentication"
	"github.com/coretrix/hitrix/service/component/password"
	"github.com/gin-gonic/gin"
	"github.com/latolukasz/beeorm"
)

const (
	MessageJwtNotFound          = "JWT Token not found"
	MessageJwtInvalidToken      = "invalid token"
	MessageJwtInvalidCredential = "invalid credential"
	UserKey                     = "user"
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

func (s *Service) GetUserFromContext(ctx context.Context) *entities.User {
	u, _ := ctx.Value(userCtxKey).(*entities.User)
	return u

}
func (s *Service) LoadUserByToken(token string) (*entities.User, error) {
	u := &entities.User{}
	_, err := s.authenticationService.VerifyAccessToken(s.ormEngine, token, u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

type contextKey struct {
	name string
}

var userCtxKey = &contextKey{"user"}

type authHeader struct {
	Token string `header:"Authorization"`
}

func (s *Service) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		h := authHeader{}
		err := c.ShouldBindHeader(&h)
		if err != nil {
			c.Next()
			return
		}
		parts := strings.Split(h.Token, " ")
		if len(parts) < 2 {
			c.Next()
			return
		}
		token := parts[1]
		loggedInUser, err := s.LoadUserByToken(token)
		if err != nil || loggedInUser == nil {
			c.Next()
			return
		}
		ctx := context.WithValue(c.Request.Context(), userCtxKey, loggedInUser)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func New(ormEngine *beeorm.Engine, authenticationService *authentication.Authentication, passwordService password.IPassword) *Service {
	return &Service{
		ormEngine:             ormEngine,
		authenticationService: authenticationService,
		passwordService:       passwordService,
	}
}
