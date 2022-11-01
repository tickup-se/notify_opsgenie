package main

import (
	"fmt"
	"os"
	//"github.com/opsgenie/opsgenie-go-sdk-v2/client"
	//"github.com/opsgenie/opsgenie-go-sdk-v2/incident"
)

func main() {
	fmt.Println("Parsing parameters")
	allArguments := os.Args[1:]
	fmt.Println(allArguments)
	fmt.Println(len(allArguments))
	fmt.Println("GITHUB_ACTION:", os.Getenv("GITHUB_ACTION"))
	fmt.Println("GITHUB_ACTOR:", os.Getenv("GITHUB_ACTOR"))
	fmt.Println("GITHUB_EVENT_NAME:", os.Getenv("GITHUB_EVENT_NAME"))
	fmt.Println("GITHUB_REF:", os.Getenv("GITHUB_REF"))
	fmt.Println("GITHUB_REF_NAME:", os.Getenv("GITHUB_REF_NAME"))
	fmt.Println("GITHUB_REPOSITORY:", os.Getenv("GITHUB_REPOSITORY"))
	fmt.Println("GITHUB_REPOSITORY_OWNER:", os.Getenv("GITHUB_REPOSITORY_OWNER"))
	fmt.Println("GITHUB_SHA:", os.Getenv("GITHUB_SHA"))
	fmt.Println("GITHUB_WORKFLOW:", os.Getenv("GITHUB_WORKFLOW"))
	fmt.Println("RUNNER_ARCH:", os.Getenv("RUNNER_ARCH"))
	fmt.Println("RUNNER_NAME:", os.Getenv("RUNNER_NAME"))
	fmt.Println("RUNNER_OS:", os.Getenv("RUNNER_OS"))
}
