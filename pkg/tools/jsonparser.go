package tools

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
	var machines []data.Machine = ReadData(); // implement automatic ids
	newObject := &data.Machine{Status: data.Status{
		Id: strconv.Itoa(len(machines)+1),
		Hostname: hostname, 
		Ip: ip, 
		Country: country, 
		Username: username,
		Os: os, 
	}};

	jData, err := json.Marshal(append(machines, *newObject));
	if err != nil{
		fmt.Println("Error: Marshal data");
	}
	
	err = ioutil.WriteFile("utils/data.json", jData, 0640);
	return machines;
}

func SelectById(machineID string) data.Machine{
	var machines []data.Machine = ReadData();
	var currentMachine data.Machine;
	for i := 0; i < len(machines); i++ {
		if machineID == machines[i].Id {
			currentMachine = machines[i];
		}
	}

	return currentMachine;
}