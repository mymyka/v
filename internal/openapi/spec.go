package openapi

// Header

type Spec struct {
	OpenAPI    string          `json:"openapi" yaml:"openapi"`
	Info       Info            `json:"info" yaml:"info"`
	Servers    []Server        `json:"servers" yaml:"servers"`
	Components Components      `json:"components" yaml:"components"`
	Paths      map[string]Path `json:"paths" yaml:"paths"`
}

type Info struct {
	Title       string  `json:"title" yaml:"title"`
	Description string  `json:"description" yaml:"description"`
	Version     string  `json:"version" yaml:"version"`
	Contact     Contact `json:"contact" yaml:"contact"`
	License     License `json:"license" yaml:"license"`
}

type Contact struct {
	Name  string `json:"name" yaml:"name"`
	Email string `json:"email" yaml:"email"`
}

type License struct {
	Name string `json:"name" yaml:"name"`
	URL  string `json:"url" yaml:"url"`
}

type Server struct {
	URL         string `json:"url" yaml:"url"`
	Description string `json:"description" yaml:"description"`
}

// Paths

type Path struct {
	Methods map[string]Method
}

type Method struct {
	Tags        []string    `json:"tags" yaml:"tags"`
	Summary     string      `json:"summary" yaml:"summary"`
	Description string      `json:"description" yaml:"description"`
	OperationId string      `json:"operationId" yaml:"operationId"`
	RequestBody RequestBody `json:"requestBody" yaml:"requestBody"`
	Responses   Responses   `json:"responses" yaml:"responses"`
}

type RequestBody struct {
	Required bool    `json:"required" yaml:"required"`
	Content  Content `json:"content" yaml:"content"`
}

type Content struct {
	ContentTypes []ContentType
}

type ContentType struct {
	ContentType string
	Schema      Schema    `json:"schema" yaml:"schema"`
	Examples    []Example `json:"examples" yaml:"examples"`
}

type Schema struct {
	Ref string `json:"$ref" yaml:"$ref"`
}

type Example struct {
	Name    string
	Summary string  `json:"summary" yaml:"summary"`
	Value   []Value `json:"value" yaml:"value"`
}

type Value struct {
	Name  string
	Value string
}

// Responses

type Responses struct {
	StatusCodes []StatusCode
}

type StatusCode struct {
	Code        string   `json:"-" yaml:"-"` // HTTP status code like "200", "400", etc.
	Description string   `json:"description" yaml:"description"`
	Content     Content  `json:"content" yaml:"content"`
	Headers     []Header `json:"headers" yaml:"headers"`
}

type Header struct {
	Name        string `json:"-" yaml:"-"`
	Description string `json:"description" yaml:"description"`
	Schema      Schema `json:"schema" yaml:"schema"`
}

// Components

type Components struct {
	Schemas map[string]ComponentSchema `json:"schemas" yaml:"schemas"`
}

type ComponentSchema struct {
	Type                 string              `json:"type" yaml:"type"`
	Format               string              `json:"format,omitempty" yaml:"format,omitempty"`
	Description          string              `json:"description,omitempty" yaml:"description,omitempty"`
	Example              interface{}         `json:"example,omitempty" yaml:"example,omitempty"`
	Required             []string            `json:"required,omitempty" yaml:"required,omitempty"`
	Properties           map[string]Property `json:"properties,omitempty" yaml:"properties,omitempty"`
	Items                *ComponentSchema    `json:"items,omitempty" yaml:"items,omitempty"`
	AdditionalProperties *bool               `json:"additionalProperties,omitempty" yaml:"additionalProperties,omitempty"`
	MinLength            *int                `json:"minLength,omitempty" yaml:"minLength,omitempty"`
	MaxLength            *int                `json:"maxLength,omitempty" yaml:"maxLength,omitempty"`
	Minimum              *float64            `json:"minimum,omitempty" yaml:"minimum,omitempty"`
	Maximum              *float64            `json:"maximum,omitempty" yaml:"maximum,omitempty"`
	Ref                  string              `json:"$ref,omitempty" yaml:"$ref,omitempty"`
}

type Property struct {
	Type        string              `json:"type" yaml:"type"`
	Format      string              `json:"format,omitempty" yaml:"format,omitempty"`
	Description string              `json:"description,omitempty" yaml:"description,omitempty"`
	Example     interface{}         `json:"example,omitempty" yaml:"example,omitempty"`
	MinLength   *int                `json:"minLength,omitempty" yaml:"minLength,omitempty"`
	MaxLength   *int                `json:"maxLength,omitempty" yaml:"maxLength,omitempty"`
	Minimum     *float64            `json:"minimum,omitempty" yaml:"minimum,omitempty"`
	Maximum     *float64            `json:"maximum,omitempty" yaml:"maximum,omitempty"`
	Items       *ComponentSchema    `json:"items,omitempty" yaml:"items,omitempty"`
	Properties  map[string]Property `json:"properties,omitempty" yaml:"properties,omitempty"`
	Required    []string            `json:"required,omitempty" yaml:"required,omitempty"`
	Ref         string              `json:"$ref,omitempty" yaml:"$ref,omitempty"`
}
