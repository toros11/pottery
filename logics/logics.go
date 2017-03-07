package logics

import (
	"fmt"
	"github.com/qb0C80aE/clay/extension"
	"text/template"
)

func HookSubmodules() {
}

func init() {
	funcMap := template.FuncMap{
		"add": func(a, b int) int { return a + b },
		"sub": func(a, b int) int { return a - b },
		"mul": func(a, b int) int { return a * b },
		"div": func(a, b int) int { return a / b },
		"map": func(pairs ...interface{}) (map[interface{}]interface{}, error) {
			if len(pairs)%2 == 1 {
				return nil, fmt.Errorf("Number of arguments must be even")
			}
			m := make(map[interface{}]interface{}, len(pairs)/2)
			for i := 0; i < len(pairs); i += 2 {
				m[pairs[i]] = pairs[i+1]
			}
			return m, nil
		},
		"putToMap": func(target map[interface{}]interface{}, key interface{}, value interface{}) map[interface{}]interface{} {
			target[key] = value
			return target
		},
		"getFromMap": func(target map[interface{}]interface{}, key interface{}) interface{} {
			value := target[key]
			return value
		},
		"deleteFromMap": func(target map[interface{}]interface{}, key interface{}) map[interface{}]interface{} {
			delete(target, key)
			return target
		},
	}
	extension.RegisterTemplateFuncMap(funcMap)
}
