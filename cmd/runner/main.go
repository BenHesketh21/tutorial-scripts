package main

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/BenHesketh21/tutorial-scripts/pkg/logger"
	"github.com/BenHesketh21/tutorial-scripts/pkg/tutorial"
	"github.com/BenHesketh21/tutorial-scripts/pkg/typing"
)

func main() {
	logger := logger.InitLogger(os.Stdout)
	logger.Info.Println("Application Starting")

	res := tutorial.Tutorial{}
	file, err := os.Open("examples/docker-images-tutorial.json")

	if err != nil {
		logger.Error.Fatal(err)
	}

	byteValue, err := ioutil.ReadAll(file)

	if err != nil {
		logger.Error.Fatal(err)
	}

	json.Unmarshal(byteValue, &res)

	typing.SimulateType("My name is ben and this is being automatically typed", 100)
}
