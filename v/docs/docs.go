package docs

const (
	StringType = "string"
	IntType    = "int"
)

type RootDocs struct {
	Groups []*Group
}

type Group struct {
	Path   string
	Groups []*Group
	Routes []*Route
}

type Route struct {
	Method string
	Path   string
}

type Property struct {
	Name        string
	Type        string
	Format      string
	Description string
	Example     string
	MaxLength   int
	MinLength   int
	MinValue    int
	MaxValue    int
	Required    bool
}
