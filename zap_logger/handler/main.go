package handler

type (
	Handler interface {
		Do(param interface{}) error
	}
)
