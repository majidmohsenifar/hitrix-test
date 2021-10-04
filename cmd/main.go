package main

import (
	"fmt"
	"hitrix-test/graph"
	"hitrix-test/graph/generated"
	"hitrix-test/internal/entities"
	"hitrix-test/servicecontainer"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/coretrix/hitrix"
	"github.com/coretrix/hitrix/service/registry"
	"github.com/gin-gonic/gin"
	"github.com/latolukasz/beeorm"
)

func main() {
	s, deferFunc := hitrix.New(
		"tiny-store", "someSecretThatShouldBeChangedLater",
	).RegisterDIGlobalService(
		registry.ServiceProviderErrorLogger(),
		registry.ServiceProviderConfigDirectory("./config"), //register config service. As param you should point to the folder of your config file
		registry.ServiceProviderOrmRegistry(RegisterEntities),
		registry.ServiceProviderOrmEngine(),
		registry.ServiceProviderJWT(),      //register JWT DI service
		registry.ServiceProviderPassword(), //register pasword DI service
		servicecontainer.ServiceProviderProductService(),
	).RegisterDIRequestService(
		registry.ServiceProviderOrmEngineForContext(true),
	).Build()
	defer deferFunc()
	updateSchema()
	productService := servicecontainer.ProductService()
	graphqlResolver := graph.NewResolver(productService)
	s.RunServer(9999, generated.NewExecutableSchema(generated.Config{Resolvers: graphqlResolver}), func(ginEngine *gin.Engine) {
		//TODO register middleware here
	}, func(server *handler.Server) {

	})
}

func RegisterEntities(registry *beeorm.Registry) {
	registry.RegisterEntity(&entities.Product{})

}

func updateSchema() {
	engine := servicecontainer.ORMEngine()
	alters := engine.GetAlters()
	for _, alter := range alters {
		fmt.Println("alter", alter.SQL)
		alter.Exec()
	}
}
