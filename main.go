package main

import (
	"fmt"
	"time"

	"github.com/dantemogrim/imir/cmd"
	"github.com/dantemogrim/imir/pkg/utils"
	// "github.com/gernest/wow/spin"
)

func main() {
	cmd.ParseFlags()

	utils.SpinnerStart()


	// IMIR, error := api.Fetch()

	time.Sleep(2 * time.Second)

	// spinner.PersistWith(spin.Spinner{Frames: []string{""}}, "")

	fmt.Printf(utils.ClearScreen)

	// if error != nil {
	// 	fmt.Println(styles.ApiError(messages.ConnFailure + error.Error()))
	// 	return
	// }

	fmt.Printf(utils.ClearScreen)

	// status := styles.BackwardRotation(messages.BackwardRotation)

	// if !IMIR {
	// 	status = styles.ForwardRotation(messages.ForwardRotation)
	// }

//	fmt.Println(status)
}
