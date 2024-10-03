package filter

import (
	"sync"
	"time"
)

const (
	ConnNumber = 65500  // port max 65536 It means the maximum number of people
)

type (
	FilterNumber struct {
		number map[int]string
	}
)

func AddConnAddres(flr *FilterNumber) {
	flr.number 
}