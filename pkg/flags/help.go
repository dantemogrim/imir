package flags

import (
	"flag"
	"fmt"
	"os"
)

func helpOption() {
	fmt.Println("Usage:")
	flag.PrintDefaults()
	os.Exit(0)
}
