package dbsampler

type (
	Row          []Field
	Rows         []Row
	TableGraph   []*Table
	TriggerGraph []*Trigger
	RoutineGraph []*Routine
	ViewGraph    []*View
	ColumnGraph  []*Column
)

// Field...
type Field struct {
	Column string
	Value string
}

// Trigger...
type Trigger struct {
	Name              string
	CreateSQL         string
	ActionTiming      string
	EventManipulation string
	EventObjectTable  string
	ActionOrientation string
	Definer           string
	SQLMode           string
	CharSet           string
	Collation         string
}

// Routine...
type Routine struct {
	Name            string
	Type            string
	CreateSQL       string
	SecurityType    string
	Definer         string
	ParamList       string
	Returns         string
	IsDeterministic string
	SQLMode         string
	CharSet         string
	Collation       string
}

// View...
type View struct {
	Name         string
	Columns      ColumnGraph
	CreateSQL    string
	CharSet      string
	Collation    string
	SecurityType string
	Definer      string
}

// Dependency...
type Dependency struct {
	TableName            string
	ColumnName           string
	ReferencedColumnName string
}

// Table stores the details of a single database table.
type Table struct {
	Name         string
	CreateSQL    string
	CharSet      string
	Collation    string
	Dependencies []*Dependency
	Triggers     TriggerGraph
	Rows Rows
}

// Column...
type Column struct {
	Name            string
	OrdinalPosition int
	Type            string
}
