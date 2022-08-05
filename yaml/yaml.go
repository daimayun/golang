package yaml

import (
	"github.com/daimayun/golang/file"
	y "gopkg.in/yaml.v3"
)

// NewYaml 实例化YAML文件
func NewYaml(yamlFilePath string, out interface{}) (err error) {
	var b []byte
	b, err = file.GetContentsToByte(yamlFilePath)
	if err != nil {
		return
	}
	err = y.Unmarshal(b, out)
	return
}
