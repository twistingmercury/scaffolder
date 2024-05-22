package conf

var (
	// build information: this is set at compile time using LDFlags
	buildVer    = "n/a"
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
