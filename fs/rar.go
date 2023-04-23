package fs

import (
	"fmt"
	_ "io"
	_ "os"
	"strings"
	"errors"
	 
	"github.com/mholt/archiver/v3"
)

func readRar(f archiver.File, size int64) (string, error) {
	buffer := make([]byte, size)
	_, err  := f.ReadCloser.Read(buffer)
	if err != nil {
		return "", err
	}
	return string(buffer), nil	
}

func getRarPass(rar_file string) (string, error) {
	rar := archiver.Rar{
		ContinueOnError: false,
	}
	for _, pass := range passwords {
		rar.Password = pass
		err := rar.Walk(rar_file, func(f archiver.File) error {return nil})
		if err == nil {
			return pass, nil
		}
		fmt.Println(err)
	}
	return "", errors.New("can't guess the password. get the password and then update the passwords slice.")
}

func ListRar(rar_file string) ([]string, error) {
	ret := []string{}
	z := archiver.Rar{
		ContinueOnError: true,
		ImplicitTopLevelFolder: false,
	}
	guessed_pass, err := getRarPass(rar_file)
	if err != nil {
		return nil, err
	}
	z.Password = guessed_pass;
	_ = z.Walk(rar_file, func(f archiver.File) error {
		file_info := f.FileInfo
		splited_path := strings.Split(string(file_info.Name()), "/")
		file_name := strings.TrimSuffix(splited_path[len(splited_path)-1], "\n")
		if file_regex.MatchString(file_name) {
			size := file_info.Size()
			content, err := readRar(f, size)
			if err == nil {
				ret = append(ret, content)
			}
		}
		return nil
	})
	return ret, nil
}