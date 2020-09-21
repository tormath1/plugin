package plugin

// Type defines the type of the plugin
type Type int

//go:generate ./bin/enumer -type=Type -transform=lower -output=type_string.go
const (
	Local Type = iota
	Remote
	Embedded
)


