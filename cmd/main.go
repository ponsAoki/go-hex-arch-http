package main

import (
	entity "hex/internal/adapters/entities/arithmetic"
	"hex/internal/adapters/frameworks/right/db"
	usecase "hex/internal/adapters/interactors/usecases/arithmetic"
	repository "hex/internal/adapters/repositories"
	"hex/internal/config"
	entity_ports "hex/internal/ports/entities"
	framework_ports "hex/internal/ports/frameworks"
	interactor_ports "hex/internal/ports/interactors"
	repository_ports "hex/internal/ports/repositories"
	"log"

	HTTP "hex/internal/adapters/frameworks/left/http"
)

func main() {
	var err error

	//ports
	var arith entity_ports.ArithmeticPort
	var arithRepoAdapter repository_ports.ArithmeticRepositoryPort
	var nabeRepoAdapter repository_ports.NabeatsuRepositoryPort
	var arithUcAdapter interactor_ports.ArithmeticApplicationServicePort
	var httpAdapter framework_ports.HttpPort

	dbAdapter, err := db.NewAdapter(config.Env.DB_DRIVER, config.Env.DATA_SOURCE_NAME)
	if err != nil {
		log.Fatalf("failed to initiate database connection: %v", err)
	}
	defer dbAdapter.CloseDbConnection()

	arith = entity.NewArithmetic()

	arithRepoAdapter = repository.NewArithmeticRepository(dbAdapter.DB)
	nabeRepoAdapter = repository.NewNabeatsuRepository(dbAdapter.DB)

	arithUcAdapter = usecase.NewArithmeticApplicationService(arithRepoAdapter, nabeRepoAdapter, arith)

	httpAdapter = HTTP.NewAdapter(arithUcAdapter)
	httpAdapter.Run()
}
