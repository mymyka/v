package docs

import "reflect"

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

func Req[T Validatable]() func(*Route) {
	return func(r *Route) {
		var req T

		b := ValidationBuilderImpl{}
		req.Validator(&b)

		var BodyProps []DocsProperty
		var QueryProps []DocsProperty
		var PathProps []DocsProperty

		for _, f := range b.stringFields {
			fieldPointer := f.field
			rules := f.rules

			tags := getFieldTagsFromPointer(&req, fieldPointer)

			for loc, name := range tags {
				if name == "" {
					continue
				}

				docsProp := DocsProperty{
					Name:     name,
					Location: loc,
					Type:     "int",
				}

				for _, r := range rules {
					r.Docs(&docsProp)
				}

				switch loc {
				case "path":
					PathProps = append(PathProps, docsProp)
				case "json":
					BodyProps = append(BodyProps, docsProp)
				case "form":
					QueryProps = append(QueryProps, docsProp)
				}
			}
		}

		r.Request.Body = BodyProps
		r.Request.Query = QueryProps
		r.Request.Path = PathProps
	}
}

func Res[T any]() func(*Route) {
	return func(r *Route) {
	}
}

func Err[T any](code int) func(*Route) {
	return func(r *Route) {}
}

func getFieldTagsFromPointer(structValue interface{}, fieldPtr interface{}) map[string]string {
	structVal := reflect.ValueOf(structValue)
	structType := structVal.Type()

	// Handle pointer to struct
	if structType.Kind() == reflect.Ptr {
		structVal = structVal.Elem()
		structType = structVal.Type()
	}

	fieldPtrVal := reflect.ValueOf(fieldPtr)
	if fieldPtrVal.Kind() != reflect.Ptr {
		return nil
	}

	fieldAddr := fieldPtrVal.Pointer()

	// Iterate through struct fields
	for i := 0; i < structVal.NumField(); i++ {
		field := structVal.Field(i)
		if field.CanAddr() {
			// The key fix: compare the actual addresses correctly
			if field.Addr().Pointer() == fieldAddr {
				structField := structType.Field(i)
				return map[string]string{
					"param": structField.Tag.Get("param"),
					"json":  structField.Tag.Get("json"),
					"query": structField.Tag.Get("query"),
					"form":  structField.Tag.Get("form"),
				}
			}
		}
	}

	return nil
}
