package errors

type Error string

func (e Error) Error() string {
	return string(e)
}

const NoEntityByID = Error("No such entity with this ID found")
