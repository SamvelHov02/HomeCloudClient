package cli

type Command struct {
	Name        string
	Description string
	// Mapps flags and parameter values
	FlagsParam map[string]string
	Run        func(cmd *Command)
}

func (c *Command) Init(name string) {
	c.Name = name
	c.FlagsParam = make(map[string]string, 2)
}

func (c *Command) Build(args []string) {
	var key string
	for _, arg := range args {
		if arg[0] == '-' {
			key = arg
		} else {
			c.FlagsParam[key] = arg
		}
	}
}

func (c *Command) Execute() {
	c.Run(c)
}
