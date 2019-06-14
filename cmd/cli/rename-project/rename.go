package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/otiai10/copy"
)

var wd, _ = os.Getwd()
var paths = strings.Split(wd, string(os.PathSeparator))
var projectName = paths[len(paths)-1]

func getWalkFunc(newProjectName string) filepath.WalkFunc {
	return func(path string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !!fi.IsDir() {
			return nil
		}

		var ignore = regexp.MustCompile(`^vendor\/.*?`)

		if ignore.MatchString(path) == true {
			return filepath.SkipDir
		}

		matched, _ := filepath.Match("*.go", fi.Name())

		if matched && fi.Name() != "rename.go" {
			read, err := ioutil.ReadFile(path)
			if err != nil {
				panic(err)
			}

			newContents := strings.Replace(string(read), projectName+"/", newProjectName+"/", -1)

			err = ioutil.WriteFile(path, []byte(newContents), 0)
			if err != nil {
				panic(err)
			}
		}

		return nil
	}
}

func RenameProject(newProjectName string) {
	copy.Copy("../"+projectName+"/cmd", "../"+newProjectName+"/cmd")
	copy.Copy("../"+projectName+"/internal", "../"+newProjectName+"/internal")
	copy.Copy("../"+projectName+"/.gitignore", "../"+newProjectName+"/.gitignore")
	copy.Copy("../"+projectName+"/Gopkg.lock", "../"+newProjectName+"/Gopkg.lock")
	copy.Copy("../"+projectName+"/Gopkg.toml", "../"+newProjectName+"/Gopkg.toml")
	copy.Copy("../"+projectName+"/README.md", "../"+newProjectName+"/README.md")
	copy.Copy("../"+projectName+"/internal/config/samples/dev.yaml.sample", "../"+newProjectName+"/internal/config/dev.yaml")
	copy.Copy("../"+projectName+"/internal/config/samples/prod.yaml.sample", "../"+newProjectName+"/internal/config/prod.yaml")

	err := filepath.Walk("../"+newProjectName, getWalkFunc(newProjectName))
	if err != nil {
		panic(err)
	}

	DepEnsure(newProjectName)

	fmt.Println("Successfully!")
}

func DepEnsure(projectName string) {
	cmd := exec.Command("sh", "-c", "dep ensure")
	cmd.Dir = "../" + projectName
	cmd.Start()
}

func main() {

	prompt := promptui.Select{
		Label: "Select Option",
		Items: []string{
			"❆ Rename, backup and setup project",
			"❆ Setup current project with dep",
			"❆ Quit",
		},
	}

	i, _, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	switch i {
	case 0:
		validate := func(input string) error {
			match, _ := regexp.MatchString(`[^/?*:;{}\\]+\\.[^/?*:;{}\\]+`, input)
			if match == true {
				return errors.New("Invalid Project Name")
			}
			return nil
		}

		prompt := promptui.Prompt{
			Label:    "Project Name",
			Validate: validate,
		}

		result, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		RenameProject(result)

	case 1:
		copy.Copy("../"+projectName+"/internal/config/samples/dev.yaml.sample", "../"+projectName+"/internal/config/dev.yaml")
		copy.Copy("../"+projectName+"/internal/config/samples/prod.yaml.sample", "../"+projectName+"/internal/config/prod.yaml")
		DepEnsure(projectName)
		fmt.Println("Successfully!")

	default:
		os.Exit(3)
	}
}
