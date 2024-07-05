package conn

import (
	
)

func SendDataTarget(sec_str string, connection_type string) string{
	switch(connection_type){
	case "0":
		return "DNS"
	case "1":
		return "HTTP"
	case "2":
		return "WEBSOCKET"
	default:
		return "NOT IMPLEMENTED"
	}
	return ""
}
