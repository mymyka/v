package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/mymyka/v/v/builder"
	"github.com/mymyka/v/v/docs"
	v "github.com/mymyka/v/v/validators"
	"github.com/mymyka/v/v/vecho"
)

type ReristerUserReq struct {
	Username string
	Likes    int
}

type Tokens struct {
	Username string
	Likes    int
}

type LoginUserReq struct {
	Username string
	Likes    int
}

func (r ReristerUserReq) Validator(b builder.ValidationBuilder) {
	b.ForString(
		&r.Username,
		v.RequiredString(),
		v.MinLength(1),
		v.MaxLength(100),
	)
	b.ForInt(
		&r.Likes,
		v.RequiredInt(),
		v.MinInt(22),
		v.MaxInt(100),
	)
}

func main() {
	r := ReristerUserReq{
		Username: "LOL",
		Likes:    22,
	}

	b := builder.ValidationBuilderImpl{}
	r.Validator(&b)

	if err := b.Validate(); err != nil {
		fmt.Printf("Validation failed: %v\n", err)
	} else {
		fmt.Println("Validation passed!")
	}

	e := vecho.New(echo.New())
	a := e.Group("/a")
	o := a.GET("a", handler).WithDocs(
		docs.Description("ABC"),
		docs.Summary("aaa"),
		docs.OperationId("aaaa"),

		docs.Req[ReristerUserReq](),
		docs.Res[Tokens](),
	)

	fmt.Println(o.R.Description)
	fmt.Println(o.R.Method)
}

func handler(ctx echo.Context) error { return nil }
