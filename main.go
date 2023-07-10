package main

import "github.com/sunflower10086/restful-api-demo/cmd"

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		panic(err)
	}
}
