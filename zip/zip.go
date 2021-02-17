package zip

import (
	"archive/zip"
	"fmt"
	"io/ioutil"
	"os"
)

// Compress compress directory and export zip archive
func Compress(dst, src string) error {
	outFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer outFile.Close()

	// Create a new zip archive.
	w := zip.NewWriter(outFile)

	// Add some files to the archive.
	var errStr = addFiles(w, src+string(os.PathSeparator), "")
	if errStr != "" {
		return fmt.Errorf(errStr)
	}

	// Make sure to check the error on Close.
	err = w.Close()
	if err != nil {
		return err
	}

	return nil
}

func addFiles(w *zip.Writer, basePath, baseInZip string) string {
	var errStr string

	// Open the Directory
	files, err := ioutil.ReadDir(basePath)
	if err != nil {
		errStr += err.Error() + "\n"
	}

	for _, file := range files {
		if !file.IsDir() {
			dat, err := ioutil.ReadFile(basePath + file.Name())
			if err != nil {
				errStr += err.Error() + "\n"
			}

			// Add some files to the archive.
			f, err := w.Create(baseInZip + file.Name())
			if err != nil {
				errStr += err.Error() + "\n"
			}
			_, err = f.Write(dat)
			if err != nil {
				errStr += err.Error() + "\n"
			}
		} else if file.IsDir() {
			newBase := basePath + file.Name() + string(os.PathSeparator)

			errStr += addFiles(w, newBase, baseInZip+file.Name()+"/")
		}
	}

	return errStr
}
