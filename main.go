package main

import (
	"fmt"
	"os"

	"github.com/MakeNowJust/heredoc"
	"github.com/mazznoer/csscolorparser"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println(heredoc.Doc(`Usage: hexstring <css color>...`))
		os.Exit(1)
	}

	for _, arg := range os.Args[1:] {
		s, err := hexstring(arg)
		if err != nil {
			fmt.Printf("color %q is not a valid CSS color: %s\n", arg, err.Error())
			os.Exit(1)
		}

		fmt.Printf("%s ", s)
	}
	fmt.Println()
}

func hexstring(htmlcolor string) (string, error) {
	c, err := csscolorparser.Parse(htmlcolor)
	if err != nil {
		return "", err
	}
	return c.HexString(), nil
}
