package docs

const (
	StringType = "string"
	IntType    = "int"
)

type RootDocs struct {
	Groups []*Group `json:"groups"`
}

type Group struct {
	Path   string   `json:"path"`
	Groups []*Group `json:"groups"`
	Routes []*Route `json:"routes"`
}

type Route struct {
	Method      string `json:"method"`
	Path        string `json:"path"`
	Description string `json:"description"`
	Summary     string `json:"summary"`
	OperationId string `json:"operationId"`

	Request  ReqDocs  `json:"request"`
	Responce RespDocs `json:"response"`
}

type ReqDocs struct {
	Path    []DocsProperty `json:"path"`
	Query   []DocsProperty `json:"query"`
	Body    []DocsProperty `json:"body"`
	Headers []DocsProperty `json:"headers"`
}

type RespDocs struct {
	Body    []DocsProperty `json:"body"`
	Headers []DocsProperty `json:"headers"`
}

type DocsProperty struct {
	Name        string `json:"name"`
	Location    string `json:"localtion"`
	Type        string `json:"type"`
	Format      string `json:"format"`
	Description string `json:"description"`
	Example     string `json:"example"`
	MaxLength   int    `json:"maxLength"`
	MinLength   int    `json:"minLength"`
	MinValue    int    `json:"minValue"`
	MaxValue    int    `json:"maxValue"`
	Required    bool   `json:"required"`
}
