package main

import (
	"fmt"
	"reflect"
)

// Simplified structs for demonstration
type ReristerUserReq struct {
	Username string `param:"usernameParam" query:"userameQuery" form:"usernameForm" json:"usernameJson"`
	Likes    int    `param:"likesParam" query:"likesQuery" form:"likesForm" json:"likesJson"`
	Lol      Lol    `json:"lolJSON"`
	A        int
}

type Lol struct {
	A int
}

func main() {
	u := ReristerUserReq{
		Username: "a",
		Likes:    25,
	}

	// Method 1: Using field pointer (your original approach, but fixed)
	fmt.Println("=== Method 1: Using field pointer ===")
	up := &u.Lol
	tags := getFieldTagsFromPointer(&u, up) // Pass pointer to struct, not struct value

	if tags != nil {
		fmt.Printf("param tag: %s\n", tags["param"])
		fmt.Printf("json tag: %s\n", tags["json"])
		fmt.Printf("query tag: %s\n", tags["query"])
		fmt.Printf("form tag: %s\n", tags["form"])
	} else {
		fmt.Println("Field not found")
	}

	// Method 2: Using field name (simpler alternative)
	fmt.Println("\n=== Method 2: Using field name ===")
	tags2 := getFieldTagsByName(u, "Username")
	if tags2 != nil {
		fmt.Printf("param tag: %s\n", tags2["param"])
		fmt.Printf("json tag: %s\n", tags2["json"])
		fmt.Printf("query tag: %s\n", tags2["query"])
		fmt.Printf("form tag: %s\n", tags2["form"])
	}

	// Method 3: Test with Likes field
	fmt.Println("\n=== Method 3: Testing Likes field ===")
	likesPtr := &u.Likes
	tags3 := getFieldTagsFromPointer(&u, likesPtr) // Pass pointer to struct
	if tags3 != nil {
		fmt.Printf("param tag: %s\n", tags3["param"])
		fmt.Printf("json tag: %s\n", tags3["json"])
		fmt.Printf("query tag: %s\n", tags3["query"])
		fmt.Printf("form tag: %s\n", tags3["form"])
	}
}

// Alternative simpler approach using field name
func getFieldTagsByName(structValue interface{}, fieldName string) map[string]string {
	structVal := reflect.ValueOf(structValue)
	structType := structVal.Type()

	// Handle pointer to struct
	if structType.Kind() == reflect.Ptr {
		structVal = structVal.Elem()
		structType = structVal.Type()
	}

	// Find field by name
	if field, found := structType.FieldByName(fieldName); found {
		return map[string]string{
			"param": field.Tag.Get("param"),
			"json":  field.Tag.Get("json"),
			"query": field.Tag.Get("query"),
			"form":  field.Tag.Get("form"),
		}
	}

	return nil
}
