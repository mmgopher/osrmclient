package model

type  Route struct {
	Destination string `json:"destination"`
	Duration    float64 `json:"duration"`
	Distance    float64 `json:"distance"`
}


type  Response struct {
	Source string `json:"source"`
	Routes []Route `json:"routes"`
}


type ByTime []Route

func (r ByTime) Len() int           { return len(r) }
func (r ByTime) Less(i, j int) bool {
	if (r[i].Duration == r[j].Duration) {
		return r[i].Distance < r[j].Distance
	} else {
		return r[i].Duration < r[j].Duration
	}
}
func (r ByTime) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }


