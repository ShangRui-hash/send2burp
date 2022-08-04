package debug

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ShangRui-hash/send2burp/config"
)

func Println(a ...interface{}) (n int, err error) {
	if !config.CurrentConf.Debug {
		return 0, nil
	}
	return fmt.Fprintln(os.Stderr, a...)
}

func WriteFile(filename string, a ...interface{}) error {
	if !config.CurrentConf.Debug {
		return nil
	}
	return ioutil.WriteFile(filename, []byte(fmt.Sprintf("%s", a...)), 0644)
}
