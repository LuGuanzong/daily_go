package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

var Conf Config

// LoadConfFromYaml 配置配置文件
func LoadConfFromYaml(filename string) {
	fmt.Printf("开始读取配置文件 %s", filename)
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("configFile:%s\n", err)
		fmt.Printf("Load config [%s] err: %s\n", filename, err)
		os.Exit(1)
	}
	defer closeFile(file)

	decode := yaml.NewDecoder(file)
	err = decode.Decode(&Conf)
	if err != nil {
		fmt.Printf("yaml decode error, err:%s\n", err)
		os.Exit(1)
	}
}

func closeFile(file *os.File) {
	_ = file.Close()
}
