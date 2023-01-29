package readfiles

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

type FilesInfo struct {
	IsFolder bool
	Path     string
}

func ReadFolder(folder string) []FilesInfo {
	var (
		root  string
		files []FilesInfo
		err   error
	)

	root = folder
	// filepath.Walk
	files, err = FilePathWalkDir(root)
	if err != nil {
		panic(err)
	}
	// ioutil.ReadDir
	// files, err = IOReadDir(root)
	// if err != nil {
	// 	panic(err)
	// }
	// //os.File.Readdir
	// files, err = OSReadDir(root)
	// if err != nil {
	// 	panic(err)
	// }
	//
	// for _, file := range files {
	// 	fmt.Println(file)
	// }

	return files
}

func FilePathWalkDir(root string) ([]FilesInfo, error) {
	var files []string
	var fileStruct []FilesInfo
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {

		if !info.IsDir() {
			files = append(files, path)
			fileStruct = append(fileStruct, FilesInfo{IsFolder: false, Path: path})
		}

		return nil
	})
	return fileStruct, err
}

func IOReadDir(root string) ([]string, error) {
	var files []string
	fileInfo, err := ioutil.ReadDir(root)
	if err != nil {
		return files, err
	}

	for _, file := range fileInfo {
		files = append(files, file.Name())
	}
	return files, nil
}

func OSReadDir(root string) ([]string, error) {
	var files []string
	f, err := os.Open(root)
	if err != nil {
		return files, err
	}
	fileInfo, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		return files, err
	}

	for _, file := range fileInfo {
		files = append(files, file.Name())
	}
	return files, nil
}
