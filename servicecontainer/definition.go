package servicecontainer

import (
	"hitrix-test/internal/product"

	"github.com/coretrix/hitrix/service"
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
