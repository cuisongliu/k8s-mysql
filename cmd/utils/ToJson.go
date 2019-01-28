package utils

import (
	"bytes"
	"fmt"
	"io"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/yaml"
)

var bs = []byte(`
kind: Namespace
metadata:
   name: test
---
kind: bbb
name: aaa`)

/* Out put
{"kind":"Namespace","metadata":{"name":"test"}}
{"kind":"bbb","name":"aaa"}
*/
func ToJson() {
	reader := bytes.NewReader(bs)
	ext := runtime.RawExtension{}
	d := yaml.NewYAMLOrJSONDecoder(reader, 4096)
	for {
		if err := d.Decode(&ext); err != nil {
			if err == io.EOF {
				return
			}
			return
		}
		fmt.Println(string(ext.Raw))
	}
}

//YamlCallback is
type YamlCallback func([]byte) error

//YamlHandler is
func YamlHandler(rawBytes []byte, fn YamlCallback) (err error) {
	reader := bytes.NewReader(rawBytes)
	ext := runtime.RawExtension{}
	d := yaml.NewYAMLOrJSONDecoder(reader, 4096)
	for {
		if err = d.Decode(&ext); err != nil {
			if err == io.EOF {
				return nil
			}
			return fmt.Errorf("decode yaml json failed: %v", err)
		}
		//Raw is already to json
		if err = fn(ext.Raw); err != nil {
			return fmt.Errorf("handler yaml callback fn failed: %v", err)
		}
	}
}
