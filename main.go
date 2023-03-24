package main

import (
	"bufio"
	"github.com/nlsrmll/logger"
	"log"
)

func main() {
	logger
	//Einen text zum Dokument hinzufügen
	text := "Neue Zeile hinzufügen\n"
	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(text)
	if err != nil {
		log.Fatalln(err)
	}

	err = writer.Flush()
	if err != nil {
		log.Fatal(err)
	}

}
