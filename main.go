package main

import (
	"fmt"
	"time"

	"github.com/dantemogrim/imir/pkg/api"
	"github.com/dantemogrim/imir/pkg/flags"
	"github.com/dantemogrim/imir/pkg/messages"
	"github.com/dantemogrim/imir/pkg/styles"
	"github.com/dantemogrim/imir/pkg/utils"
)

func main() {
	flags.Init()

	utils.SpinnerStart()

	IMIR, error := api.Fetch()

	// Simulate loading time 4 dramatic effect.
	time.Sleep(2 * time.Second)

	if error != nil {
		fmt.Println(styles.Err(messages.ConnFailure + error.Error()))
		return
	}

	fmt.Println(utils.ClearScreen)

	fmt.Println(messages.Result(IMIR))
}
