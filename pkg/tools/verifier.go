package tools

import (
	"regexp"
)

func VerifyAddr(destination string) bool{
	re := regexp.MustCompile(`[0-9]{1,3}.[0-9]{1,3}.[0-9]{1,3}.[0-9]{1,3}`); // can be simplified
	if(re.Find([]byte(destination)) != nil){
		return true;
	} else {
		return false;
	}
}