package cli

import (
	"os/exec"
	"strings"

	"github.com/tesh254/simp/helpers"

	"github.com/fatih/color"
)

// CLIMethods holds all cli methods
type CLIMethods struct{}

func (c *CLIMethods) parseCommands(command string) []string {
	return strings.Fields(command)
}

// Run initializes running the commands in simp.json file
func (c *CLIMethods) Run(order string) {
	arguments := c.parseCommands(order)

	root := arguments[0]

	commandArguments := arguments[1:]

	c.RunCommand(root, commandArguments, true)
}

// RunAllCommands loops to execute all commands
func (c *CLIMethods) RunAllCommands(commands []string) {
	for _, element := range commands {
		c.Run(element)
	}
}

// RunCommand handles running a single command
func (c *CLIMethods) RunCommand(path string, args []string, debug bool) {
	cmd := exec.Command(path, args...)

	var log helpers.Logger

	var b []byte

	b, err := cmd.CombinedOutput()

	if err != nil {
		color.Red("😡 {Error}: \n" + string(err.Error()))
		// Error log
		errorLog := "Log: " + helpers.NewSHA1Hash(16) + " -> " + path + " " + strings.Join(args, ",") + "\n" + "---------------------------------------------------------------------------" + "\n" + "---------------------------------------------------------------------------" + "\n" + "[ERROR] \n" + err.Error() + "\n" + "---------------------------------------------------------------------------" + "\n" + "---------------------------------------------------------------------------" + "\n\n"

		log.CreateLoggerFileAndPopulate(errorLog)
	}

	color.Green(string(b[:len(b)]))

	outputLog := "Log: " + helpers.NewSHA1Hash(16) + " -> " + path + " " + strings.Join(args, ",") + "\n" + "---------------------------------------------------------------------------" + "\n" + "---------------------------------------------------------------------------" + "\n" + "[OUTPUT] \n" + string(b[:len(b)]) + "\n" + "---------------------------------------------------------------------------" + "\n" + "---------------------------------------------------------------------------" + "\n\n"

	log.CreateLoggerFileAndPopulate(outputLog)
}
