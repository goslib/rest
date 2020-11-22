package rest

// The options to alter the default behaviors of the package.
type config struct {
	// Whether log out the response error or not.
	Logger bool
	// The separator between modules used in the default logger.
	TaggingSeparator string

	Debug bool
}

var conf = &config{
	true,
	"›",
	true,
}

// The exported package configuration, get and set the parameters.
func GetConfigure() *config {
	return conf
}
