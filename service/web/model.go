package web

import "time"

//Vdoc model
type Vdoc struct {
	URL        string
	SubmitTime time.Time
	Supplier   string
}

//Doc document
type Doc struct {
	URL string
	Org string
}
