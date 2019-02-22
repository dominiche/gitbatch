package gitlab

import (
	"encoding/json"
	"fmt"
	"gitbatch/config"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

var apiPath = config.GitPath + "/api/v3/"

type idName struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func GetProjectUrlsByGroup(gitPath string, groupName string) []string {
	var groups []struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}

	var httpBasePath = gitPath + "/api/v3/groups"
	//1. search group name, get group:id
	searchPath := httpBasePath + "?search=" + groupName
	searchBody := getResult(searchPath)
	if err := json.Unmarshal(searchBody, &groups); err != nil {
		fmt.Println("Unmarshal response json error:", err)
		os.Exit(1)
	}

	var groupId int
	for _, group := range groups {
		if groupName == group.Name {
			groupId = group.Id
		}
	}
	if groupId == 0 {
		fmt.Println("can't find group: " + groupName)
		os.Exit(1)
	}

	//2. base groupName:id, get projects's url
	var urls []string
	projectsPath := httpBasePath + "/" + strconv.Itoa(groupId) + "/projects" + "?per_page=50"
	for i := 1; true; i++ {
		var projects []struct {
			SslUrl string `json:"ssh_url_to_repo"`
		}
		projectsPagePath := projectsPath + "&page=" + strconv.Itoa(i)
		fmt.Println("fetching projects:", projectsPagePath)
		projectsBody := getResult(projectsPagePath)
		if err := json.Unmarshal(projectsBody, &projects); err != nil {
			fmt.Println("Unmarshal response json error:", err)
			os.Exit(1)
		}

		if len(projects) == 0 {
			break
		}

		for _, project := range projects {
			urls = append(urls, project.SslUrl)
		}
	}

	return urls
}

func getResult(url string) []byte {
	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	request.Header.Add("PRIVATE-TOKEN", config.PrivateToken)

	if err != nil {
		panic(err)
	}
	response, err := client.Do(request)
	defer response.Body.Close()
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	return body
}

type FileInfo struct {
	FileName      string `json:"file_name"`
	Size          int    `json:"size"`
	Encoding      string `json:"encoding"`
	Content       string `json:"content"`
	ContentSha256 string `json:"content_sha256"`
	Ref           string `json:"ref"`
}

func GetFileFromRepository(projectName string, filePath string, ref string) FileInfo {
	var projects []idName
	//1. search project name, get project:id
	searchUrl := apiPath + "projects?simple=true&search=" + projectName
	fmt.Println("Get", searchUrl)
	searchBody := getResult(searchUrl)
	if err := json.Unmarshal(searchBody, &projects); err != nil {
		fmt.Println("Unmarshal response json error:", err)
		os.Exit(1)
	}

	var projectId int
	for _, project := range projects {
		if projectName == project.Name {
			projectId = project.Id
		}
	}
	if projectId == 0 {
		fmt.Println("can't find project: " + projectName)
		os.Exit(1)
	}

	//2. base projectName:id, get file
	var fileInfo FileInfo
	fileRequestUrl := apiPath + "projects/" + strconv.Itoa(projectId) +
		"/repository/files?file_path=" + filePath + "&ref=" + ref
	fmt.Println("Get", fileRequestUrl)
	fileBody := getResult(fileRequestUrl)
	if err := json.Unmarshal(fileBody, &fileInfo); err != nil {
		fmt.Println("Unmarshal response json error:", err)
		os.Exit(1)
	}

	return fileInfo
}
