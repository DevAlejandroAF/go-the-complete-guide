package main

import "fmt"

type stringMap map[string]string

func main() {
	userNames := make([]string, 2, 5)

	userNames[0] = "Julie"
	userNames[1] = "Jose"

	userNames = append(userNames, "Max")
	userNames = append(userNames, "Manuel")

	fmt.Println(userNames)

	websites := make(stringMap)
	websites["LinkedIn"] = "https://linkedin.com"
}
