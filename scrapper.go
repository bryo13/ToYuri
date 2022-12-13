package main

import (
	"fmt"
	"log"
	"os/exec"
)

type Scrap struct{}

func NewScrapper() *Scrap {
	return &Scrap{}
}

func (s *Scrap) IsOnline() bool {

	out, err := exec.Command("ping", "8.8.8.8").Output()
	fmt.Println(out)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
