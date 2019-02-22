package common

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func DoClone(dirName string, projectUrls []string) {
	if len(projectUrls) == 0 {
		fmt.Println("no project to clone!!")
		os.Exit(1)
	}

	//build dir
	pwd, _ := os.Getwd()
	var fullDir = pwd
	if filepath.Base(pwd) != dirName {
		err := os.Mkdir(dirName, 0777)
		if err != nil {
			panic(err)
		}
		fullDir = pwd + string(filepath.Separator) + dirName
	}
	fmt.Println(fullDir)
	fmt.Println("--------------------------------------------")
	for i := 0; i < len(projectUrls); i++ {
		url := projectUrls[i]
		//fmt.Println(url)
		command := exec.Command("git", "clone", url)
		command.Dir = fullDir
		command.Stdout = os.Stdout
		command.Stderr = os.Stderr
		err := command.Run()
		if err != nil {
			fmt.Printf("git clone failed : %s\n", err)
		}
		fmt.Println("--------------------------------------------")
	}
	fmt.Println("generating parent pom ...!")
	generateParentPom(fullDir, projectUrls)
	fmt.Println("Finish!")
}

func generateParentPom(groupDir string, projectUrls []string) {
	pom := groupDir + string(filepath.Separator) + "pom.xml"
	file, err := os.Create(pom)
	if err != nil {
		fmt.Println("generate Parent Pom failed:", err)
		return
	}
	defer file.Close()

	var b strings.Builder
	s1 := `<project xmlns="http://maven.apache.org/POM/4.0.0" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
  xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/xsd/maven-4.0.0.xsd">
  <modelVersion>4.0.0</modelVersion>
  <groupId>com.example</groupId>
  <artifactId>parent</artifactId>
  <version>0.0.1-SNAPSHOT</version>
  <packaging>pom</packaging>

  <modules>
`
	b.WriteString(s1)

	for _, url := range projectUrls {
		base := filepath.Base(url)
		projectName := strings.TrimSuffix(base, filepath.Ext(base))
		s := "    <module>" + projectName + `</module>
`
		b.WriteString(s)
	}

	s2 := ` </modules>

  <name>parent</name>

  <properties>
    <project.build.sourceEncoding>UTF-8</project.build.sourceEncoding>
    <skipTests>true</skipTests>
  </properties>
</project>
`
	b.WriteString(s2)

	pomString := b.String()
	_, e := io.WriteString(file, pomString)
	if e != nil {
		fmt.Println("write pom.xml failed:", err)
		fmt.Println("xml content is:\n", pomString)
		return
	}
	fmt.Println("generate parent pom.xml success!")
}
