package main

import (
	"fmt"
	"reflect"
)

func convertMapToStructure(m map[string]interface{}, structure interface{}) error {
	structValue := reflect.ValueOf(structure).Elem()
	structType := structValue.Type()

	// if structValue.Kind() == reflect.Interface {
	// 	return fmt.Errorf("interface value idk what the fuck")
	// }

	// fmt.Println()
	numFields := structValue.NumField()
	for i := 0; i < numFields; i++ {
		structField := structType.Field(i)
		f, ok := m[structField.Tag.Get("json")]
		if !ok {
			return fmt.Errorf("field name: %s isn't present in the map", structField.Name)
		}

		var value reflect.Value
		kind := structField.Type.Kind()

		switch kind {
		case reflect.Struct:
			strPtr := reflect.New(structField.Type)
			tmp := strPtr.Interface()
			panikIfErr(convertMapToStructure(f.(map[string]interface{}), tmp))
			value = reflect.Indirect(strPtr)
		case reflect.Array:
			fallthrough
		case reflect.Slice:
			items := reflect.Indirect(reflect.New(reflect.SliceOf(structField.Type.Elem())))
			for _, s := range f.([]map[string]interface{}) {
				strPtr := reflect.New(structField.Type.Elem())

				switch structField.Type.Elem().Kind() {
				case reflect.Interface:
					fmt.Println(structField)
					fmt.Println(f)
					crash("interface not implemented")
				default:
					crashIfErr(convertMapToStructure(s, strPtr.Interface()))
					value = reflect.Indirect(strPtr)
					items = reflect.Append(items, value)
				}
			}
			value = items
		default:
			value = reflect.ValueOf(f)
		}

		if ok {
			field := structValue.Field(i)
			field.Set(value)
		}
	}

	return nil
}
