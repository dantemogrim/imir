package main

import (
	"fmt"
	"os"
	"time"

	"github.com/dantemogrim/imir/cmd"
	"github.com/dantemogrim/imir/pkg/api"
	"github.com/dantemogrim/imir/pkg/styles"
	"github.com/gernest/wow"
	"github.com/gernest/wow/spin"
)

func main() {
	cmd.ParseFlags()

	spinner := wow.New(os.Stdout, spin.Get(spin.Star), " Stargazing.. 🔭")
	spinner.Start()

	isRetrograde, error := api.FetchMercuryRetrogradeStatus()

	time.Sleep(2 * time.Second)

	spinner.PersistWith(spin.Spinner{Frames: []string{""}}, "")

	// Clear screen
	fmt.Printf("\x1bc")

	if error != nil {
		errorMessage := styles.ApiError("Wasn't able to connect to the Mercury Retrograde API. Status: " + error.Error())
		fmt.Println(errorMessage)
		return
	}

	fmt.Printf("\x1bc")

	status := styles.Backward("Mercury is in retrograde!")

	if !isRetrograde {
		status = styles.Forward("Mercury is not in retrograde.")
	}

	fmt.Println(status)
}
