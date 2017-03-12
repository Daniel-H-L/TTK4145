package Network

type Backup struct {
	Orders    [3][4]int
	Direction int
	Floor     int
	State     int
}

//Key for Backup in StandardData is IP.

type StandardData struct {
	Type         int
	IP           string
	IsAlive      string
	Backup       map[string]*Backup
	DescendantNr int
	SetLights    [3][4]int
}
