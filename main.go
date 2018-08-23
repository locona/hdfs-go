package main

import (
	"os"

	"github.com/colinmarc/hdfs"
	"github.com/k0kubun/pp"
)

func main() {
	client, err := hdfs.New("localhost:8020")
	// b, err := client.ReadFile("/tmp/sample_1_1_1/0_0_0.csv")
	// pp.Println(string(b), err)
	err = client.Walk("/tmp/samplesamplesamplesamplesample", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			pp.Println(info.Name())
		}
		return nil
	})
	pp.Println(err)
	//
	// file, _ := client.Open("/mobydick.txt")
	//
	// buf := make([]byte, 59)
	// file.ReadAt(buf, 48847)
	//
	// fmt.Println(string(buf))
}
