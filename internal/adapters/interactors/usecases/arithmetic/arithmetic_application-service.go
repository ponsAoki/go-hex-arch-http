package usecase

import (
	usecase "hex/internal/adapters/interactors/usecases/nabeatsu"
	arith_ports "hex/internal/ports/entities"
	repository_ports "hex/internal/ports/repositories"
)

type ArithmeticApplicationService struct {
	arithRepo repository_ports.ArithmeticRepositoryPort
	nabeRepo  repository_ports.NabeatsuRepositoryPort
	arith     arith_ports.ArithmeticPort
}

func NewArithmeticApplicationService(arithRepo repository_ports.ArithmeticRepositoryPort, nabeRepo repository_ports.NabeatsuRepositoryPort, arith arith_ports.ArithmeticPort) *ArithmeticApplicationService {
	return &ArithmeticApplicationService{arithRepo, nabeRepo, arith}
}

func (adapter *ArithmeticApplicationService) GetAddition(a, b int32) (int32, error) {
	answer, err := adapter.arith.Addition(a, b)
	if err != nil {
		return 0, err
	}

	err = adapter.arithRepo.AddToHistory(a, b, answer, "addition")
	if err != nil {
		return 0, err
	}

	//ナベアツユースケース
	_, err = usecase.NewCreateNabeatsuService(adapter.nabeRepo).Handle(a, b, answer, "addition")
	if err != nil {
		return answer, err
	}

	return answer, nil
}

func (adapter *ArithmeticApplicationService) GetSubtraction(a, b int32) (int32, error) {
	answer, err := adapter.arith.Subtraction(a, b)
	if err != nil {
		return 0, err
	}

	err = adapter.arithRepo.AddToHistory(a, b, answer, "subtraction")
	if err != nil {
		return 0, err
	}

	//ナベアツユースケース
	_, err = usecase.NewCreateNabeatsuService(adapter.nabeRepo).Handle(a, b, answer, "subtraction")
	if err != nil {
		return answer, err
	}

	return answer, nil
}

func (adapter *ArithmeticApplicationService) GetMultiplication(a, b int32) (int32, error) {
	answer, err := adapter.arith.Multiplication(a, b)
	if err != nil {
		return 0, err
	}

	err = adapter.arithRepo.AddToHistory(a, b, answer, "multiplication")
	if err != nil {
		return 0, err
	}

	//ナベアツユースケース
	_, err = usecase.NewCreateNabeatsuService(adapter.nabeRepo).Handle(a, b, answer, "multiplication")
	if err != nil {
		return answer, err
	}

	return answer, nil
}

func (adapter *ArithmeticApplicationService) GetDivision(a, b int32) (int32, error) {
	answer, err := adapter.arith.Division(a, b)
	if err != nil {
		return 0, err
	}

	err = adapter.arithRepo.AddToHistory(a, b, answer, "division")
	if err != nil {
		return 0, err
	}

	//ナベアツユースケース
	_, err = usecase.NewCreateNabeatsuService(adapter.nabeRepo).Handle(a, b, answer, "division")
	if err != nil {
		return answer, err
	}

	return answer, nil
}
