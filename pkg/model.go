package pkg

import "errors"

var (
	PathNotFoundErr  = errors.New("path not found")
	InvalidValueType = errors.New("invalid value type")
)

type CompareType int

const (
	CompareTypeEq CompareType = iota
	CompareTypeEqCount
)

type CompareConfig struct {
	Path        string      `json:"expect_path"`
	CompareType CompareType `json:"compare_type"`
}

type CompareStrategy func(a, b any) (bool, error)
