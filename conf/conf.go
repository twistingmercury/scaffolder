package conf

var (
	// build information: this is set at compile time using LDFlags
	buildVer  = "n/a"
	buildDate = "n/a"
)

func BuildVersion() string {
	return buildVer
}

func BuildDate() string {
	return buildDate
}
