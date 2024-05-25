package operations

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
	"HermesC2/pkg/data"
)

func ReadData() []data.Machine{
	var machines []data.Machine;
	jData, err := ioutil.ReadFile("utils/data.json");
	if err != nil{
		fmt.Println("Error: openning json");
	}
	
	err = json.Unmarshal([]byte(jData), &machines);
	if err != nil{
		fmt.Println("Error: Unmarshal json");
	}
	
	return machines;
}