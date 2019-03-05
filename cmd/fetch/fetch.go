package fetch

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

var fetch = &cobra.Command{
	Use:   "fetch",
	Short: "batch git fetch",
	Long:  "only support 'git fetch'",
	Run: func(cmd *cobra.Command, args []string) {
		var pullCmd *exec.Cmd
		pullCmd = exec.Command("git", "fetch")
		pullCmd.Stdout = os.Stdout
		pullCmd.Stderr = os.Stderr
		doFetch(pullCmd)
	},
}

func Cmd() *cobra.Command {
	return fetch
}

func doFetch(theCmd *exec.Cmd) {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	dir, err = filepath.Abs(dir)
	if err != nil {
		panic(err)
	}
	ifs, err := ioutil.ReadDir(dir)
	for _, info := range ifs {
		if info.IsDir() {
			cmd := new(exec.Cmd)
			*cmd = *theCmd
			fmt.Printf("--------------------------------------------\n%s:\n", info.Name())
			cmd.Dir = dir + string(filepath.Separator) + info.Name()
			cmd.Run()
		}
	}
}
