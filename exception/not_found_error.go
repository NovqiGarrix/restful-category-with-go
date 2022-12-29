package exception

type NotFoundException struct {
	Error string
}

func NewNotFoundException(error error) NotFoundException {
	return NotFoundException{Error: error.Error()}
}
