package cpumeter

import (
	"os"
	"strings"
)

const (
	PROC_1        string = "/proc/1/attr/current"
	NOT_CONTAINER string = "unconfined"
)

func IsContainer() (bool, error) {
	data, err := os.ReadFile(PROC_1)
	if err != nil {
		return false, err
	}
	return !strings.HasPrefix(string(data), NOT_CONTAINER), nil
}
