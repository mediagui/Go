package moduletoimport

import (
	"fmt"

	"github.com/michelidigoraz/saludos/saludos"
)

func Greeting(name string) string {
	saludos.Saludos()
	return fmt.Sprintf("Hi %s", name)

}
