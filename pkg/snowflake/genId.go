package snowflake

import (
	"fmt"
	"time"

	"github.com/sony/sonyflake"
)

var (
	sonyFlake     *sonyflake.Sonyflake
	sonyMachineId uint16
)

func Init(machineId uint16) (err error) {
	sonyMachineId = machineId
	t, _ := time.Parse("2006-01-02", "2022-11-17")
	settings := sonyflake.Settings{
		StartTime: t,
		MachineID: getMachineId,
	}
	sonyFlake = sonyflake.NewSonyflake(settings)
	return
}

func getMachineId() (uint16, error) {
	return sonyMachineId, nil
}

func GetId() (id uint64, err error) {
	if sonyFlake == nil {
		err = fmt.Errorf("sony flake not inited")
		return
	}
	id, err = sonyFlake.NextID()
	return
}
