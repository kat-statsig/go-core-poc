package main

import (
	"fmt"
	"strconv"
	"time"

	statsig "github.com/kat-statsig/go-core-poc"
)

func init() {
    fmt.Println("This runs when the package is imported")
}

func main() {
	for i := 0; i < 100; i++ {
		user := statsig.NewStatsigUser(
			"test-user"+strconv.Itoa(i),
			"test@test.com",
			"127.0.0.1",
			"test-user-agent",
			"US",
			"en-US",
			"1.0.0",
			map[string]interface{}{},
			map[string]interface{}{},
			map[string]string{},
			map[string]string{},
		)
		fmt.Println(user.GetInnerRef())
		user = nil
		time.Sleep(1 * time.Second)
	}
} 