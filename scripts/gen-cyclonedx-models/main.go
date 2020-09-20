package main

import (
	"fmt"
	"github.com/xuri/xgen"
	"os"
)

func main() {
	err := gen_cyclonedx_models("cyclonedx", "./schema", ".")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func gen_cyclonedx_models(pkg string, inputPath string, outputPath string) error {
	if err := xgen.PrepareOutputDir(outputPath); err != nil {
		return nil
	}

	files, err := xgen.GetFileList(inputPath)
	if err != nil {
		return nil
	}

	for _, file := range files {
		if err = xgen.NewParser(&xgen.Options{
			FilePath:            file,
			OutputDir:           outputPath,
			Lang:                "Go",
			Package:             pkg,
			IncludeMap:          make(map[string]bool),
			LocalNameNSMap:      make(map[string]string),
			NSSchemaLocationMap: make(map[string]string),
			ParseFileList:       make(map[string]bool),
			ParseFileMap:        make(map[string][]interface{}),
			ProtoTree:           make([]interface{}, 0),
			RemoteSchema:        make(map[string][]byte),
		}).Parse(); err != nil {
			return fmt.Errorf("process error on %s: %s\r\n", file, err.Error())
		}
	}

	return nil
}
