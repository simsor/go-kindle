package kindle

// GetWifiState returns the current wireless connection state
func GetWifiState() string {
	state, err := RawLIPC("com.lab126.wifid", "cmState")
	if err != nil {
		return "ERROR"
	}

	return state
}

// GetWifiESSID returns the ESSID of the current WiFi network
func GetWifiESSID() string {
	hash, err := RawLIPCHashArray("com.lab126.wifid", "currentEssid")
	if err != nil {
		return ""
	}

	if len(hash) > 0 {
		return hash[0]["essid"]
	}
	return ""
}
