package conn

import (
	"fmt"
	"net"
	"context"
	//"Hermes/pkg/enc"
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

var dnsServer string = "127.0.0.1:53";
var Resolv = &net.Resolver{
	PreferGo: true,
	Dial: func(ctx context.Context, network, address string) (net.Conn, error){
		var d net.Dialer;
		return d.DialContext(ctx, "udp", dnsServer);
	},
	
};

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
		ips, err := Resolv.LookupIPAddr(context.Background(),chunk+".localhost");
		if err != nil {
			fmt.Println("Resolution error");
		}
		fmt.Println(ips);
	}
	//handle response
	
	return "a"
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