package events

import (
	"fmt"
)

func TemplateEvent(err error){
	if err != nil {
		fmt.Println("[MAIN][TEMPLATE][ERROR] Template parsing.");
	}
}

func TemplateEventCustom(err error, message string){
	if err != nil {
		fmt.Printf("[MAIN][TEMPLATE][ERROR] %s.\n",message);
	}
}

func ConnectionEventCustom(err error, message string){
	if err != nil {
		fmt.Printf("[CONNECTION][HTTP][ERROR] %s.\n", message);
	}
}