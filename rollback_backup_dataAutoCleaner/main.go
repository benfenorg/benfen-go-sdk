package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"time"
)

var basePath string
var maxDataFolds = 11

func main() {
	config, err := readConfigFile()
	if err != nil {
		return
	}
	maxDataFolds = config.Node.MaxDataFolds

	sortDataFolds(config.Node.Path[0])
	// create a ticker： 60*60 * time.Second, 1 hour
	ticker := time.NewTicker(60 * 60 * time.Second)
	// start goroutine todo something
	go func() {
		for {
			select {
			case <-ticker.C:

				for _, path := range config.Node.Path {
					fmt.Println("schedule start ...")
					sortDataFolds(path)
					fmt.Println("schedule end ...")
				}

			}
		}
	}()

	// main goroutine waiting forever
	select {}

}

func readConfigFile() (*Config, error) {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("can not get working path .. ", err)
		return nil, err
	}

	fmt.Println("current working path：", wd)

	data, err := ioutil.ReadFile("./rollback_backup_dataAutoCleaner/config.yaml")
	if err != nil {
		log.Fatalf("can not read file ：%v", err)
	}

	// parse config file.
	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("can not decode file ：%v", err)
	}
	return &config, nil
	// using config file
}

func sortDataFolds(path string) {
	println("start clean data ......")
	root := path // current directory

	subdirectories, err := getSubdirectories(root)
	if err != nil {
		fmt.Println("can not read sub directors:", err)
		return
	}

	sortSubdirectoriesByTime(subdirectories)

	count := 0
	for _, subdir := range subdirectories {
		fmt.Println(subdir)
		count = count + 1
	}

	deleteCount := 0
	if count >= maxDataFolds {
		needDeleteCount := count - maxDataFolds
		for _, subdir := range subdirectories {
			fmt.Println(subdir)

			if deleteCount < needDeleteCount {
				deleteCount = deleteCount + 1
				err := os.RemoveAll(subdir)
				if err != nil {
					fmt.Println("can not remove subdir :", err)
					return
				}

				fmt.Println("remove successfully.....")
			}
		}
	}

}

// get sub directories
func getSubdirectories(root string) ([]string, error) {
	var subdirectories []string

	files, err := ioutil.ReadDir(root)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if file.IsDir() {
			subdirectories = append(subdirectories, filepath.Join(root, file.Name()))
		}
	}

	return subdirectories, nil
}

// sort sub directories by time
func sortSubdirectoriesByTime(subdirectories []string) {
	sort.Slice(
		subdirectories, func(i, j int) bool {
			info1, _ := os.Stat(subdirectories[i])
			info2, _ := os.Stat(subdirectories[j])
			return info1.ModTime().Before(info2.ModTime())
		},
	)
}
