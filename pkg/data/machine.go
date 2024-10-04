package data

type Status struct{
	Id string
	Hostname string
	Ip	 string
	Country string
	Username string
	Os string
}

type Machine struct{
	Status 
}
