package main

import (
	"context"
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
		registry.ServiceProviderJWT(),
		registry.ServiceProviderPassword(),
		registry.ServiceProviderClock(),
		registry.ServiceProviderGenerator(),
		registry.ServiceProviderAuthentication(),
		registry.ServiceProviderUUID(),

		servicecontainer.ServiceProviderProductService(),
		servicecontainer.ServiceProviderAuthService(),
		servicecontainer.ServiceProviderBasketService(),
	).RegisterDIRequestService(
		registry.ServiceProviderOrmEngineForContext(true),
	).Build()
	defer deferFunc()
	updateSchema()
	runBackgroundConsumer()
	updateRedisIndex()
	productService := servicecontainer.ProductService()
	authService := servicecontainer.AuthService()
	basketService := servicecontainer.BasketService()
	graphqlResolver := graph.NewResolver(productService, authService, basketService)

	s.RunServer(9999, generated.NewExecutableSchema(generated.Config{Resolvers: graphqlResolver}), func(ginEngine *gin.Engine) {
		ginEngine.Use(authService.AuthMiddleware())
	}, func(server *handler.Server) {

	})
}

func RegisterEntities(registry *beeorm.Registry) {
	registry.RegisterEntity(&entities.Product{})
	registry.RegisterEntity(&entities.User{})
}

func updateSchema() {
	engine := servicecontainer.ORMEngine()
	alters := engine.GetAlters()
	for _, alter := range alters {
		fmt.Println("alter", alter.SQL)
		alter.Exec()
	}
}

func runBackgroundConsumer() {
	engine := servicecontainer.ORMEngine()
	consumer := beeorm.NewBackgroundConsumer(engine)
	go consumer.Digest(context.Background())
}

func updateRedisIndex() {
	engine := servicecontainer.ORMEngine()
	alters := engine.GetRedisSearchIndexAlters()
	for _, alter := range alters {
		fmt.Println("alter", alter)
		alter.Execute()
	}
	go engine.GetRedisSearch("search").ForceReindex("entities.User")
	go engine.GetRedisSearch("search").ForceReindex("entities.Product")
}
