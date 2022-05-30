package downloading

import (
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	e "wbio/malware/errorinfo"
)

func DownloadMalware(targetDirectory string, fileName string) {

	path := filepath.FromSlash(targetDirectory + "/" + fileName)

	if isMalwareDownloaded(path) {
		log.Printf("[DOWNLOAD] Malware already downloaded (%s)", path)
		return
	}

	out, err := os.Create(path)
	e.Check(err, true)
	defer out.Close()

	url := "https://secure.eicar.org/eicar.com.txt"
	log.Printf("[DOWNLOAD] Downloading malware from %s\n", url)
	resp, err := http.Get(url)
	e.Check(err, true)
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	e.Check(err, true)

	log.Printf("[DOWNLOAD] Malware downloaded and saved to %s\n", path)
}

func isMalwareDownloaded(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	} else if errors.Is(err, os.ErrNotExist) {
		return false
	} else {
		e.Check(err, true)
		return false
	}
}
