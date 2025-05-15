package conn

import (
	"fmt"
	"net"
	"errors"
	"strings"
	"net/http"
	"io/ioutil"
	"Hermes/pkg/enc"
	"Hermes/pkg/data"
	"Hermes/pkg/tools"
	"Hermes/pkg/events"
)

type Interact interface{
	Dns() string
	Http(machineID string) string
	Websocket() string
	getType() string
}

type Data struct{
	CommandString string
	ConnectionType string
}

var dnsServer string = "127.0.0.1:53";
var Resolv = &net.Resolver{
	PreferGo: true,
	
}

func SendDataTarget(i Interact, machineID string) string{
	switch(i.getType()){
	case "0":
		return i.Dns();
	case "1":
		return i.Http(machineID);
	case "2":
		return i.Websocket();
	default:
		return "NOT IMPLEMENTED";
	}
}

func (d Data) Dns() string{
	return "a";
}

func (d Data) Http(machineID string) string{
	var currentMachine data.Machine = tools.SelectById(machineID);

	if(tools.VerifyAddr(currentMachine.Ip) == false){
		events.ConnectionEventCustom(errors.New(""), "Invalid Ip");
		return "Invalid Ip";
	} else {
		var URL string = fmt.Sprintf("http://%s:4000/",currentMachine.Ip) // format check
		var command *strings.Reader = strings.NewReader(enc.Encrypt(d.CommandString));

		res, err := http.Post(URL, "text/html", command);
		events.ConnectionEventCustom(err, "Invalid request");
		if( err == nil){
			body, err := ioutil.ReadAll(res.Body);
			events.ConnectionEventCustom(err, "No response");
			return string(body);
		}
	}
	return "no connection"
}

func (d Data) Websocket() string{
	return "WEBSOCKET";
}

func (d Data) getType() string{
	return d.ConnectionType;
}
