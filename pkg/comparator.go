package pkg

import (
	"strconv"
	"strings"
	"sync"
)

// Compare: based on a config, this functions allows compare
// json specific json attributes. It`s supports arrays and multi level nested objects.
func Compare(expect, actual any, config CompareConfig) bool {
	var wg sync.WaitGroup

	var expectValues map[string]any
	var expectErr error

	var actualValues map[string]any
	var actualErr error

	wg.Add(2)
	go func() {
		defer wg.Done()
		expectValues, expectErr = getValuesToCompare(expect, "", config.Path)
	}()

	go func() {
		defer wg.Done()
		actualValues, actualErr = getValuesToCompare(actual, "", config.Path)
	}()

	wg.Wait()

	if expectErr != nil || actualErr != nil {
		return false
	}

	for i, expectValue := range expectValues {
		compareStrategy := compareStrategies[config.CompareType]
		if compareStrategy == nil {
			return false
		}
		compareResult, err := compareStrategy(expectValue, actualValues[i])
		if err != nil {
			return false
		}
		return compareResult
	}
	return true
}

func getValuesToCompare(input any, previousPath, path string) (map[string]any, error) {
	if input == nil {
		if path != "" {
			return map[string]any{}, PathNotFoundErr
		}
		return map[string]any{}, nil
	}

	pathValues := map[string]any{}
	var err error

	pathParts := strings.Split(path, ".")
	currentPathPart := pathParts[0]

	currentPartValue, err := getPartValue(input, currentPathPart)
	if err != nil {
		return map[string]any{}, err
	}

	if len(pathParts) == 1 {
		return parseResult(currentPartValue, previousPath, pathParts[0]), nil
	}

	switch v := currentPartValue.(type) {
	case []any:
		pathValues, err = processArrayPathPart(v, previousPath, currentPathPart, pathParts)
		break
	case map[string]any:
		newPreviousPath := previousPath + "." + currentPathPart
		if previousPath == "" {
			newPreviousPath = currentPathPart
		}
		pathValues, err = getValuesToCompare(v, newPreviousPath, strings.Join(pathParts[1:], "."))
		break
	default:
		pathValues, err = map[string]any{}, PathNotFoundErr
	}

	if err != nil {
		return map[string]any{}, err
	}

	return pathValues, nil
}

func processArrayPathPart(arrayValue []any, previousPath string, currentPathPart string, pathParts []string) (map[string]any, error) {
	pathValues := map[string]any{}
	var err error

	desiredIndex, parseErr := strconv.ParseInt(pathParts[1], 10, 64)
	wantSpecificArrayIdx := parseErr == nil && desiredIndex >= 0
	if wantSpecificArrayIdx {
		newPreviousPath := previousPath + "." + currentPathPart
		if previousPath == "" {
			newPreviousPath = currentPathPart
		}
		return getValuesToCompare(arrayValue, newPreviousPath, strings.Join(pathParts[1:], "."))
	}

	for i, part := range arrayValue {
		newPreviousPath := previousPath + "." + strconv.FormatInt(int64(i), 10)
		if previousPath == "" {
			newPreviousPath = strconv.FormatInt(int64(i), 10)
		}
		var arrayPathValues map[string]any
		arrayPathValues, err = getValuesToCompare(part, newPreviousPath, strings.Join(pathParts[1:], "."))
		if err != nil {
			return pathValues, err
		}
		for key, value := range arrayPathValues {
			pathValues[key] = value
		}
	}

	return pathValues, err
}

func getPartValue(jsonValue any, partPath string) (any, error) {
	switch v := jsonValue.(type) {
	case []any:
		idx, err := strconv.ParseInt(partPath, 10, 64)
		if err != nil {
			return nil, PathNotFoundErr
		}

		if idx >= int64(len(v)) {
			return nil, PathNotFoundErr
		}

		return v[idx], nil
	case map[string]any:
		return v[partPath], nil
	}
	return nil, PathNotFoundErr
}

func parseResult(nextPartValue any, previousPath string, partPath string) map[string]any {
	result := map[string]any{}

	switch v := nextPartValue.(type) {
	case []any:
	case map[string]any:
		for key, value := range v {
			valuePath := previousPath + "." + key
			if previousPath == "" {
				valuePath = key
			}
			result[valuePath] = value
		}
		return result
	default:
		valuePath := previousPath + "." + partPath
		if previousPath == "" {
			valuePath = partPath
		}
		result[valuePath] = nextPartValue
		return result
	}

	return result
}
