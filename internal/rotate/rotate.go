package rotate

import (
	"archive/tar"
	"errors"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

func PackFile(filepath, directory string) error {
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Error open file:", err)
		return err
	}
	defer file.Close()

	fileContent, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println("Error open file:", err)
		return err
	}

	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("Error open file:", err)
		return err
	}

	tarFilename, found := strings.CutSuffix(fileInfo.Name(), ".log")
	if !found {
		fmt.Println("Not found .log files", err)
		return errors.New("log file is missing")
	}
	tarFilename += "_" + strconv.Itoa(int(time.Now().Unix())) + ".tar.gz"
	if directory != "" {
		err = os.MkdirAll(directory, 0777)
		if err != nil {
			fmt.Println("Error creating tar directory:", err)
			return err
		}
	}

	tarFile, err := os.Create(path.Join(directory, tarFilename))
	if err != nil {
		fmt.Println("Error creating tar file:", err)
		return err
	}
	defer tarFile.Close()

	tarWriter := tar.NewWriter(tarFile)
	defer tarWriter.Close()
	header := &tar.Header{
		Name:    fileInfo.Name(),
		Size:    fileInfo.Size(),
		Mode:    int64(fileInfo.Mode()),
		ModTime: fileInfo.ModTime(),
	}

	err = tarWriter.WriteHeader(header)
	if err != nil {
		fmt.Println("Error writing tar header:", err)
		return errors.New("error writing tar header")
	}

	_, err = tarWriter.Write(fileContent)
	if err != nil {
		fmt.Println("Error writing file to tar:", err)
		return errors.New("error writing file to tar")
	}

	return nil
}
