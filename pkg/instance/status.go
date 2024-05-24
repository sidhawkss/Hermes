package instance


type CmpStat struct{
	Hostname string
	Ip	 string
	Country string
	Username string
	Os string
}

type Machine struct{
	Machines []CmpStat
}

func ShowInfo() CmpStat{
	computer := CmpStat {
		Hostname: "VANILLA",
		Ip: "169.254.169.254",
		Country: "N/A",
		Username: "SpongeBOB",
		Os: "FreeBSD",
	}

	return computer
}
