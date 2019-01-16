package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type Event struct {
	Username string //upper case public

}

func handler(e Event) (string, error) {
	// if no username is given error
	if len(e.Username) == 0 {
		return "", fmt.Errorf("No Name Given")
	}
	// if name starts with D error
	if e.Username[0] == 'D' {
		return "", fmt.Errorf("Dont like : %s", e.Username)
	}
	return fmt.Sprintf("<h1>hello %s from lambda Go</h1>", e.Username), nil
}
func main() {
	lambda.Start(handler)
}
