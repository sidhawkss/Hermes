package conn

import (
	"fmt"
	"net"
	"strings"
	"context"
	"net/http"
	"io/ioutil"
	//"Hermes/pkg/enc"
)

type Interact interface{
	Dns() string
	Http(machineID string) string
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

func SendDataTarget(i Interact, machineID string) string{
	switch(i.getType()){
	case "0":
		return i.Dns();
	case "1":
		return i.Http(machineID);
	case "2":
		return i.Websocket();
	default:
		return "NOT IMPLEMENTED"
	}
}

func (d Data) Dns() string{
	for i:=16; i <= len(d.SecureString); i += 16{
		chunk := d.SecureString[i-16:i];
		ips, err := Resolv.LookupIPAddr(context.Background(),chunk+".localhost"); // format check
		if err != nil {
			fmt.Println("[CONNECTION][DNS][ERROR] resolution error");
		}
		fmt.Println(ips);
	}
	//handle responses
	return "a"
}

/* The connection needs to be closed, so the user can't control the destination URL to avoid SSRF */
/* create an interface to handle the machine, to make a ID to the machine, and perform actions based on this ID, and only will work with the ID */
func (d Data) Http(machineID string) string{
	var URL string = fmt.Sprintf("http://%s:5000/",machineID) // format check
	var command *strings.Reader = strings.NewReader(d.SecureString);
	res, err := http.Post(URL, "text/html", command);
	if err != nil {
        fmt.Println("[CONNECTION][HTTP][ERROR] invalid request");
    }

	body, err := ioutil.ReadAll(res.Body);
	fmt.Println(body);
	if err != nil {
		fmt.Println("[CONNECTION][HTTP][ERROR] no response");
	}

	return string(body);
}

func (d Data) Websocket() string{
	return "WEBSOCKET"
}

func (d Data) getType() string{
	return d.ConnectionType
}
