package usecase

import (
	entity "hex/internal/adapters/entities/nabeatsu"
	repository_ports "hex/internal/ports/repositories"
	"strconv"
	"strings"
)

type CreateNabeatsuService struct {
	repo repository_ports.NabeatsuRepositoryPort
}

func NewCreateNabeatsuService(repo repository_ports.NabeatsuRepositoryPort) *CreateNabeatsuService {
	return &CreateNabeatsuService{repo}
}

func (adapter *CreateNabeatsuService) Handle(a, b, answer int32, operation string) (*entity.Nabeatsu, error) {
	okNabeatsu := false
	nums := []int32{a, b, answer}
	for _, num := range nums {
		if nabeatsuChecker(num) {
			okNabeatsu = true
		}
		if okNabeatsu {
			break
		}
	}
	if !okNabeatsu {
		return nil, nil
	}

	nabeatsu := entity.NewNabeatsu(a, b, answer, operation)

	err := adapter.repo.CreateNabeatsu(nabeatsu)
	if err != nil {
		return nil, err
	}

	return nabeatsu, nil
}

func nabeatsuChecker(num int32) bool {
	containThree := strings.Contains(strconv.Itoa(int(num)), "3")

	if containThree || num%3 == 0 {
		return true
	}

	return false
}
