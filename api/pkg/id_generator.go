package pkg

type IDGenerator interface {
	Generate() (string, error)
}
