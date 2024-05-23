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

	// Simulate loading time for dramatic effect.
	time.Sleep(2 * time.Second)

	fmt.Printf(utils.ClearScreen)

	if error != nil {
		fmt.Println(styles.ApiError(messages.ConnFailure + error.Error()))
		return
	}

	fmt.Printf(utils.ClearScreen)

	fmt.Println(messages.Result(IMIR))
}
