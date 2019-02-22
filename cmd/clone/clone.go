package clone

import (
	"gitbatch/cmd/clone/group"
	"gitbatch/cmd/clone/repo"
	"github.com/spf13/cobra"
)

var clone = &cobra.Command{
	Use:   "clone",
	Short: "batch git clone",
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println(args)
		//url, _ := cmd.Flags().GetString("url")
		//fmt.Println("url=", url)
	},
}

func init() {
	clone.AddCommand(group.Cmd())
	clone.AddCommand(repo.Cmd())
}

func Cmd() *cobra.Command {
	return clone
}
