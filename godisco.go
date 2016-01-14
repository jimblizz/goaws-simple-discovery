package godisco

type Godisco struct {}

func (gd *Godisco) Get(service string) (ips []string, err error) {
	ips = append(ips, "10.10.10.9")
	return
}

func (gd *Godisco) GetFirst(service string) (ip string, err error) {
	ips, ipsErr := gd.Get(service)

	if len(ips) > 0 && ipsErr != nil {
		ip = ips[0]
	} else {
		err = ipsErr
	}
	return
}