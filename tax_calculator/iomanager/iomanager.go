package iomanager 

type IOManager interface {
	ReadFile()	([]string, error)
	WriteJSONFile(data interface{}) error 
}