package main

type Config struct {
	Node struct {
		Path         []string `yaml:"path"`
		MaxDataFolds int      `yaml:"maxDataFolds"`
	}
}
