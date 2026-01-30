package mcp

import (
	"encoding/json"
	"reflect"
)

// filterFields filters a struct or slice of structs to keep only specified fields.
// If fields is empty or nil, it returns a default set of fields (id, name/stem).
// supportedFields is used to validate requested fields (optional, if nil, all struct fields are allowed).
func filterFields(data interface{}, fields []string) interface{} {
	if data == nil {
		return nil
	}

	val := reflect.ValueOf(data)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	// Handle slice
	if val.Kind() == reflect.Slice {
		length := val.Len()
		result := make([]interface{}, length)
		for i := 0; i < length; i++ {
			result[i] = filterFields(val.Index(i).Interface(), fields)
		}
		return result
	}

	// Handle struct
	if val.Kind() == reflect.Struct {
		// Determine which fields to keep
		keepFields := make(map[string]bool)

		// Default fields logic
		if len(fields) == 0 {
			keepFields["id"] = true
			if hasField(val, "Name") {
				keepFields["name"] = true
			} else if hasField(val, "Stem") {
				keepFields["stem"] = true
			}
		} else {
			for _, f := range fields {
				keepFields[f] = true
			}
		}

		result := make(map[string]interface{})
		formattedStruct := jsonToMap(data) // Convert struct to map using json tags

		for k, v := range formattedStruct {
			if keepFields[k] {
				result[k] = v
			}
		}
		return result
	}

	// Handle map
	if val.Kind() == reflect.Map {
		keepFields := make(map[string]bool)
		if len(fields) == 0 {
			keepFields["id"] = true
			keepFields["name"] = true
			keepFields["stem"] = true
		} else {
			for _, f := range fields {
				keepFields[f] = true
			}
		}

		result := make(map[string]interface{})
		iter := val.MapRange()
		for iter.Next() {
			k := iter.Key().String()
			if keepFields[k] {
				result[k] = iter.Value().Interface()
			}
		}
		return result

	}

	return data
}

// Helper to check if a struct has a field by name (case insensitive for simplicity in logic, but reflect is case sensitive)
func hasField(v reflect.Value, name string) bool {
	return v.FieldByName(name).IsValid()
}

// Helper to convert struct to map[string]interface{} using json tags
func jsonToMap(v interface{}) map[string]interface{} {
	b, _ := json.Marshal(v)
	var m map[string]interface{}
	json.Unmarshal(b, &m)
	return m
}

// parseFields parses the "fields" argument from input args
func parseFields(args map[string]interface{}) []string {
	if v, ok := args["fields"]; ok {
		if arr, ok := v.([]interface{}); ok {
			var fields []string
			for _, item := range arr {
				if s, ok := item.(string); ok {
					fields = append(fields, s)
				}
			}
			return fields
		}
	}
	return nil
}

// Available field definitions for documentation
const (
	FieldDescOrganisation  = "Available fields: id, name"
	FieldDescQualification = "Available fields: id, name, organisationId"
	FieldDescSyllabus      = "Available fields: id, name, code, qualificationId"
	FieldDescChapter       = "Available fields: id, name, parentId, syllabusId" // removed children as we don't return trees
	FieldDescPaperSeries   = "Available fields: id, name, syllabusId"
	FieldDescPaperCode     = "Available fields: id, name, syllabusId"
	FieldDescPastPaper     = "Available fields: id, name, year, syllabusId, paperCodeId, paperSeriesId, questionNumber"
	FieldDescQuestion      = "Available fields: id, stem, difficult, status, syllabusId, pastPaperId, questionContents, chapters"
)
