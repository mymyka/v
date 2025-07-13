package main

import (
	"fmt"

	"github.com/mymyka/v/v"
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

func (r ReristerUserReq) Validator(b v.ValidationBuilder) {
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

	b := v.ValidationBuilderImpl{}
	r.Validator(&b)

	if err := b.Validate(); err != nil {
		fmt.Printf("Validation failed: %v\n", err)
	} else {
		fmt.Println("Validation passed!")
	}

	// e := vecho.New(echo.New())
	// e.Docs.Info.Title = "a"
	// a := e.Group("/a")
	// a.GET("a", handler).WithDocs(
	// 	vecho.Description("aaaa"),
	// 	vecho.Summary("aaa"),
	// 	vecho.OperationId("aaaa"),

	// 	vecho.Req[ReristerUserReq](),
	// 	vecho.Res[Tokens](),
	// )
}
