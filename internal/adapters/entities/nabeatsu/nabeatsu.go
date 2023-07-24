package entity

type Nabeatsu struct {
	A, B, Answer int32
	Operation    string
}

func NewNabeatsu(a, b, answer int32, operation string) *Nabeatsu {
	return &Nabeatsu{A: a, B: b, Answer: answer, Operation: operation}
}
