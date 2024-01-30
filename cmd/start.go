package cmd

/*
Still lots to do & learn
But tl;dr this file can
Check if you need/your Proxy is working
Adds a spinner while it starts, if it doesn't think it's started & it will ask you if you want to start it

It will check the ../dir you have it installed and lists all the projects you have, THAT contain a makefile or a docker compose
it will then ask you to select one, and it will prioitise the makefile if it exists, if not it will use docker compose up -d



*/

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/fatih/color"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start a selected project using Docker Compose",
	Long:  `Start a selected project using Docker Compose. For example:`,
	Run:   startProject,
}

var cursors = []string{"|", "/", "-", "\\"}
var green = color.New(color.FgGreen).SprintFunc()
var red = color.New(color.FgRed).SprintFunc()

func init() {
	err := godotenv.Load()
	if err != nil {
		return
	}
	rootCmd.AddCommand(startCmd)
}

func startProject(cmd *cobra.Command, args []string) {

	//before we do anything lets do some preflight checks
	preChecks()

	path := "../"
	projects := listProjects(path)

	fmt.Println("Select a project to start:")
	for i, project := range projects {
		fmt.Printf("[%d] %s\n", i+1, green("✔ ")+project)
	}

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter your choice: ")
	scanner.Scan()
	choice := scanner.Text()

	choiceIndex, err := strconv.Atoi(choice)
	if err != nil || choiceIndex < 1 || choiceIndex > len(projects) {
		fmt.Println("Invalid choice")
		return
	}
	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	fmt.Println("Starting project:", projects[choiceIndex-1])
	var hasMakefile, hasDockerCompose bool
	for _, file := range files {
		if file.IsDir() && file.Name() == projects[choiceIndex-1] {
			fmt.Printf("Subdirectory: %s\n", file.Name())

			subDirPath := filepath.Join(path, file.Name())

			subFiles, err := os.ReadDir(subDirPath)
			if err != nil {
				fmt.Println("Error reading subdirectory:", err)
				continue
			}

			for _, subFile := range subFiles {
				switch subFile.Name() {
				case "makefile":
					hasMakefile = true
				case "docker-compose.yml", "docker-compose.yaml":
					hasDockerCompose = true
				}
			}

		}
	}

	if hasMakefile {
		fmt.Printf("Starting project %s using makefile\n", projects[choiceIndex-1])
		exec.Command("make", "up").Run()
	} else if hasDockerCompose {
		fmt.Printf("Starting project %s using docker-compose\n", projects[choiceIndex-1])
		exec.Command("docker-compose", "up", "-d").Run()
	}
}

func preChecks() {
	//@todo:
	// ✔ check if on VPN
	//check if proxy is running
	print("sh", "-c", "docker ps | grep proxy")
	//maybe find a way to better check if the proxy is running, bind it?
	proxyCheck := exec.Command("sh", "-c", "docker ps | grep proxy")
	var out bytes.Buffer
	proxyCheck.Stdout = &out
	result := proxyCheck.Run()

	if result != nil {

		fmt.Println(red("✖ "), os.Getenv("PROXY_NAME")+" is not running, would you like to start it? [Y, N]")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		choice := scanner.Text()

		if choice == "Y" || choice == "y" || choice == "Yes" || choice == "yes" {
			fmt.Println("Starting", os.Getenv("PROXY_NAME"))
			stopSpinner := make(chan bool)
			go startSpinner(stopSpinner)
			path := "../" + os.Getenv("PROXY_NAME")

			cmd := exec.Command("make", "up")
			cmd.Dir = path

			if err := cmd.Run(); err != nil {
				fmt.Println("Error running Docker command:", err)
				stopSpinner <- true
				return
			}

			stopSpinner <- true
			proxyCheck.Run()
			// Additional logic after the Docker command completes
			fmt.Println(green("✔ "), os.Getenv("PROXY_NAME")+" Started")
			fmt.Println("------------")

		}
	} else {
		if out.String() != "" {
			fmt.Println(green("✔ ") + os.Getenv("PROXY_NAME") + " is already running")
			fmt.Println("------------")
		}
	}
}

func listProjects(path string) []string {
	var projects []string

	directories, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Printf("Error reading directory: %s\n", err)
		return nil
	}

	for _, dir := range directories {
		if dir.IsDir() && strings.HasPrefix(dir.Name(), os.Getenv("FILE_PREFIX")) {
			files, err := ioutil.ReadDir(filepath.Join(path, dir.Name()))
			if err != nil {
				fmt.Printf("Error reading directory '%s': %s\n", filepath.Join(path, dir.Name()), err)
				continue
			}

			for _, file := range files {
				if file.Name() == "makefile" || file.Name() == "docker-compose.yml" || file.Name() == "docker-compose.yaml" {
					projects = append(projects, dir.Name())
					break
				}
			}
		}
	}

	return projects
}

func startSpinner(stop chan bool) {
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	i := 0
	for {
		select {
		case <-stop:
			fmt.Print("\r \r") // Clear the spinner
			return
		case <-ticker.C:
			fmt.Printf("\r%s", cursors[i%len(cursors)])
			i++
		}
	}
}
