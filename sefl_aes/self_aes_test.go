package sefl_aes

import (
	"fmt"
	"testing"
)

func TestNewAes(t *testing.T) {
	opt := WithKey("1111")
	res := NewAes(opt, WithHeader(2))
	fmt.Println(res.key)
	fmt.Println(res.head)
}
