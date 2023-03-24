package qesygo

import (
	"errors"
	"sync"
	"time"
)

// 0(1位，且始终为0)|时间戳(41位)|工作机器id(10位)|序列号(12位)
type SNOW struct {
	machineID     int64
	snow          int64
	lastTimeStamp int64
	lock          sync.Mutex
}

func CreatId() int64 { //获取唯一ID号
	snow, _ := newSnow(654)
	return snow.getID()
}

func (snow *SNOW) initTimeStamp() {
	snow.lastTimeStamp = time.Now().UnixNano() / 1e6
}

func newSnow(machineID int64) (*SNOW, error) {
	snow := &SNOW{
		machineID:     0,
		snow:          0,
		lastTimeStamp: time.Now().UnixNano() / 1e6,
		lock:          sync.Mutex{},
	}
	err := snow.setMachineID(machineID)
	if err != nil {
		return &SNOW{}, err
	}
	return snow, err
}

func (snow *SNOW) setMachineID(id int64) error {
	if id > 1024 {
		return errors.New("Machine id must lower than 1024!")
	}
	snow.machineID = id << 12 // 左移12位变成机器号
	return nil
}

func (snow *SNOW) getID() int64 {
	// snow.lock.Lock()
	// defer snow.lock.Unlock()
	return snow.snowID()
}

func (snow *SNOW) snowID() int64 {
	curTimeStamp := time.Now().UnixNano() / 1e6
	if curTimeStamp == snow.lastTimeStamp { // 请求id的时间发生冲突
		// fmt.Println("current time is ", curTimeStamp, "last time is ", snow.lastTimeStamp, " cur snow is ", snow.snow)
		snow.snow++ // 防止冲突直接加一
		if snow.snow > 4095 {
			time.Sleep(time.Nanosecond) // 无法加一的时候直接睡眠并置0，一定不会冲突，相当于同时修改snow以及时间戳
			curTimeStamp = time.Now().UnixNano() / 1e6
			snow.lastTimeStamp = curTimeStamp
			snow.snow = 0
		}
		timeStampInSnmow := curTimeStamp & 0x1FFFFFFFFFF // 保留低41位的值
		timeStampInSnmow <<= 22                          // 作为时间戳的41位时间值
		return timeStampInSnmow | snow.machineID | snow.snow
	} else {
		snow.snow = 0
		snow.lastTimeStamp = curTimeStamp
		timeStampInSnmow := curTimeStamp & 0x1FFFFFFFFFF // 保留低41位的值
		timeStampInSnmow <<= 22                          // 作为时间戳的41位时间值
		return timeStampInSnmow | snow.machineID | snow.snow
	}
}
