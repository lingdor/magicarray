package magicarray

// NamingDefault use the order of Json,hump
const NamingDefault = NamingJsonFirst

// NamingJson use the order of json,hump format
const NamingJsonFirst = 51

// NamingOrm use order of gorm or underline split rule for naming
const NamingOrmFirst = 62

// NamingHump Naming for the hump rule format,example: helloWord
const NamingHump = 1

// NamingUnderLine Naming for the underline split words rule format, example: hello_world
const NamingUnderLine = 2

// NamingRaw Naming for raw value of rule
const NamingRaw = 3

const NamingJsonFormat = 5
const NamingGormFormat = 6

func GenerateNamingFormatter(code int) func(string) string {
	return func(input string) string {

		return input
	}
}
