package operations

import (
	"fmt"
	"strconv"
	"io/ioutil"
	"encoding/json"
	"Hermes/pkg/data"
)

func OpenJson() []data.Machine{
	var machines []data.Machine;
	jData, err := ioutil.ReadFile("utils/data.json");
	if err != nil{
		fmt.Println("Error: openning json");
	}
	
	err = json.Unmarshal([]byte(jData), &machines);
	if err != nil{
		fmt.Println("Error: Unmarshal data");
	}
	
	return machines;
}

func ReadData() []data.Machine{
	var machines []data.Machine = OpenJson();
	
	return machines;
}

func WriteData(hostname string, ip string, country string, username string, os string) []data.Machine{
	var machines []data.Machine = OpenJson();
	// implement automatic id

	var d []data.Machine = ReadData();

	newObject := &data.Machine{ Status: data.Status{
		Id: strconv.Itoa(len(d)+1),
		Hostname: hostname, 
		Ip: ip, 
		Country: country, 
		Username: username,
		Os: os, 
	}};

	machines = append(machines, *newObject);
	jData, err := json.Marshal(machines);
	if err != nil{
		fmt.Println("Error: Marshal data");
	}

	err = ioutil.WriteFile("utils/data.json", jData, 0640);

	return machines;
}