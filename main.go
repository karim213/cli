package main

import (
	"os"
	"flag"
	"fmt"
)

func main()  {
	// checking if the command contains parameters
	if len(os.Args) < 2 {
		fmt.Println("Some parameters messing")
		os.Exit(1)
	}
	// switch on command type
	switch os.Args[1] {

		case "get":
			// `get` subCommand
			getCmd := flag.NewFlagSet("get", flag.ExitOnError)
			// inputs `get` command
			getAll := getCmd.Bool("all", false, "Get all videos")
			getID := getCmd.String("id", "", "Youtube video ID")
			getVideos(getCmd, getAll, getID)

		case "add":
			// `add` subCommand
			addCmd := flag.NewFlagSet("add", flag.ExitOnError)
			// inputs `add` command
			addId := addCmd.String("id", "", "Youtube video ID")
			addTitle := addCmd.String("title", "", "Youtube video ID")
			addDescription := addCmd.String("description", "", "Youtube video ID")
			addImageUrl := addCmd.String("imageUrl", "", "Youtube video ID")
			addUrl := addCmd.String("url", "", "Youtube video ID")
			saveVideos(addCmd, addId, addTitle, addDescription, addImageUrl, addUrl)

		default:
			fmt.Println("expected 'get' or 'add' subcommands")
			os.Exit(1)
		
	}
}
