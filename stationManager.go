package main

import "sync"

type stationManager struct{
	isPlatformFree bool
	lock  *sync.Mutex
	trainQueue []train
}

func newStationManager() *stationManager {
	newSm := &stationManager{
		isPlatformFree : true,
		lock : &sync.Mutex{},
	}

	return newSm

}



func (sm *stationManager) canLand(t train) bool{
	sm.lock.Lock()
	defer sm.lock.Unlock()

	if sm.isPlatformFree{
		sm.isPlatformFree = false
		return true
	}
	sm.trainQueue = append(sm.trainQueue, t)
	return false
}


func (sm *stationManager) notifyFree() {
	sm.lock.Lock()
	defer sm.lock.Unlock()
	if !sm.isPlatformFree{
		sm.isPlatformFree = true
	}
	if len(sm.trainQueue) > 0{
		firstTrain := sm.trainQueue[0]
		sm.trainQueue = sm.trainQueue[1:]
		firstTrain.permitArrival()
	}
}