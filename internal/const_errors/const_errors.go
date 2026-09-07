package const_errors

type ConstError string

func (e ConstError) Error() string {
	return string(e)
}

const NoEntityByID = ConstError("no entity with this ID found")
