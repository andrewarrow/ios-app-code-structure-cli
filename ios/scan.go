package ios

import (
	"fmt"
	"io/ioutil"
)

func Scan(dir string) {
	files, _ := ioutil.ReadDir(dir)
	for _, file := range files {
		name := file.Name()
		fmt.Println(name)
	}
}
