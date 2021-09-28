package status

type Status struct {
	Status struct {
		Water int
		Wind  int
	}
}

func (s Status) WaterStatus() string {
	status := "aman"
	if s.Status.Water >= 6 && s.Status.Water <= 8 {
		status = "siaga"
	} else if s.Status.Water > 8 {
		status = "bahaya"
	}
	return status
}

func (s Status) WindStatus() string {
	status := "aman"
	if s.Status.Wind >= 7 && s.Status.Wind <= 15 {
		status = "siaga"
	} else if s.Status.Wind > 15 {
		status = "bahaya"
	}
	return status
}
