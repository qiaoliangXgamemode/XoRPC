package XoRPC

type reylogin struct {
	ServiceID         int
	ServiceName       string
	Serviceweight     int
	ServiceEncrypt    bool
	ServiceFilter     bool
	ServiceFiltertype string
	// Public token vry spDimain network
	PublicToken string
	// Private token vry WideArea network
	PrivateToken string
}

type recvlogin struct {
}

func wrtieLoologin(*reylogin) {

}

// func rpcPareslogin(reftext *loologin) {
// 	return reftext
// }
