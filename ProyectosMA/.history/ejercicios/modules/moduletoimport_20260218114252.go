package modules

import (
	"fmt"

	m "modules"

	"github.com/michelidigoraz/saludos/saludos"
)

func Greeting(name string) string {
	saludos.Saludos()
	m.Greeting()
	return fmt.Sprintf("Hi %s", name)

}
