package vecho

func Description(d string) func() {
	return func() {}
}

func Summary(s string) func() { return func() {} }

func OperationId(i string) func() { return func() {} }

func Req[T any]() func() {
	return func() {}
}

func Res[T any]() func() {
	return func() {}
}

func Err[T any](code int) func() {
	return func() {}
}
