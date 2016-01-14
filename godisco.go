package godisco

type Godisco struct {}

func (gd *Godisco) Get(service string) (ips []string, err error) {

	ips = append(ips, "10.10.10.9")

	return
}