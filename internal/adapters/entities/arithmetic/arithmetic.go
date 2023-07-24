package entity

type Arithmetic struct{}

func NewArithmetic() *Arithmetic {
	return &Arithmetic{}
}

func (arith Arithmetic) Addition(a, b int32) (int32, error) {
	return a + b, nil
}

func (arith Arithmetic) Subtraction(a, b int32) (int32, error) {
	return a - b, nil
}

func (arith Arithmetic) Multiplication(a, b int32) (int32, error) {
	return a * b, nil
}

func (arith Arithmetic) Division(a, b int32) (int32, error) {
	return a / b, nil
}
