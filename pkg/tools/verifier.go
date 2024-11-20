package tools

import (
	"fmt"
	"regexp"
)

func VerifyAddr (destination string) bool{
	re := regexp.MustCompile(`(([a-z]|[A-Z]|[0-9])+\.([a-z]|[A-Z])+)|([0-9]{0,3}\.){0,3}?([0-9]{0,3})`); // can be simplified
	if(re.Find([]byte(destination)) != nil){
		return true;
	}else{
		fmt.Println("[TOOLS][VERIFYDESTINATION][ERROR] Invalid destination.");
		return false;
	}
}