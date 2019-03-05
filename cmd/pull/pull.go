package pull

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

var pull = &cobra.Command{
	Use:   "pull",
	Short: "batch git pull",
	Long:  "only support 'git pull'",
	Run: func(cmd *cobra.Command, args []string) {
		var pullCmd *exec.Cmd
		pullCmd = exec.Command("git", "pull")
		pullCmd.Stdout = os.Stdout
		pullCmd.Stderr = os.Stderr
		doPull(pullCmd)
	},
}

func Cmd() *cobra.Command {
	return pull
}

func doPull(theCmd *exec.Cmd) {
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
