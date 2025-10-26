package models

type HttpResponse[T any] struct {
	Url          string
	Request      string
	Response     *string
	ResponseCode *int
	Data         *T
}
