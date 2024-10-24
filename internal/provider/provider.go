package provider

type NumberProvider interface {
	Load() ([]int, error)
}
