package api

import (
	"fmt"
	"gitbatch/api/gitlab"
	"gitbatch/common"
	"gitbatch/config"
	"os"
	"strings"
)

func Clone(groupName string) {
	fmt.Println("git clone group:", groupName, "...")

	var projectUrls []string

	switch strings.ToLower(config.GitType) {
	case "gitlab":
		projectUrls = gitlab.GetProjectUrlsByGroup(config.GitPath, groupName)
	default:
		fmt.Print("git.type=" + config.GitType)
		fmt.Println(", only support gitlab now!!")
		os.Exit(1)
	}

	common.DoClone(groupName, projectUrls)
}

//ref是分支名
func GetFileFromRepository(projectName string, filePath string, ref string) gitlab.FileInfo {
	return gitlab.GetFileFromRepository(projectName, filePath, ref)
}
