package main

import (
	"github.com/kwesilarbi/store-microservices/order/config"
	"github.com/kwesilarbi/store-microservices/order/internal/adapters/db"
	"github.com/kwesilarbi/store-microservices/order/internal/adapters/grpc"
	"github.com/kwesilarbi/store-microservices/order/internal/application/core/api"
	"log"
)

func main() {
	dbAdapter, err := db.NewAdapter(config.GetDataSourceURL())
	if err != nil {
		log.Fatalf("failed to connect to database. Error: %v", err)
	}

	application := api.NewApplication(dbAdapter)
	grpcAdapter := grpc.NewAdapter(application, config.GetApplicationPort())
	grpcAdapter.Run()
}