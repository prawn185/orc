package cmd

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists available projects",
	Long:  `Lists available projects. For example:`,
	Run:   availableProjects,
}

func init() {
	rootCmd.AddCommand(listCmd)
}

/*
This was my old method to get all the projects, my new one is now in start.go fun! :)
*/
func availableProjects(cmd *cobra.Command, args []string) {
	path := "../"
	file1 := "docker-compose.yml"
	file2 := "makefile"
	green := color.New(color.FgGreen).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()
	prefix := "project-"
	directories, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Printf("Error reading directory: %s\n", err)
		return
	}

	for _, dir := range directories {
		if dir.IsDir() && strings.HasPrefix(dir.Name(), prefix) {
			dirPath := filepath.Join(path, dir.Name())
			files, err := ioutil.ReadDir(dirPath)
			if err != nil {
				fmt.Printf("Error reading directory '%s': %s\n", dirPath, err)
				continue
			}

			found1, found2 := false, false
			for _, file := range files {
				if file.Name() == file1 {
					found1 = true
				} else if file.Name() == file2 {
					found2 = true
				}
			}

			if found1 || found2 {
				fmt.Println(green("✔"), dir.Name())
			} else {
				fmt.Println(red("✖"), dir.Name())
			}

		}
	}
}
