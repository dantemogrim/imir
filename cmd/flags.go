package cmd

import (
	"fmt"
	"os"

	// "reflect"
	"time"

	"github.com/dantemogrim/imir/pkg/api"
	"github.com/dantemogrim/imir/pkg/messages"
	"github.com/dantemogrim/imir/pkg/styles"
	"github.com/integrii/flaggy"
)

var (
    dateFlag string
    aboutFlag = false
)

func ParseFlags() {
    flaggy.DefaultParser.ShowVersionWithVersionFlag = false
    flaggy.DefaultParser.ShowHelpOnUnexpected = true

    flaggy.SetName("imir")
    flaggy.SetDescription("Is Mercury In Retrograde?")

    flaggy.String(&dateFlag, "d", "date YYYY-MM-DD", "Check if Mercury was in retrograde during a specific date.")
    flaggy.Bool(&aboutFlag, "a", "about", "Show information on the concept behind Mercury Retrograde.")

    flaggy.Parse()

    if dateFlag != ""{
        // Access and store the 2nd argument (excluding program name & date flag).
        userDateInput := os.Args[2]

        // Convert the user input to a time.Time object (to validate).
        parsedUserDate, dateError := time.Parse("2006-01-02", userDateInput)

        if dateError != nil {
            errorMessage := styles.ApiError("There was an issue when validating the given date.\nError: " + dateError.Error())
            usageMessage := "imir --date <yyyy-mm-dd>\nExample: $ imir --date 2018-04-28"
            fmt.Println(errorMessage)
            fmt.Println(usageMessage)
            os.Exit(1)
        }

        // Format the date back to a string (to later concatenate with API endpoint).
        formattedDateString := parsedUserDate.Format("2006-01-02")

       // fmt.Println("Type of formattedDateString: ", reflect.TypeOf(formattedDateString))
        fetchResult, fetchError := api.FetchByDateInput(formattedDateString)

        if fetchError != nil {
            fmt.Println("error")
            os.Exit(1)
        }


        message := messages.DatedResult(parsedUserDate, fetchResult)
        fmt.Println(message)
        os.Exit(0)
	}

	if aboutFlag {
		fmt.Println("Used about flag.")
        // TODO import about pager.
        os.Exit(0)
	}

}
