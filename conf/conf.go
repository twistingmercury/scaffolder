package conf

var (
	// build information: this is set at compile time using LDFlags
	buildVer    = "1.0.1"
	buildDate   = "n/a"
	buildCommit = "n/a"
)

func BuildVersion() string {
	return buildVer
}

func BuildData() string {
	return buildDate
}

func BuildCommit() string {
	return buildCommit
}
