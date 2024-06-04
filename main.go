package main

import (
	"fmt"

	"github.com/dantemogrim/imir/pkg/api"
	"github.com/dantemogrim/imir/pkg/flags"
	"github.com/dantemogrim/imir/pkg/messages"
	"github.com/dantemogrim/imir/pkg/styles"
)

func main() {
	flags.Init()

	IMIR, error := api.Fetch()

	if error != nil {
		fmt.Println(styles.Err(messages.ConnFailure + error.Error()))
		return
	}

	fmt.Println(messages.Result(IMIR))
}
