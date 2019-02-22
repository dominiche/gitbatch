package checkout

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

var checkout = &cobra.Command{
	Use:   "checkout",
	Short: "batch git checkout",
	Long:  "only support two usages: \n git checkout branch_name  \n git checkout -b new_branch_name origin_branch ",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("Only one argument needed!")
			os.Exit(1)
		}
		arg0 := args[0]
		newBranchName, _ := cmd.Flags().GetString("branch")
		var checkoutCmd *exec.Cmd
		if newBranchName == "" {
			fmt.Printf("checkout %s:\n", arg0)
			checkoutCmd = exec.Command("git", "checkout", arg0)
		} else {
			fmt.Printf("new branch %s, then checkout:\n", newBranchName)
			checkoutCmd = exec.Command("git", "checkout", "-b", newBranchName, arg0)
		}
		checkoutCmd.Stdout = os.Stdout
		checkoutCmd.Stderr = os.Stderr

		doCheckout(checkoutCmd)
	},
}

func init() {
	checkout.Flags().StringP("branch", "b", "", "new_branch name")
}

func Cmd() *cobra.Command {
	return checkout
}

func doCheckout(checkoutCmd *exec.Cmd) {
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
			*cmd = *checkoutCmd
			fmt.Printf("--------------------------------------------\n%s:\n", info.Name())
			cmd.Dir = dir + string(filepath.Separator) + info.Name()
			cmd.Run()
		}
	}
}
