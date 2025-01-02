package TestPackage

var TestVar string = "You got it"

type FileTest struct {
	*fileTest
}

type fileTest struct {
	fd      int
	name    string
	dirinfo string
	nepipe  int32
}

func main() {
}
