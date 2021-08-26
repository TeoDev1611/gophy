package cmds

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
)

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
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

func TouchFile(args ...string) error {
	if len(args) <= 1 {
		_, err := os.Stat(args[0])
		if os.IsNotExist(err) {
			file, err := os.Create(args[0])
			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()
		} else {
			currentTime := time.Now().Local()
			err = os.Chtimes(args[0], currentTime, currentTime)
			if err != nil {
				fmt.Println(err)
			}
		}
		return nil
	} else {
		return errors.New("Error not argument needed")
	}
}
