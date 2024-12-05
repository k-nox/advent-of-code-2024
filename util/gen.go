package util

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"strings"
	"text/template"
)

type Data struct {
	Day  int
	Part string
}

func Generate(d int) error {
	pkgName := fmt.Sprintf("day%02d", d)
	err := createInputs(pkgName)
	if err != nil {
		return fmt.Errorf("error creating inputs: %w", err)
	}

	err = genPkg(d, pkgName)
	if err != nil {
		return fmt.Errorf("error generating package: %w", err)
	}

	return nil
}

func createInputs(pkgName string) error {
	inputDir := fmt.Sprintf("input/%s", pkgName)
	err := createDirIfNotExist(inputDir)
	if err != nil {
		return err
	}

	err = createFileInDir(inputDir, "input.txt")
	if err != nil {
		return err
	}

	err = createFileInDir(inputDir, "sample.txt")
	if err != nil {
		return err
	}

	return nil
}

func createDirIfNotExist(name string) error {
	err := os.Mkdir(name, 0750)
	if err != nil && !errors.Is(err, fs.ErrExist) {
		return fmt.Errorf("failed to create %s dir: %w", name, err)
	}
	return nil
}

func createFileInDir(dir string, name string) error {
	fullPath := fmt.Sprintf("%s/%s", dir, name)
	f, err := os.Create(fullPath)
	if err != nil {
		return fmt.Errorf("error creating file %s: %w", fullPath, err)
	}
	defer f.Close()
	return nil
}

func genPkg(d int, pkgName string) error {
	err := createDirIfNotExist(pkgName)
	if err != nil {
		return err
	}

	for _, part := range []string{"One", "Two"} {
		err := genPartFile(d, pkgName, part)
		if err != nil {
			return err
		}
	}

	return nil
}

func genPartFile(d int, pkgName string, part string) error {
	tmplFile := "part.tmpl"
	tmpl, err := template.New(tmplFile).ParseFiles(fmt.Sprintf("util/%s", tmplFile))
	if err != nil {
		return fmt.Errorf("error parsing template: %w", err)
	}
	f, err := os.Create(fmt.Sprintf("%s/part%s.go", pkgName, strings.ToLower(part)))
	if err != nil {
		return fmt.Errorf("error creating part%s.go: %w", strings.ToLower(part), err)
	}
	defer f.Close()
	err = tmpl.Execute(f, Data{
		Day:  d,
		Part: part,
	})

	if err != nil {
		return fmt.Errorf("error execturing template: %w", err)
	}
	return nil
}
