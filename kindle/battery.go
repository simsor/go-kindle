package kindle

import (
	"strconv"
)

// GetBatteryLevel returns the current percentage of the battery
func GetBatteryLevel() int {
	res, err := RawLIPC("com.lab126.powerd", "battLevel")
	if err != nil {
		return 0
	}

	lvl, err := strconv.Atoi(res)
	if err != nil {
		return 0
	}

	return lvl
}
