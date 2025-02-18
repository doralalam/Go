package iomanager

type IOmanager interface {
	ReadingFromFile() ([]string, error)
	WriteJSON(data any) error
}
