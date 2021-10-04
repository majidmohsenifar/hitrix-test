package servicecontainer

import (
	"hitrix-test/internal/product"

	"github.com/coretrix/hitrix/service"
	"github.com/latolukasz/beeorm"
)

func ProductService() *product.Service {
	return service.GetServiceRequired("product_service").(*product.Service)
}

func ORMEngine() *beeorm.Engine {
	return service.GetServiceRequired(service.ORMEngineGlobalService).(*beeorm.Engine)
}
