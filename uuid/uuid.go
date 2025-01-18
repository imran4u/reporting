package uuid

type Generator interface {
	Generate() string
	Parse(uuid string) error
}

func new() Generator {
	return &generator{}
}
