package model

import "dfetch/internal/getsysinfo"

type SystemInfo struct {
	DistroName   string
	ID           string
	Kernel       string
	CPU          string
	Memory       string
	Username     string
	Hostname     string
	LocalIP      string
	IPVersion    string
	Uptime       string
	Battery      int
	BatteryState string
	DE           string
	SessionType  string
	Shell        string
}

func CollectSystemInfo() SystemInfo {
	DistroName, id := getsysinfo.Distro()
	localIP := getsysinfo.LocalIP()
	battery, batteryStatus := getsysinfo.Battery()

	de, sessionType := getsysinfo.DesktopEnvironment()

	return SystemInfo{
		DistroName:   DistroName,
		ID:           id,
		Kernel:       getsysinfo.Kernel(),
		CPU:          getsysinfo.Cpu(),
		Memory:       getsysinfo.Memory(),
		Username:     getsysinfo.Username(),
		Hostname:     getsysinfo.Hostname(),
		LocalIP:      localIP,
		Uptime:       getsysinfo.Uptime(),
		Battery:      battery,
		BatteryState: batteryStatus,
		DE:           de,
		SessionType:  sessionType,
		Shell:        getsysinfo.Shell(),
	}
}
