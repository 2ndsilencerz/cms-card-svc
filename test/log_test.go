package test

import (
	"bufio"
	"os"
	"testing"

	"github.com/2ndsilencerz/cms-card-svc/configs/utils"
)

func TestLog(t *testing.T) {
	contents := "log testing"
	utils.LogToFile(contents)
	contents = utils.InstantTimeString() + " " + contents

	file, err := os.OpenFile("logs/log", os.O_RDONLY, 0777)
	if err != nil {
		t.Errorf("Error: %v", err)
		return
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	var contentGotten string
	for fileScanner.Scan() {
		contentGotten = fileScanner.Text()
	}

	if contentGotten != contents {
		t.Errorf("Expected %v got %v", contents, contentGotten)
	}

	err = file.Close()
	if err != nil {
		t.Errorf("Error: %v", err)
	}
}
