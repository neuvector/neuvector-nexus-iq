package main

import (
	"fmt"
	"github.com/go-swagger/go-swagger/cmd/swagger/commands/generate"
	"github.com/jessevdk/go-flags"
	"log"
	"os"
)

func main() {
	log.SetOutput(os.Stdout)

	err := genNexusIqClient()
	if err != nil {
		log.Fatalln(err)
	}
}

func genNexusIqClient() error {
	tags := []string{
		"Application",
		"ThirdPartyScan",
		"Organization",
		"License",
	}

	models := []string{
		"Applications",
		"Application",
		"ApplicationTag",
		"NewApplication",
		"ScanStatusReference",
		"ScanStatus",
		"Violations",
		"Organizations",
		"Organization",
		"License",
	}

	args := []string{
		"--spec=./schema/nexus-iq-api-94.yml",
		"--target=.",
		"--model-package=models",
		"--client-package=client",
	}

	for _, tag := range tags {
		args = append(args, fmt.Sprintf("--tags=\"%s\"", tag))
	}

	for _, model := range models {
		args = append(args, fmt.Sprintf("--model=\"%s\"", model))
	}

	c := &generate.Client{}

	_, err := flags.ParseArgs(c, args)
	if err != nil {
		return err
	}

	err = c.Execute([]string{})
	if err != nil {
		return err
	}

	return nil
}
