// Code generated from Pkl module `introduction.of.pkl`. DO NOT EDIT.
package loglevel

import (
	"encoding"
	"fmt"
)

type LogLevel string

const (
	DEBUG LogLevel = "DEBUG"
	INFO  LogLevel = "INFO"
	WARN  LogLevel = "WARN"
	ERROR LogLevel = "ERROR"
)

// String returns the string representation of LogLevel
func (rcv LogLevel) String() string {
	return string(rcv)
}

var _ encoding.BinaryUnmarshaler = new(LogLevel)

// UnmarshalBinary implements encoding.BinaryUnmarshaler for LogLevel.
func (rcv *LogLevel) UnmarshalBinary(data []byte) error {
	switch str := string(data); str {
	case "DEBUG":
		*rcv = DEBUG
	case "INFO":
		*rcv = INFO
	case "WARN":
		*rcv = WARN
	case "ERROR":
		*rcv = ERROR
	default:
		return fmt.Errorf(`illegal: "%s" is not a valid LogLevel`, str)
	}
	return nil
}
