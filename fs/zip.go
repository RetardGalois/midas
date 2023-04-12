package fs

import (
	"io/ioutil"
	"archive/zip"
	"strings"
)

func readZip(file *zip.File) (string, error) {
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

func ListZip(zip_file string) ([]string, error) {
	ret := []string{}
	zf, err := zip.OpenReader(zip_file)
	if err != nil {
		return nil, err
	}
	defer zf.Close()

	for _, file := range zf.File {
		splited_path := strings.Split(file.Name, "/")
		file_name := strings.TrimSuffix(splited_path[len(splited_path)-1], "\n")
		if file_regex.MatchString(file_name) {
			content, err := readZip(file)
			if err != nil {
				//panic(fmt.Sprintf("Can't read file. ERROR: %s", err))
				continue
			}
			ret = append(ret, content)
		}
	}
	return ret, nil
}
 
