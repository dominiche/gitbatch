package repo

import (
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"gitbatch/api"
	"gitbatch/api/gitlab"
	"gitbatch/common"
	"gitbatch/config"
	"github.com/spf13/cobra"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func Cmd() *cobra.Command {
	return repo
}

var repo = &cobra.Command{
	Use:   "repo",
	Short: "git clone projects by repo xml file",
	Long:  "the arg can be a local file path; or a git file path with the format project_name:file_path:ref",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatalln("repo name can't be empty!")
			os.Exit(1)
		}

		repoHandleFunc(args)
	},
}

type repoXml struct {
	XMLName xml.Name `xml:"manifest"`
	Project []struct {
		Name string `xml:"name,attr"`
	} `xml:"project"`
}

func repoHandleFunc(args []string) {
	var projectUrls []string
	var dirName string
	arg0 := args[0]
	if strings.Contains(arg0, ":") {
		arr := strings.Split(arg0, ":")
		if len(arr) != 3 {
			log.Fatalln("a git file path should be like 'project_name:file_path:ref' !")
			os.Exit(1)
		}
		fileInfo := api.GetFileFromRepository(arr[0], arr[1], arr[2])
		ext := filepath.Ext(fileInfo.FileName)
		if ext != ".xml" {
			log.Fatalln("error: the git file may not be a repo xml file.")
			os.Exit(1)
		}
		dirName = strings.TrimSuffix(fileInfo.FileName, ext)
		saveRepoFile(fileInfo)
		projectUrls = getProjectFromFile(fileInfo.FileName, projectUrls)

	} else if strings.HasSuffix(arg0, ".xml") {
		dirName = strings.TrimSuffix(arg0, ".xml")
		projectUrls = getProjectFromFile(arg0, projectUrls)
	} else {
		log.Fatalln("error：maybe not a repo xml format file!")
		os.Exit(1)
	}
	common.DoClone(dirName, projectUrls)
}

func saveRepoFile(fileInfo gitlab.FileInfo) {
	//save the repo xml file
	file, err := os.Create(fileInfo.FileName)
	if err != nil {
		log.Fatalln("save the repo xml file failed:", err)
		os.Exit(1)
	}
	defer file.Close()
	encodedString := fileInfo.Content
	bytes, err := base64.StdEncoding.DecodeString(encodedString)
	if err != nil {
		log.Fatalln("base64.StdEncoding.DecodeString failed:", err)
		os.Exit(1)
	}
	_, e := io.WriteString(file, string(bytes))
	if e != nil {
		log.Fatalln("write repo xml file failed:", err)
		os.Exit(1)
	}
	fmt.Println("generate parent pom.xml success!")
}

func getProjectFromFile(filePath string, projectUrls []string) []string {
	var rpl repoXml
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
	err = xml.Unmarshal(bytes, &rpl)
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
	for _, project := range rpl.Project {
		sslUrl := "git@" + path.Base(config.GitPath) + ":" + project.Name + ".git"
		projectUrls = append(projectUrls, sslUrl)
	}
	return projectUrls
}
