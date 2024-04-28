package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/MakeNowJust/heredoc"
	"github.com/docopt/docopt-go"
	"github.com/mazznoer/csscolorparser"
)

var Usage = heredoc.Doc(`Usage: hexstring [options] <css color>...

Options:
  -H  don't add an '#' in front	
  -n  don't add a newline at the end
  `)

func main() {
	args, _ := docopt.ParseDoc(Usage)

	for _, arg := range args["<css color>"].([]string) {
		s, err := hexstring(arg)
		if err != nil {
			fmt.Printf("color %q is not a valid CSS color: %s\n", arg, err.Error())
			os.Exit(1)
		}

		if flag(args, "-H") {
			s = strings.TrimPrefix(s, "#")
		}

		fmt.Printf("%s ", s)
	}
	if !flag(args, "-n") {
		fmt.Println()
	}
}

func flag(opts docopt.Opts, key string) bool {
	b, _ := opts.Bool(key)
	return b
}

func hexstring(htmlcolor string) (string, error) {
	c, err := csscolorparser.Parse(htmlcolor)
	if err != nil {
		return "", err
	}
	return c.HexString(), nil
}
