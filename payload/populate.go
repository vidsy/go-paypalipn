package payload

import (
	"fmt"
	"reflect"
	"strconv"
)

type (
	// Populate takes a reflect.Type, reflect.Value, form data and optional key
	// index and populates the reflected struct with the form data.
	Populate struct {
	}

	// Processor interface for Populaste structs that process data.
	Processor interface {
		Process(itemType reflect.Type, itemValue reflect.Value, formData *Values, fieldIndex int, countField string) []error
	}
)

// Process takes a structs type and value and recursively populates it based on the data
// in formData. 'fieldIndex' is the index for the given field being used to populate a struct with,
// for example 1 would get the data from fields `qux_1, foo_1, bar_1` etc. 'countField' is
// the field that is used to determine the amount of array items in a response.
func (p *Populate) Process(itemType reflect.Type, itemValue reflect.Value, formData *Values, fieldIndex int, countField string) []error {
	itemCount := formData.ItemCount(countField)
	var errors []error

	for i := 0; i < itemType.Elem().NumField(); i++ {
		field := itemType.Elem().Field(i)
		fieldValue := itemValue.Elem().FieldByName(field.Name)

		if fieldTag, ok := field.Tag.Lookup("form_field"); ok {
			fieldData := formData.Get(fieldTag)
			if fieldIndex != -1 {
				fieldData = formData.GetValueAtIndex(fieldTag, fieldIndex)
			}

			switch field.Type.Kind() {
			case reflect.String:
				fieldValue.SetString(fieldData)

			case reflect.Float64:
				floatValue, err := strconv.ParseFloat(fieldData, 64)
				if err != nil {
					errors = append(errors, fmt.Errorf("Error encountered parsing float value: '%s'", fieldData))
				}
				fieldValue.SetFloat(floatValue)

			case reflect.Slice:
				sliceValue := itemValue.Elem().Field(i)

				switch sliceValue.Type().Elem().Kind() {
				case reflect.String:
					formFields := formData.GetValues(fieldTag, itemCount)
					for _, value := range formFields {
						itemValue := reflect.Indirect(reflect.New(sliceValue.Type().Elem()))
						itemValue.SetString(value)
						sliceValue.Set(reflect.Append(sliceValue, itemValue))
					}

				case reflect.Struct:
					for ii := 0; ii < itemCount; ii++ {
						itemValue := reflect.New(sliceValue.Type().Elem())
						itemType := sliceValue.Type()
						errors = p.Process(itemType, itemValue, formData, ii, countField)
						sliceValue.Set(reflect.Append(sliceValue, itemValue.Elem()))
					}
				}
			}
		}
	}

	return errors
}
