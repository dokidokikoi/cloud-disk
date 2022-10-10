package test

import (
	"cloud-disk/core/helper"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestXormTest(t *testing.T) {
	fmt.Println(helper.Md5("123123"))
}
