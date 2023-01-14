package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var months = [12]string{
	"Janeiro",
	"Fevereiro",
	"Março",
	"Abril",
	"Maio",
	"Junho",
	"Julho",
	"Agosto",
	"Setembro",
	"Outubro",
	"Novembro",
	"Dezembro",
}

func listFiles(dirPath string) (FileIndex, error) {
	reader, err := os.Open(dirPath)

	if err != nil {
		return nil, err
	}

	files, err := reader.Readdir(0)

	if err != nil {
		return nil, err
	}

	var fileIndex FileIndex

	for _, file := range files {
		if !file.IsDir() {
			fileInfo := FileInfo{
				Name: file.Name(),
				LastModified: LastModified{
					Day:   fmt.Sprint(file.ModTime().Day()),
					Month: months[file.ModTime().Month()-1],
					Year:  fmt.Sprint(file.ModTime().Year()),
				},
			}

			fileIndex = append(fileIndex, fileInfo)
		}
	}

	return fileIndex, nil
}

func createFolder(dirPath string) error {
	return os.MkdirAll("./"+dirPath, os.ModePerm)
}

func moveFile(sourcePath string, destinationPath string) error {
	return os.Rename("./"+sourcePath, "./"+destinationPath)
}

func organizeFiles(sortBy string, dirPath string) error {
	fileIndex, err := listFiles(dirPath)

	if err != nil {
		return err
	}

	if sortBy == "-d" {
		for _, fileInfo := range fileIndex {
			createFolder(filepath.Join(
				fileInfo.LastModified.Year,
				fileInfo.LastModified.Month,
				fileInfo.LastModified.Day,
			))
			moveFile(
				fileInfo.Name,
				filepath.Join(
					fileInfo.LastModified.Year,
					fileInfo.LastModified.Month,
					fileInfo.LastModified.Day,
					fileInfo.Name,
				),
			)
		}
	} else if sortBy == "-n" {
		for index, fileInfo := range fileIndex {
			title := strings.ToUpper(strings.Split(fileInfo.Name, "_")[0])

			splittedTitle := strings.Split(title, " ")

			keyword := strings.ToUpper(splittedTitle[0])

			fileIndex[index].Keyword = keyword
		}

		if err != nil {
			return err
		}

		for _, fileInfo := range fileIndex {
			createFolder(fileInfo.Keyword)

			oldPath := fileInfo.Name
			newPath := filepath.Join(fileInfo.Keyword, fileInfo.Name)

			err := moveFile(oldPath, newPath)

			if err != nil {
				return err
			}
		}
	} else {
		return errors.New("argumentos incorretos")
	}

	return nil
}
