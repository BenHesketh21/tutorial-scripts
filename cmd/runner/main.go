package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"os"

	"github.com/BenHesketh21/tutorial-scripts/pkg/tutorial"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.DebugLevel)
	log.Debug("Application Starting")

	configFile := flag.String("config-file", ".config", "Config file containing the information for a tutorial.")
	flag.Parse()
	log.Debug(*configFile)
	res := tutorial.Tutorial{}
	file, err := os.Open(*configFile)

	if err != nil {
		log.Error("Unable to open config file")
		log.Fatal(err)
	}

	byteValue, err := ioutil.ReadAll(file)

	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(byteValue, &res)

	for _, p := range res.Prerequisites {
		if p.Checked {
			continue
		}
		available, err := tutorial.IsPrerequisiteAvailable(p)
		p.Checked = true
		if available {
			log.Infof("Prerequisite: %s is avaiable", p.Name)
		} else {
			exists, position, err := tutorial.DoesAlternativePrerequisiteExist(res, p.Alternative)
			if err != nil {
				log.Fatal(err)
			}
			if exists {
				available, err := tutorial.IsPrerequisiteAvailable(res.Prerequisites[position])
				res.Prerequisites[position].Checked = true
				if err != nil {
					log.Fatal(err)
				}
				if !available {
					log.Fatalf("Cannot find prerequisite %s, or the alternative %s, make sure one is installed and on the PATH", p.Name, res.Prerequisites[position].Name)
				}
			} else {
				log.Fatalf("Cannot find prerequisite %s, make sure it is installed and on the PATH", p.Name)
			}
		}
		if err != nil {
			log.Fatal(err)
		}
	}

	for i, s := range res.Steps {
		tutorial.ExecuteStep(s, i)
	}

}
