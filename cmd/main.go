package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/pavelanni/pkgname"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		p, err := pkgname.Parse(line)
		if err != nil {
			log.Println("error: ", err, " in pkg: ", line)
		} else {
			fmt.Printf("%s,%s,%s\n", p.Name, p.Version, p.Arch)
		}
	}
}
