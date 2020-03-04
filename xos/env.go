package xos

import (
	"os"
	"strconv"
	"strings"
)

func EnvBool(key string) (bool, error) {
	if val, ok := os.LookupEnv(key); ok {
		return strconv.ParseBool(val)
	}
	return false, nil
}

func EnvInt64(key string) (int64, error) {
	if val, ok := os.LookupEnv(key); ok {
		return strconv.ParseInt(val, 10, 64)
	}
	return 0, nil
}

func EnvInt64s(key string) ([]int64, error) {
	if val, ok := os.LookupEnv(key); ok {
		items := strings.Split(val, ",")
		ints := make([]int64, 0, len(items))
		for _, item := range items {
			i, err := strconv.ParseInt(item, 10, 64)
			if err != nil {
				return nil, err
			}
			ints = append(ints, i)
		}
		return ints, nil
	}
	return nil, nil
}

func EnvString(key string) (string, error) {
	return os.Getenv(key), nil
}

func EnvStrings(key string) ([]string, error) {
	return strings.Split(os.Getenv(key), ","), nil
}
