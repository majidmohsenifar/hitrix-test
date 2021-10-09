package servicecontainer

import (
	"hitrix-test/internal/auth"
	"hitrix-test/internal/order"
	"hitrix-test/internal/product"

	"github.com/coretrix/hitrix/service"
	"github.com/coretrix/hitrix/service/component/jwt"
	"github.com/latolukasz/beeorm"
)

func ProductService() *product.Service {
	return service.GetServiceRequired("product_service").(*product.Service)
}

func AuthService() *auth.Service {
	return service.GetServiceRequired("auth_service").(*auth.Service)
}

func BasketService() *order.BasketService {
	return service.GetServiceRequired("basket_service").(*order.BasketService)
}

func ORMEngine() *beeorm.Engine {
	return service.GetServiceRequired(service.ORMEngineGlobalService).(*beeorm.Engine)
}

func JWTService() *jwt.JWT {
	return service.GetServiceRequired(service.JWTService).(*jwt.JWT)
}
