package format

import (
	"fmt"
	"reflect"
	"strings"
)

type InvalidFieldPath struct {
	Field string
}

func (e *InvalidFieldPath) Error() string {
	return fmt.Sprintf("field path contains invalid field %s", e.Field)
}

type NameAndFieldPathsUnmatched struct {
	NameLen      int
	FieldPathLen int
}

func (e *NameAndFieldPathsUnmatched) Error() string {
	return fmt.Sprintf("name %d and field path %d slices do not match", e.NameLen, e.FieldPathLen)
}

type InvalidFieldPathSlice struct {
	Field string
}

func (e *InvalidFieldPathSlice) Error() string {
	return fmt.Sprintf("field path contains slice at field %s", e.Field)
}

type InvalidSliceType struct {
	Value reflect.Type
}

func (e *InvalidSliceType) Error() string {
	return fmt.Sprintf("invalid list type of %s provided", e.Value.String())
}

// GetFieldValue - Returns the value of a field identified by it's field path
func GetFieldValue(item interface{}, fieldPath []string) (value interface{}, err error) {

	if len(fieldPath) == 0 {
		return item, nil
	}

	if reflect.TypeOf(item).Kind() == reflect.Slice {
		return nil, &InvalidFieldPathSlice{Field: fieldPath[0]}
	}

	reflectValue := reflect.ValueOf(item)
	fieldValue := reflect.Indirect(reflectValue).FieldByName(fieldPath[0])
	if !fieldValue.IsValid() {
		return nil, &InvalidFieldPath{Field: fieldPath[0]}
	}

	value, err = GetFieldValue(fieldValue.Interface(), fieldPath[1:])
	if err != nil {
		return nil, err
	}

	return value, nil
}

// FormattedList - Returns a formatted string representation of a list
// of structs with selected fields represented as columns. Fields names
// and fieldPaths must align. Field paths specify the nested fields of
// a struct in the format a.b.c
func FormattedList(list interface{}, fieldNames []string, fieldPaths []string) (output string, err error) {

	if len(fieldNames) != len(fieldPaths) {
		return "", &NameAndFieldPathsUnmatched{NameLen: len(fieldNames), FieldPathLen: len(fieldPaths)}
	}

	// Format list interface to slice
	itemSlice := make([]interface{}, 0)

	switch reflect.TypeOf(list).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(list)

		for i := 0; i < s.Len(); i++ {
			itemSlice = append(itemSlice, s.Index(i).Interface())
		}
	default:
		return "", &InvalidSliceType{Value: reflect.TypeOf(list)}
	}

	// Check for empty slice
	if len(itemSlice) == 0 {
		return "no resources found", nil
	}

	// Initialise columns
	type column struct {
		header    string
		width     int
		fieldPath string
		rows      []string
	}
	columns := []*column{}
	for i, fieldName := range fieldNames {
		columns = append(columns, &column{
			header:    fieldName,
			width:     len(fieldName),
			fieldPath: fieldPaths[i],
			rows:      []string{},
		})
	}

	// Process items to columns
	for _, item := range itemSlice {
		for _, column := range columns {

			fields := strings.Split(column.fieldPath, ".")
			value, err := GetFieldValue(item, fields)
			if err != nil {
				return "", err
			}
			valueStr := fmt.Sprintf("%v", value)
			valueWidth := len(valueStr)

			// Increase column size if necessary
			if column.width < valueWidth {
				column.width = valueWidth
			}

			// Add row
			column.rows = append(column.rows, valueStr)
		}
	}

	// Format column headers
	totalWidth := 0
	output = "\n"
	for _, col := range columns {
		col.width += 3
		totalWidth += col.width
		dif := col.width - len(col.header)
		output += col.header + strings.Repeat(" ", dif)
	}
	output += "\n" + strings.Repeat("─", totalWidth) + "\n"

	// Format column rows
	for i := 0; i < len(itemSlice); i++ {
		for j, col := range columns {
			dif := col.width - len(col.rows[i])
			output += columns[j].rows[i] + strings.Repeat(" ", dif)
		}
		output += "\n"
	}

	return output, nil
}

// StandardiseWhitespace - helper function to set whitespace to single space
func StandardiseWhitespace(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

// RemoveWhitespace - helper function to remove all whitespace
func RemoveWhitespace(s string) string {
	return strings.ReplaceAll(strings.Join(strings.Fields(s), " "), " ", "")
}
