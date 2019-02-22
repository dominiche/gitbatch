package group

import (
	"gitbatch/api"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var group = &cobra.Command{
	Use:   "group",
	Short: "git clone projects by group name",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatalln("group name can't be empty!")
			os.Exit(1)
		}

		groupName := args[0]
		api.Clone(groupName)
	},
}

//func init() {
//	//clone.Flags().StringP("url", "u", "", "git project url")
//}

func Cmd() *cobra.Command {
	return group
}
