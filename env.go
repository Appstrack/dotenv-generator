package dotenv_generator

type (
	Environment struct {
		Name  string
		Value string
	}
	Environments []*Environment
)

// os.Getenv("VARNAME", "defaultvalue")
// os.Getenv("VAR1", "2")
// os.Getenv("VAR2", "3")
