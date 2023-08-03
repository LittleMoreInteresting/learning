package main

import (
	"fmt"

	"learning/aes_demo/aes_ebc"
)

func main() {

	aes_old := aes_ebc.NewAesEBC(aes_ebc.WithKey("P9PI#l8jWH@fzZtm"))
	aes_new := aes_ebc.NewAesEBC(aes_ebc.WithKey("x7MykcczUKmvcUXB"))
	phones := []string{
		"2C7nEsyRBkAhvBnnQpSTeg==",
		"rt6VbT16bEOYH9yE55qeGw==",
		"JTIAHJrOBvH6WEaTpIu9Ew==",
		"b48AQeiZrOEvvZynmf8uoQ==",
	}
	for _, phone := range phones {
		decrypt := aes_old.Decrypt(phone)

		fmt.Println(phone, decrypt, aes_new.EncryptPhone(decrypt))
	}
	fmt.Println(aes_new.DecryptPhone("aenETOKlMNkJvwQcg/qtKA=="))
}
