package main

import (
	"hex/internal/adapters/app/api"
	"hex/internal/adapters/core/arithmetic"
	"hex/internal/adapters/framework/right/db"
	"hex/internal/ports"
	"log"
	"os"

	gRPC "hex/internal/adapters/framework/left/grpc"
)

func main() {
	var err error

	//ports
	var dbAdapter ports.DbPort
	var core ports.ArithmeticPort
	var appAdapter ports.APIPort
	var gRPCAdapter ports.GrpcPort

	dbDriver := os.Getenv("DB_DRIVER")
	dataSourceName := os.Getenv("DATA_SOURCE_NAME")

	dbAdapter, err = db.NewAdapter(dbDriver, dataSourceName)
	if err != nil {
		log.Fatalf("failed to initiate database connection: %v", err)
	}
	defer dbAdapter.CloseDbConnection()

	core = arithmetic.NewAdapter()

	appAdapter = api.NewAdapter(dbAdapter, core)

	gRPCAdapter = gRPC.NewAdapter(appAdapter)
	// gRPCAdapter.Run()
}
