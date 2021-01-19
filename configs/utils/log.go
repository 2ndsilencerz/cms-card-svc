package utils

import (
	"fmt"
	"log"
	"os"
)

func loadFile() *os.File {
	file, err := os.OpenFile("logs/log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		// file, err = os.Create("./logs/log")
		// if err != nil {
		panic(fmt.Errorf("Error: %v", err))
		// }
	}
	return file
}

// LogToFile used to log into file which located in logs/log
func LogToFile(content string) {
	file := loadFile()
	//writer := io.MultiWriter(os.Stdout, file)
	log.SetOutput(file)
	log.Print(content)
	content = fmt.Sprintln(InstantTimeString(), content)
	fmt.Print(content)

	err := file.Close()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

// PrintToConsole use timestamp as prefix
func PrintToConsole(content string) {
	fmt.Printf("%v %v\n", InstantTimeString(), content)
}
