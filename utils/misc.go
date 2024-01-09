package utils

import (
	"encoding/hex"
	"fmt"
	"reflect"
	"strings"
)

func NoErr(err error) {
	if err != nil {
		panic(err)
	}
}

func IndentText(text, indent string) string {
	return strings.ReplaceAll(text, "\n", "\n"+indent)
}

func MarshalTextIndex(v interface{}, indentChar string) string {
	return formatValueText(reflect.ValueOf(v), 0, indentChar)
}

func MarshalText(v interface{}) string {
	return MarshalTextIndex(v, "	")
}

func formatValueText(v reflect.Value, indent int, indentChar string) string {
	var sb strings.Builder

	switch v.Kind() {
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			field := v.Type().Field(i)
			fieldValue := v.Field(i)

			sb.WriteString(strings.Repeat(indentChar, indent))
			sb.WriteString(field.Name)
			sb.WriteString(": ")

			if fieldValue.Kind() == reflect.Struct || (fieldValue.Kind() == reflect.Ptr && !fieldValue.IsNil() && fieldValue.Elem().Kind() == reflect.Struct) {
				sb.WriteString("\n")
				sb.WriteString(formatValueText(fieldValue, indent+1, indentChar)) // Increment indent for nested structs
			} else {
				sb.WriteString(fmt.Sprintf("%v\n", formatIndividualValue(fieldValue)))
			}
		}
	case reflect.Ptr:
		if v.IsNil() {
			sb.WriteString("nil\n")
		} else {
			sb.WriteString(formatValueText(v.Elem(), indent, indentChar))
		}

	case reflect.Slice:
		// if bytes array, print as hex or e, print as hex
		if v.Type().Elem().Kind() == reflect.Uint8 {
			sb.WriteString(fmt.Sprintf("%v\n", formatIndividualValue(v)))
		} else {
			for i := 0; i < v.Len(); i++ {
				sb.WriteString(formatValueText(v.Index(i), indent, indentChar))
			}
		}

	case reflect.Array:
		// if bytes array, print as hex or e, print as hex
		if v.Type().Elem().Kind() == reflect.Uint8 {
			sb.WriteString(fmt.Sprintf("%v\n", formatIndividualValue(v)))
		} else {
			for i := 0; i < v.Len(); i++ {
				sb.WriteString(formatValueText(v.Index(i), indent, indentChar))
			}
		}

	default:
		sb.WriteString(fmt.Sprintf("%v\n", v.Interface()))
	}

	return sb.String()
}

func formatIndividualValue(v reflect.Value) interface{} {
	if v.Kind() == reflect.Slice && v.Type().Elem().Kind() == reflect.Uint8 { // Check for []byte
		bytes := v.Bytes()
		return "0x" + hex.EncodeToString(bytes)
	}

	// handle byte arrays
	if v.Kind() == reflect.Array && v.Type().Elem().Kind() == reflect.Uint8 { // Check for [32]byte
		bytes := v.Slice(0, v.Len()).Bytes()
		return "0x" + hex.EncodeToString(bytes)
	}

	return v.Interface()
}
