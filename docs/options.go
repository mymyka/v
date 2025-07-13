package docs

func Description(d string) func(*Route) {
	return func(r *Route) {
		r.Description = d
	}
}

func Summary(s string) func(*Route) {
	return func(r *Route) {
		r.Summary = s
	}
}

func OperationId(o string) func(*Route) {
	return func(r *Route) {
		r.OperationId = o
	}
}

func Req[T any]() func(*Route) {
	return func(r *Route) {}
}

func Res[T any]() func(*Route) {
	return func(r *Route) {}
}

func Err[T any](code int) func(*Route) {
	return func(r *Route) {}
}
