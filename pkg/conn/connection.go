package conn

import (
	"fmt"
)

type Interact interface{
	Dns() string
	Http() string
	Websocket() string
	getType() string
}

type Data struct{
	SecureString string
	ConnectionType string
}

func SendDataTarget(i Interact) string{
	switch(i.getType()){
	case "0":
		return i.Dns();
	case "1":
		return i.Http();
	case "2":
		return i.Websocket();
	default:
		return "NOT IMPLEMENTED"
	}
}

func (d Data) Dns() string{
	
	for i:=16; i <= len(d.SecureString); i += 16{
		chunk := d.SecureString[i-16:i];
		fmt.Println(chunk);
		//net.LookupIP(sstr+".dns");
	}
	// send done message
	// wait response, 
	// return response,
	return "DNS SENT"
}

func (d Data) Http() string{
	return "HTTP"
}

func (d Data) Websocket() string{
	return "WEBSOCKET"
}

func (d Data) getType() string{
	return d.ConnectionType
}