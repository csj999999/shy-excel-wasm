package main

func main() {

	c := make(chan struct{})
	regFuncs()
	<-c

}

func regFuncs() {

}
