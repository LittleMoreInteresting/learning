package aes_demo

import (
	"fmt"
	"testing"
)

func TestAes_Encrypt(t *testing.T) {
	//I1q1D3G5xI43v6wFGKa5/Q==
	//13164238899
	aes := NewAes(WithKey("y27bulYuw6cmm@ln"))
	code := aes.Encrypt("13164238899")
	fmt.Println(code)
}
