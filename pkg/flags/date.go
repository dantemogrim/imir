package flags

import (
	"fmt"
	"os"
	"time"

	"github.com/dantemogrim/imir/pkg/api"
	"github.com/dantemogrim/imir/pkg/messages"
	"github.com/dantemogrim/imir/pkg/styles"
)

func dateOption () {
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
