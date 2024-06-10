package conf

const (
	// build information: this is set at compile time using LDFlags
	buildVer  = "v1.1.0"
	buildDate = "2024-06-10"
)

func BuildVersion() string {
	return buildVer
}

func BuildDate() string {
	return buildDate
}
