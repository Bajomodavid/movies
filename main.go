package main

import (
	"github.com/go-phie/gophie/cmd"
	"os"
	"fmt"
	"strings"
)

func main() {
	// for _, env := range os.Environ() {
	// 	// env is
	// 	envPair := strings.SplitN(env, "=", 2)
	// 	key := envPair[0]
	// 	value := envPair[1]

	// 	fmt.Printf("%s : %s\n", key, value)
	// }
	var allowedHosts = "http://127.0.0.1,http://105.112.124.195"
	os.Setenv("WHITE_LISTED_HOSTS", allowedHosts)
	fmt.Printf("%s", strings.Split(os.Getenv("WHITE_LISTED_HOSTS"), ","))
	cmd.Execute()
}
