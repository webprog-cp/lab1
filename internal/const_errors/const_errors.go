package const_errors

type ConstError string

func (e ConstError) Error() string {
	return string(e)
}

const NoEntityByID = ConstError("No such entity with this ID found")
