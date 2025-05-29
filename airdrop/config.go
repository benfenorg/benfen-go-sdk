package main

type Config struct {
	Node struct {
		Rpc              string  `yaml:"rpc"`
		Amount           float64 `yaml:"amount"`
		Address          string  `yaml:"address"`
		M1Mnemonic       string  `yaml:"m1mnemonic"`
		OwltoFromAddress string  `yaml:"owlto_from_address"`
	}
}
