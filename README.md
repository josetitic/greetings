# greetings
Este paquete proporciona una forma simple de obtener saludos personales en go

# Instalación
Ejecuta el siguiente comando para instalar el paquete:
```bash

go get -u github.com/josetitic/greetings

```
# Uso
Ejm:

```go
package main

import (
	"fmt"
	"log"
	"github.com/josetitic/greetings"
)

func main() {
	log.SetPrefix("greetings")
	log.SetFlags(0)

	names := []string{"Sami","Mary","José","Juan"}

	message,err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)

	message,err := greetings.Hello("Sami")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
```
