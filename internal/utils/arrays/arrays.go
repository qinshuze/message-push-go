package arrays

import (
	"strconv"
	"strings"
)

func Int32ToString(v []int32) string {
	var str = strings.Builder{}
	for i, i3 := range v {
		if i != 0 {
			str.WriteString(",")
		}
		str.WriteString(strconv.FormatInt(int64(i3), 10))
	}

	return str.String()
}

func Int16ToString(v []int16) string {
	var str = strings.Builder{}
	for i, i3 := range v {
		if i != 0 {
			str.WriteString(",")
		}
		str.WriteString(strconv.FormatInt(int64(i3), 10))
	}

	return str.String()
}

func Int8ToString(v []int8) string {
	var str = strings.Builder{}
	for i, i3 := range v {
		if i != 0 {
			str.WriteString(",")
		}
		str.WriteString(strconv.FormatInt(int64(i3), 10))
	}

	return str.String()
}

func StrToString(v []string) string {
	var str = strings.Builder{}
	for i, i3 := range v {
		if i != 0 {
			str.WriteString(",")
		}
		str.WriteString(i3)
	}

	return str.String()
}

func ByteToString(v []byte) string {
	var str = strings.Builder{}
	for i, i3 := range v {
		if i != 0 {
			str.WriteString(",")
		}
		str.WriteString(string(i3))
	}

	return str.String()
}

func BoolToString(v []byte) string {
	var str = strings.Builder{}
	for i, i3 := range v {
		if i != 0 {
			str.WriteString(",")
		}
		str.WriteString(string(i3))
	}

	return str.String()
}

func ToString[T any](v []T) string {
	var str = strings.Builder{}
	for i, t := range v {
		if i != 0 {
			str.WriteString(",")
		}

		if any(t) == nil {
			str.WriteString("nil")
			continue
		}

		switch v1 := any(t).(type) {
		case string:
			str.WriteString(v1)
			continue
		case int:
			str.WriteString(strconv.FormatInt(int64(v1), 10))
			continue
		case int8:
			str.WriteString(strconv.FormatInt(int64(v1), 10))
			continue
		case int16:
			str.WriteString(strconv.FormatInt(int64(v1), 10))
			continue
		case int32:
			str.WriteString(strconv.FormatInt(int64(v1), 10))
			continue
		case int64:
			str.WriteString(strconv.FormatInt(v1, 10))
			continue
		case byte:
			str.WriteString(string(v1))
			continue
		case float32:
			str.WriteString(strconv.FormatFloat(float64(v1), 'f', -1, 64))
			continue
		case float64:
			str.WriteString(strconv.FormatFloat(v1, 'f', -1, 64))
			continue
		case bool:
			str.WriteString(strconv.FormatBool(v1))
			continue
		}
	}

	return str.String()
}

func Int64ToString(v []int64) string {
	var str = strings.Builder{}
	for i, i3 := range v {
		if i != 0 {
			str.WriteString(",")
		}
		str.WriteString(strconv.FormatInt(i3, 10))
	}

	return str.String()
}
