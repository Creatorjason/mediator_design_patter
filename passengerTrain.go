package main


import "fmt"


type passengerTrain struct{
	mediator mediator
}

func (ps *passengerTrain) requestArrival(){
	if ps.mediator.canLand(ps){
		fmt.Println("Passenger Train: Landing...")
	}else{
		fmt.Println("Passenger Train: Waiting...")
	}
}

func (ps *passengerTrain) departure(){
	fmt.Println("Passenger  Train: Leaving...")
	ps.mediator.notifyFree()
}

func (ps *passengerTrain) permitArrival(){
	fmt.Println("Passenger Train: Arrival Permitted. Landing")
}