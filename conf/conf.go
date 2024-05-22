package conf

var (
	// build information: this is set at compile time using LDFlags
	buildVer    = "1.0.1"
	buildCommit = "7064669"
)

func BuildVersion() string {
	return buildVer
}

func BuildCommit() string {
	return buildCommit
}
