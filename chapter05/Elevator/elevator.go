package main

//const STATUS_PAUSE = 0; //暂停状态
//const STATUS_UP = 1;    //上行状态
//const STATUS_DOWN = 2;  //下行状态

//type Elevator struct {
//	motionStatus int
//	totalFloor int
//	pressFloor map[int]bool
//	currentFloor int
//}

func (e *Elevator) pressElevator(f int) {
	e.pressFloor[f] = true
}


