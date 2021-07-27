package testactions

import (
	"fmt"
	"gopkg.in/yaml.v2"
)

func Asdf() {
	_, _ = yaml.Marshal("asdf")
	fmt.Printf("hello go")
}
