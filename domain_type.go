package XoRPC

func Route_Type(r interface{}) bool {
	//	if r interface type of deal func.
	switch r.(type) {
	case FindRoute:
		if findRoute, ok := r.(FindRoute); ok {
			return route_findRoute(findRoute)
		}
	case ReturnRoute:
		if returnRoute, ok := r.(ReturnRoute); ok {
			return route_null(returnRoute)
		}
	}
	return false
}

// pass RETURNNODE
func route_null(r ReturnRoute) bool {
	if r.Type == "FINDNODERTURN" {
		if addres := IsIPV64(r.Addres); addres == 6 || addres == 4 && r.Port < 65535 && r.NULL == true {
			return true
		}
	}
	return false
}

// pass Find Route TO.
func route_findRoute(r FindRoute) bool {
	if r.Type == "FINDNODE" {
		return true
	}
	return false
}
