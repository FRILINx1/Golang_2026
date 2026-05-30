package main

import (
	"fmt"
	"reflect"
	"strings"
)

type Server struct {
	Host       string   `json:"host"`
	Port       int      `json:"port"`
	Debug      bool     `json:"debug"`
	AllowedIPs []string `json:"allowed_ips"`
}

func ToYAML(v any) (string, error) {
	return toYAMLValue(reflect.ValueOf(v), 0)
}

func toYAMLValue(val reflect.Value, depth int) (string, error) {
	if val.Kind() == reflect.Ptr {
		if val.IsNil() {
			return "null", nil
		}
		val = val.Elem()
	}

	indent := strings.Repeat("  ", depth)

	switch val.Kind() {
	case reflect.String:
		return fmt.Sprintf(`"%s"`, val.String()), nil
		
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return fmt.Sprintf(`%d`, val.Int()), nil
		
	case reflect.Bool:
		return fmt.Sprintf(`%t`, val.Bool()), nil
		
	case reflect.Slice, reflect.Array:
		if val.Len() == 0 {
			return "[]", nil
		}
		var sb strings.Builder
		for i := 0; i < val.Len(); i++ {
			sb.WriteString("\n" + indent)
			
			elemStr, err := toYAMLValue(val.Index(i), depth+1)
			if err != nil {
				return "", err
			}
			
			sb.WriteString("- " + elemStr)
		}
		return sb.String(), nil
		
	case reflect.Struct:
		var sb strings.Builder
		t := val.Type()
		
		for i := 0; i < val.NumField(); i++ {
			field := t.Field(i)
			fieldVal := val.Field(i)
			
			tag := field.Tag.Get("json")
			if tag == "" {
				tag = field.Name
			}
			
			valStr, err := toYAMLValue(fieldVal, depth+1)
			if err != nil {
				return "", err
			}
			
			if valStr != "" && valStr[0] == '\n' {
				sb.WriteString(indent + tag + ":" + valStr)
			} else {
				sb.WriteString(indent + tag + ": " + valStr)
			}
			
			if i < val.NumField()-1 {
				sb.WriteString("\n")
			}
		}
		return sb.String(), nil
		
	default:
		return "", fmt.Errorf("непідтримуваний тип: %s", val.Kind())
	}
}

func ToJSON(v any) (string, error) {
	return toJSONValue(reflect.ValueOf(v), 1)
}

func toJSONValue(val reflect.Value, depth int) (string, error) {
	if val.Kind() == reflect.Ptr {
		if val.IsNil() {
			return "null", nil
		}
		val = val.Elem()
	}

	indent := strings.Repeat("\t", depth)
	prevIndent := strings.Repeat("\t", depth-1)

	switch val.Kind() {
	case reflect.String:
		return fmt.Sprintf(`"%s"`, val.String()), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return fmt.Sprintf(`%d`, val.Int()), nil
	case reflect.Bool:
		return fmt.Sprintf(`%t`, val.Bool()), nil
	case reflect.Slice, reflect.Array:
		if val.Len() == 0 {
			return "[]", nil
		}
		var sb strings.Builder
		sb.WriteString("[\n")
		for i := 0; i < val.Len(); i++ {
			sb.WriteString(indent)
			elemStr, _ := toJSONValue(val.Index(i), depth+1)
			sb.WriteString(elemStr)
			if i < val.Len()-1 {
				sb.WriteString(",")
			}
			sb.WriteString("\n")
		}
		sb.WriteString(prevIndent + "]")
		return sb.String(), nil
	case reflect.Struct:
		var sb strings.Builder
		sb.WriteString("{\n")
		t := val.Type()
		for i := 0; i < val.NumField(); i++ {
			field := t.Field(i)
			tag := field.Tag.Get("json")
			if tag == "" {
				tag = field.Name
			}
			sb.WriteString(indent + `"` + tag + `": `)
			valStr, _ := toJSONValue(val.Field(i), depth+1)
			sb.WriteString(valStr)
			if i < val.NumField()-1 {
				sb.WriteString(",")
			}
			sb.WriteString("\n")
		}
		sb.WriteString(prevIndent + "}")
		return sb.String(), nil
	default:
		return "", fmt.Errorf("unsupported type")
	}
}

func main() {
	srv := Server{
		Host:       "localhost",
		Port:       8080,
		Debug:      true,
		AllowedIPs: []string{"192.168.1.1", "10.0.0.1"},
	}

	yamlResult, _ := ToYAML(srv)
	fmt.Println(" YAML ")
	fmt.Println(yamlResult)
}