package helper

var version = "1.0.0"           // cannot be accessed outside of package (lowercase first letter)
var Application = "Go Language" // can be accessed globally (uppercase first letter)

// can be accessed globally (uppercase first letter)
func SayHello(name string) string {
	return "Hello " + name
}

// cannot be accessed outside of package (lowercase first letter)
func sayGoodbye(name string) string {
	return "Goodbye " + name
}