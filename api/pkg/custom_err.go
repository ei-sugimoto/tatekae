package pkg

type CustomErr struct {
	Code int
	Err  error
}

func NewCustomErr(code int, err error) *CustomErr {
	return &CustomErr{
		Code: code,
		Err:  err,
	}
}

func (c *CustomErr) Error() string {
	return c.Err.Error()
}
