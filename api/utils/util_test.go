package utils

import (
	"fmt"
	"testing"
)

func TestNewUUID(t *testing.T) {
	str:="hello"
	fmt.Println(NewUUID(str))
}
