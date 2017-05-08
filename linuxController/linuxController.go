package linuxController

import (
	"os/exec"
)

func MuteVolume(mute bool) {
	if mute == true {
		exec.Command("amixer", "-D", "pulse", "sset", "Master", "mute").Run()
	} else {
		exec.Command("amixer", "-D", "pulse", "sset", "Master", "unmute").Run()
	}
}

func GetUser() string {
	username, _ := exec.Command("echo", "$USER").Output()
	return string(username)
}

func TurnVolumeUp(volumeIncrease string) {
	var formatedCommand string = volumeIncrease + "%+"
	exec.Command("amixer", "-D", "pulse", "sset", "Master", formatedCommand).Run()
}

func TurnVolumeDown(volumeDecrease string) {
	var formatedCommand string = volumeDecrease + "%-"
	exec.Command("amixer", "-D", "pulse", "sset", "Master", formatedCommand).Run()
}

func TriggerLockScreen() {
	exec.Command("gnome-screensaver-command", "-l").Start()
}

func RestartComputer(password string) {
	command := exec.Command("sudo", "-S", "reboot")
	stdin, _ := command.StdinPipe()
	stdin.Write([]byte(password))
	stdin.Close()
	stuff.Run()
}

func UpdateComputer(password string) {
	command := exec.Command("sudo", "-S", "apt", "upgrade")
	stdin, _ := command.StdinPipe()
	stdin.Write([]byte(password))
	stdin.Close()
	stuff.Run()
}

func ShutDownComputer(password string) {
	command := exec.Command("sudo", "-S", "poweroff")
	stdin, _ := command.StdinPipe()
	stdin.Write([]byte(password))
	stdin.Close()
	stuff.Run()
}
