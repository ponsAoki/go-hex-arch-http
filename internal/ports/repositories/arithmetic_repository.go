package ports

type ArithmeticRepositoryPort interface {
	AddToHistory(a, b, answer int32, operation string) error
}
