package main

import (
	"log"
	"math"
	"strconv"
	"strings"
)

func main() {
	hashUpdate()
}

func hashUpdate() {
	remoteAddrs := "192.168.1.1"
	segments := strings.Split(remoteAddrs, ".")
	l := len(segments)
	var inhash float64
	for i := range l {
		num, _ := strconv.Atoi(segments[i])
		// log.Printf("%d | %d hash %f", num, l-(i+1), float64(num)*math.Pow(256, float64(l-1-i)))
		// log.Printf("%d %f  %d", num, float64(num)*math.Pow(256, float64(l-i-1)), l-1-i)
		inhash += float64(num) * math.Pow(256, float64(l-1-i))
	}
	log.Printf("hash key %d", int(inhash))
	inhash = math.Mod(inhash, math.Pow(2, 32))
	// inhash = hash % (2 ^ 32)
	log.Printf("hash key %d", int(inhash))
}
