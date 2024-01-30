/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

// Mmm structs
// I actually lvoe this part of go
type Repository struct {
	FullName string `json:"full_name"`
	CloneURL string `json:"clone_url"`
}

var cloneCmd = &cobra.Command{
	Use:   "clone",
	Short: "Gets all git repos, tries to figure out which one you wanna download, and then downloads it",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: clone,
}

func init() {
	godotenv.Load()
	rootCmd.AddCommand(cloneCmd)
}

func clone(cmd *cobra.Command, args []string) {
	fmt.Println("clone called")
	//we'll do this one in indivudual funcs
	//connect to GH
	//get all repos
	repos := connect()
	//ask user which one they want

	fmt.Println("Select a project to start:")
	for i, repo := range repos {
		fmt.Printf("[%d] %s\n", i+1, repo.FullName)
	}

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter your choice: ")
	scanner.Scan()
	choice := scanner.Text()

	choiceIndex, err := strconv.Atoi(choice)
	if err != nil || choiceIndex < 1 || choiceIndex > len(repos) {
		fmt.Println("Invalid choice")
		return
	}

	//download it

}

func connect() []Repository {

	// i'll need to restructure this if I make a github app, can't seem to get pricate repos working

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.github.com/user/repos", nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Authorization", "token "+os.Getenv("GITHUB_ACCESS_TOKEN"))

	response, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var repositories []Repository
	err = json.Unmarshal(responseData, &repositories)
	if err != nil {
		log.Fatal(err)
	}

	for _, repo := range repositories {
		fmt.Println(repo.FullName)
		fmt.Println(repo.CloneURL)
	}
	return repositories
}
