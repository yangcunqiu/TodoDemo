package model

type Response[T interface{}] struct {
	Code    int
	Message string
	Data    T
}
