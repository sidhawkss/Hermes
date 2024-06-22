package build

import (
	"fmt"
	"Hermes/pkg/enc"
)

func dnsInteraction(input string){
	var encrypted_text string = enc.Encrypt(input);
	for i:=16; i <= len(encrypted_text); i += 16{
		fmt.Println(encrypted_text[i-16:i]);
	}
}