package main

import "fmt"

type Elevator struct {
	motionStatus int
	totalFloor int
	pressFloor map[int]bool
	currentFloor int
}
const STATUS_PAUSE = 0; //暂停状态
const STATUS_UP = 1;    //上行状态
const STATUS_DOWN = 2;  //下行状态
const STATUS_STOP = -1; //停止运行

func NewElevator(t int) (elevator *Elevator) {
	pref := map[int]bool{}
	for i := 1; i <= t; i++ {
		pref[i] = false
	}
	return &Elevator{
		motionStatus: STATUS_PAUSE,
		totalFloor:   t,
		pressFloor:   pref,
		currentFloor: 1,
	}
}
func (e *Elevator) PushButton(pressFloor int) (floorYorN bool) {
	_, ok := e.pressFloor[pressFloor]
	if !ok {
		return false
	}
	e.pressFloor[pressFloor] = true
	return true
}
func (e *Elevator) gotoNextFloor() (aimFloor int) {
	aimFloor = -1
	if e.motionStatus == STATUS_PAUSE {
		//todo
		for k, v := range e.pressFloor {
			if v {
				aimFloor = k
				break
			}
		}
	} else if e.motionStatus == STATUS_UP {
		//todo
		for floor := e.currentFloor + 1; floor <= e.totalFloor; floor++ {
			if e.pressFloor[floor] {
				aimFloor = floor
				break
			}
		}
		if aimFloor == -1 && e.currentFloor > 1 {
			for floor := e.currentFloor - 1; floor >= 1; floor-- {
				if e.pressFloor[floor] {
					aimFloor = floor
					break
				}
			}
		}
	} else if e.motionStatus == STATUS_DOWN {
		//todo
		for floor := e.currentFloor - 1; floor >= 1; floor-- {
			if e.pressFloor[floor] {
				aimFloor = floor
				break
			}
		}
		if aimFloor == -1 && e.currentFloor < e.totalFloor {
			for floor := e.currentFloor + 1; floor <= e.totalFloor; floor++ {
				if e.pressFloor[floor] {
					aimFloor = floor
					break
				}
			}
		}
	}
	if aimFloor == -1 {
		aimFloor = e.currentFloor
		e.motionStatus = STATUS_PAUSE
	}
	if aimFloor == e.currentFloor {
		e.motionStatus = STATUS_PAUSE
	} else if aimFloor > e.currentFloor {
		if aimFloor < e.totalFloor {
			e.motionStatus = STATUS_UP
		} else {
			e.motionStatus = STATUS_DOWN
		}
	} else if aimFloor < e.currentFloor {
		if aimFloor > 1 {
			e.motionStatus = STATUS_DOWN
		} else {
			e.motionStatus = STATUS_UP
		}
	}
	e.currentFloor=aimFloor
	e.pressFloor[e.currentFloor]=false
	return aimFloor
}

func main() {
	ele:=NewElevator(5)
	ele.PushButton(2)
	airFloor:=ele.gotoNextFloor()
	fmt.Println(airFloor,ele.currentFloor)
	ele.PushButton(1)
	ele.PushButton(5)
	ele.PushButton(4)
	airFloor1:=ele.gotoNextFloor()
	fmt.Println("第一次要去的",airFloor1)
	airFloor2:=ele.gotoNextFloor()
	fmt.Println("第二次要去的",airFloor2)
	airFloor3:=ele.gotoNextFloor()
	fmt.Println("第三次要去的",airFloor3)



}
