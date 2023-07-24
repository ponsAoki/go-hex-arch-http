package ports

import entity "hex/internal/adapters/entities/nabeatsu"

type NabeatsuRepositoryPort interface {
	CreateNabeatsu(nabeatsu *entity.Nabeatsu) error
}
