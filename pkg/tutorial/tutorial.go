package tutorial

import (
	"fmt"
	"os/exec"
	"strings"
	"time"

	color "github.com/fatih/color"

	"github.com/BenHesketh21/tutorial-scripts/pkg/typing"
	log "github.com/sirupsen/logrus"
)

type Prerequisite struct {
	Name                       string `json:"name"`
	CheckInstallVersionCommand string `json:"checkInstallVersionCommand"`
	Version                    string `json:"version"`
	Alternative                string `json:"alternative"`
	Checked                    bool
}

type Step struct {
	BeforeCommandMessage string `json:"beforeCommandMessage"`
	Command              string `json:"command"`
	AfterCommandMessage  string `json"afterCommandMessage"`
}

type Tutorial struct {
	Name          string         `json:"name"`
	Prerequisites []Prerequisite `json:"prerequisites"`
	Steps         []Step         `json:"steps"`
}

func IsPrerequisiteAvailable(prerequisite Prerequisite) (bool, error) {
	log.Debugf("Checking is %s is installed", prerequisite.Name)
	installVersionCommand := exec.Command("bash", "-c", prerequisite.CheckInstallVersionCommand)
	output, err := installVersionCommand.Output()
	if err != nil {
		log.Warnf("Cannot find %s: %s", prerequisite.Name, err.Error())
		return false, nil
	}
	log.Debug(string(output))
	if strings.Contains(string(output), prerequisite.Version) {
		log.Debugf("%s: %s is installed", prerequisite.Name, prerequisite.Version)
		return true, nil
	}
	return false, nil
}

func DoesAlternativePrerequisiteExist(tutorial Tutorial, alternative string) (bool, int, error) {
	for i, p := range tutorial.Prerequisites {
		if p.Name == alternative {
			return true, i, nil
		}
	}
	return false, 0, nil
}

func waiter() {
	i := 0
	for {
		log.Debug(i)
		time.Sleep(time.Second * 1)
		if i == 5 {
			fmt.Print(" (Press <Enter> to execute command)")
			return
		}
		i++
	}
}

func ExecuteStep(step Step, stepNumber int) {
	color.Magenta(fmt.Sprint(stepNumber) + ": " + step.BeforeCommandMessage)
	fmt.Println("")
	fmt.Println("")
	typing.SimulateType(step.Command, 50)
	go waiter()
	fmt.Scanln()
	cmd := exec.Command("bash", "-c", step.Command)
	output, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(string(output))
	time.Sleep(time.Second * 1)
	fmt.Println("")
	fmt.Println("")
	color.Red(step.AfterCommandMessage)
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	time.Sleep(time.Second * 1)
}
