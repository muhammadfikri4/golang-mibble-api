package helper

func ReturnPoiter[T any](val T) *T {
	return &val
}
