package cmds

import (
	"errors"
	"io/ioutil"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
)

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func ListFiles(args ...string) error {
	if len(args) <= 1 {
		wd, _ := os.Getwd()
		files, _ := ioutil.ReadDir(wd)
		for _, f := range files {
			t := table.NewWriter()
			t.SetOutputMirror(os.Stdout)
			t.AppendHeader(table.Row{"NAME", "SIZE", "MODE", "TIME"})
			t.AppendRows([]table.Row{
				{f.Name(), f.Size(), f.Mode(), f.ModTime()},
			})
			t.AppendSeparator()
			t.Render()
		}
		return nil
	} else {
		return errors.New("Error not argument needed")
	}
}
