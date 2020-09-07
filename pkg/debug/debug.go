package debug

import(
	"fmt"
	"log"
)

func Err(err string) error {
	return fmt.Errorf(err)
}

func PrintErrLog(smg string) {
	log.Printf("[ERROR]: %s\n", smg)
}

func Fatal(err error) {
	log.Fatal(err)
}
