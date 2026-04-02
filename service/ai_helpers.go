package service

import "strings"

// extractJSONFromAIResponse extracts a JSON string from an AI response that may be
// wrapped in markdown code fences or prefixed/suffixed with extra explanation text.
// It handles these common AI output patterns in order:
//  1. ```json ... ``` fenced code blocks
//  2. ``` ... ``` fenced code blocks
//  3. Responses where JSON starts after leading prose (by finding the first '{' or '[')
func extractJSONFromAIResponse(response string) string {
	response = strings.TrimSpace(response)
	if response == "" {
		return response
	}

	// Strip markdown code fences: ```json ... ``` or ``` ... ```
	if idx := strings.Index(response, "```json"); idx != -1 {
		start := idx + len("```json")
		if end := strings.Index(response[start:], "```"); end != -1 {
			return strings.TrimSpace(response[start : start+end])
		}
	}
	if idx := strings.Index(response, "```"); idx != -1 {
		start := idx + len("```")
		if end := strings.Index(response[start:], "```"); end != -1 {
			return strings.TrimSpace(response[start : start+end])
		}
	}

	// Find the first JSON object or array boundary and return from there.
	// The AI is instructed to output only JSON so trailing prose after the JSON
	// payload is not expected; go's json.Unmarshal ignores trailing whitespace
	// and the decoder stops at the end of the outermost value.
	objStart := strings.Index(response, "{")
	arrStart := strings.Index(response, "[")

	start := -1
	switch {
	case objStart != -1 && arrStart != -1:
		if objStart < arrStart {
			start = objStart
		} else {
			start = arrStart
		}
	case objStart != -1:
		start = objStart
	case arrStart != -1:
		start = arrStart
	}

	if start != -1 {
		return response[start:]
	}
	return response
}
