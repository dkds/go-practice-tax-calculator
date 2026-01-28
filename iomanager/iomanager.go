package iomanager

type IOManager interface {
	ReadLines() ([]string, error)
	WriteJson(data any) error
}
