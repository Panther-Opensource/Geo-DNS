package structs

type Data struct {
	Pops map[string]string `json:"pops"`
}

type Pop struct {
	Ip       string
	Hostname string
}
