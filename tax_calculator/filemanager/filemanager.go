package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"time"
)

type Filemanger struct {
	Inputfile  string
	Outputfile string
}

func New(inputpath, outputpath string) Filemanger {
	return Filemanger{
		Inputfile:  inputpath,
		Outputfile: outputpath,
	}
}

func (fm Filemanger) ReadFile() ([]string, error) {
	file, err := os.Open(fm.Inputfile)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	var sprices []string
	for scanner.Scan() {
		sprices = append(sprices, scanner.Text())
	}
	defer file.Close()
	err = scanner.Err()
	if err != nil {
		panic(err)
		return nil, err
	}
	return sprices, nil
}

// we can use any instead of interface{}
func (wf Filemanger) WriteJSONFile(data interface{}) error {
	file, err := os.Create(wf.Outputfile)
	if err != nil {
		errors.New("Couldn't create a new file")
		os.Exit(1)
	}
	encoder := json.NewEncoder(file)
	time.Sleep(5 * time.Second)
	err = encoder.Encode(data)
	if err != nil {
		file.Close()
		errors.New("Failed to convert data")
	}
	file.Close()
	return nil
}
