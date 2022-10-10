package test

import (
	"cloud-disk/core/helper"
	"fmt"
	"testing"
	"time"
)

func TestRandCode(t *testing.T) {
	fmt.Println(time.Now().UnixNano())
	fmt.Println(helper.RandCode())
}
