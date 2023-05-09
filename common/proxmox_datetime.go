package common

import (
	"strconv"
	"strings"
	"time"
)

type ProxmoxDateTimeValue struct {
	Default  uint64
	selected uint64
}

func (e *ProxmoxDateTimeValue) Set(value string) error {
	if strings.Contains(value, "T") {
		t, err := time.Parse("2006-01-02T15:04:05Z", value)
		if err != nil {
			return err
		}
		e.selected = uint64(t.Unix())
	} else {
		t, err := strconv.Atoi(value)
		if err != nil {
			return err
		}
		e.selected = uint64(t)
	}
	return nil
}

func (e *ProxmoxDateTimeValue) String() string {
	return strconv.FormatUint(e.selected, 10)
}

func (e *ProxmoxDateTimeValue) Value() uint64 {
	return e.selected
}
