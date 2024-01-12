package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/gonuts/commander"
)

func makeCmdClean(filename string) *commander.Command {
	cmdClean := func(cmd *commander.Command, args []string) error {
		if len(args) != 0 {
			cmd.Usage()
			return nil
		}
		w, err := os.Create(filename + "_")
		if err != nil {
			return err
		}
		defer w.Close()
		f, err := os.Open(filename)
		if err != nil {
			return err
		}
		br := bufio.NewReader(f)
		for {
			b, _, err := br.ReadLine()
			if err != nil {
				if err != io.EOF {
					return err
				}
				break
			}
			line := string(b)
			if !strings.HasPrefix(line, "-") {
				_, err = fmt.Fprintf(w, "%s\n", line)
				if err != nil {
					return err
				}
			}
		}
		f.Close()
		w.Close()
		err = os.Remove(filename)
		if err != nil {
			return err
		}
		return os.Rename(filename+"_", filename)
	}

	return &commander.Command{
		Run:       cmdClean,
		UsageLine: "clean",
		Short:     "remove all done items",
	}
}
