package pkg

var compareStrategies = map[CompareType]CompareStrategy{
	CompareTypeEq:      CompareTypeEqFunc,
	CompareTypeEqCount: CompareTypeEqCountFunc,
}

func CompareTypeEqFunc(a, b any) (bool, error) {
	switch typedA := a.(type) {
	case []any:
		typedB, ok := b.([]any)
		if !ok {
			return false, InvalidValueType
		}
		for idxA, valueA := range typedA {
			if valueA != typedB[idxA] {
				return false, nil
			}
		}
		return true, nil
	case map[string]any:
		typedB, ok := b.(map[string]any)
		if !ok {
			return false, InvalidValueType
		}
		for keyA, valueA := range typedA {
			if valueA != typedB[keyA] {
				return false, nil
			}
		}
		return true, nil
	default:
		return a == b, nil
	}
}

func CompareTypeEqCountFunc(a, b any) (bool, error) {
	switch typedA := a.(type) {
	case []any:
		typedB, ok := b.([]any)
		if !ok {
			return false, InvalidValueType
		}
		return len(typedA) == len(typedB), nil
	case map[string]any:
		typedB, ok := b.(map[string]any)
		if !ok {
			return false, InvalidValueType
		}
		return len(typedA) == len(typedB), nil
	default:
		return false, InvalidValueType
	}
}
