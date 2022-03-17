package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/kortschak/ct"
)

func main() {
	green := ct.Fg(ct.BoldGreen).Paint
	yellow := ct.Fg(ct.BoldYellow).Paint
	dec := json.NewDecoder(os.Stdin)
	for {
		var line json.RawMessage
		err := dec.Decode(&line)
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
		os.Stdout.Write(line)
		fmt.Println()
		var entry struct {
			LogOrigin struct {
				File string `json:"file.name"`
				Line int    `json:"file.line"`
			} `json:"log.origin"`
			Message string
		}
		err = json.Unmarshal(line, &entry)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s:%d\n", yellow(entry.LogOrigin.File), yellow(entry.LogOrigin.Line))
		for _, l := range strings.Split(entry.Message, "\n") {
			fmt.Printf("\t%s\n", green(l))
		}
		fmt.Println()
	}
}
