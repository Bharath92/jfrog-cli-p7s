package commands

import (
	"errors"
	"os"
	"strconv"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/jfrog/jfrog-cli-core/v2/plugins/components"
	"github.com/jfrog/jfrog-client-go/utils/log"
)

func GetHelloCommand() components.Command {
	return components.Command{
		Name:        "pipelines",
		Description: "Pipelines GUI.",
		Aliases:     []string{"p7s"},
		Arguments:   getHelloArguments(),
		Flags:       getHelloFlags(),
		EnvVars:     getHelloEnvVar(),
		Action: func(c *components.Context) error {
			dir, err := os.Getwd()
			if err != nil {
				log.Error(err)
			}
			log.Error(dir)
			r, err := git.PlainOpen(dir)
			if err != nil {
				log.Error(err)
			}
			log.Error(r.Head())
			remote, err := r.Remotes()
			if err != nil {
				log.Error(err)
			}
			log.Error(remote[0].Config().URLs[0])
			return nil
		},
	}
}

func getHelloArguments() []components.Argument {
	return []components.Argument{
		{
			Name:        "addressee",
			Description: "The name of the person you would like to greet.",
		},
	}
}

func getHelloFlags() []components.Flag {
	return []components.Flag{
		components.BoolFlag{
			Name:         "shout",
			Description:  "Makes output uppercase.",
			DefaultValue: false,
		},
		components.StringFlag{
			Name:         "repeat",
			Description:  "Greets multiple times.",
			DefaultValue: "1",
		},
	}
}

func getHelloEnvVar() []components.EnvVar {
	return []components.EnvVar{
		{
			Name:        "HELLO_FROG_GREET_PREFIX",
			Default:     "A new greet from your plugin template: ",
			Description: "Adds a prefix to every greet.",
		},
	}
}

type helloConfiguration struct {
	addressee string
	shout     bool
	repeat    int
	prefix    string
}

func helloCmd(c *components.Context) error {
	if len(c.Arguments) != 1 {
		return errors.New("Wrong number of arguments. Expected: 1, " + "Received: " + strconv.Itoa(len(c.Arguments)))
	}
	var conf = new(helloConfiguration)
	conf.addressee = c.Arguments[0]
	conf.shout = c.GetBoolFlagValue("shout")

	repeat, err := strconv.Atoi(c.GetStringFlagValue("repeat"))
	if err != nil {
		return err
	}
	conf.repeat = repeat

	conf.prefix = os.Getenv("HELLO_FROG_GREET_PREFIX")
	if conf.prefix == "" {
		conf.prefix = "New greeting: "
	}

	log.Output(doGreet(conf))
	return nil
}

func doGreet(c *helloConfiguration) string {
	greet := c.prefix + "Hello " + c.addressee + "!\n"

	if c.shout {
		greet = strings.ToUpper(greet)
	}

	return strings.TrimSpace(strings.Repeat(greet, c.repeat))
}
