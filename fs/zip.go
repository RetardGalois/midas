package fs

import (
	"io/ioutil"
	"github.com/yeka/zip"
	"strings"
	"errors"
)

func readZip(file *zip.File, password string) (string, error) {
	if password != "" {
		file.SetPassword(password)
	}
	fc, err := file.Open()
	if err != nil {
		return "", err
	}
	defer fc.Close()
	content, err := ioutil.ReadAll(fc)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func getZipPass(file *zip.File) (string, error) {
	for _, pass := range passwords {
		file.SetPassword(pass)
		if _, err := readZip(file, pass); err == nil {
			return pass, nil
		}
	}
	return "", errors.New("unable to guess the password. get the password and then update the passwords slice.")
}

func ListZip(zip_file string) ([]string, error) {
	ret := []string{}
	zf, err := zip.OpenReader(zip_file)
	if err != nil {
		return nil, err
	}
	defer zf.Close()

	possible_pass := ""
	for _, file := range zf.File {
		splited_path := strings.Split(file.Name, "/")
		file_name := strings.TrimSuffix(splited_path[len(splited_path)-1], "\n")
		if file_regex.MatchString(file_name) {
			if (file.IsEncrypted() && possible_pass == "") {
				guessed_pass, err := getZipPass(file)
				if err == nil {
					possible_pass = guessed_pass
				} else {
					return nil, err
				}
			}
			content, err := readZip(file, possible_pass)
			if err != nil {
				//panic(fmt.Sprintf("Can't read file. ERROR: %s", err))
				continue
			}
			ret = append(ret, content)
		}
	}
	return ret, nil
}
 
