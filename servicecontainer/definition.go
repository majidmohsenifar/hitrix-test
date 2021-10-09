package servicecontainer

import (
	"hitrix-test/internal/auth"
	"hitrix-test/internal/order"
	"hitrix-test/internal/product"

	"github.com/coretrix/hitrix/service"
	"github.com/coretrix/hitrix/service/component/authentication"
	"github.com/coretrix/hitrix/service/component/password"
	"github.com/latolukasz/beeorm"
	"github.com/sarulabs/di"
)

func ServiceProviderProductService() *service.DefinitionGlobal {
	return &service.DefinitionGlobal{
		Name: "product_service",
		Build: func(ctn di.Container) (interface{}, error) {
			ormEngine := ctn.Get(service.ORMEngineGlobalService).(*beeorm.Engine)
			return product.New(ormEngine), nil
		},
	}

}

func ServiceProviderAuthService() *service.DefinitionGlobal {
	return &service.DefinitionGlobal{
		Name: "auth_service",
		Build: func(ctn di.Container) (interface{}, error) {
			ormEngine := ctn.Get(service.ORMEngineGlobalService).(*beeorm.Engine)
			authenticationService := ctn.Get(service.AuthenticationService).(*authentication.Authentication)
			passwordService := ctn.Get(service.PasswordService).(password.IPassword)
			return auth.New(ormEngine, authenticationService, passwordService), nil
		},
	}
}

func ServiceProviderBasketService() *service.DefinitionGlobal {
	return &service.DefinitionGlobal{
		Name: "basket_service",
		Build: func(ctn di.Container) (interface{}, error) {
			return order.NewBasketService(), nil
		},
	}
}
