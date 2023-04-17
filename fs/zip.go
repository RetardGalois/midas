package fs

import (
	"io/ioutil"
	"github.com/yeka/zip"
	"strings"
	"errors"
)

var passwords = []string{
	"Borwita",
	"GoldenEagleCloud",
	"Free",
	"OTTOMAN",
	"@Daisy_Cloud",
	"ygygg777",
	"OTTOMAN**1115",
	"OTTOLIKE",
	"@cheatbreakercloud",
	"https://t.me/MariaLogs",
	"@crypton_logs",
	"@TOR_LOG",
	"@RaccoonCloud",
	"@FASTCLOUDD",
	"@observercloud",
	"@MikaLogsfree",
	"https://t.me/REDLINEPREMIUM/",
	"https://t.me/+WlFsI97nclo2ZjJl",
	"@freshlogs1only",
	"LOGS",
	"PASS",
	"redlinevip",
	"@freshlogs1only",
	"t.me/stake_logs",
	"https://t.me/Zeuscloudfree",
	"@NoxyCloud",
	"@segacloud",
	"@spiderLogs",
	"https://t.me/EMClouds",
	"https://t.me/PremCloud",
	"https://t.me/prdscloud",
	"https://t.me/magicianlogs",
	"https://t.me/tor_log",
	"https://t.me/crypton_logs",
	"https://t.me/Wooden_Cloud",
	"@tokyocloudd",
	"https://t.me/Heavenlogscloud",
	"@milkywaycloud",
	"https://t.me/tor_log",
	"t.me/xl0gs",
	"@sigmcloud",
	"@zet_cloud",
	"https://t.me/klaus_cloud_public",
	"https://t.me/RedlineClouds1",
	"https://t.me/ObserverPrivate",
}


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

func getPass(file *zip.File) (string, error) {
	for _, pass := range passwords {
		file.SetPassword(pass)
		if _, err := readZip(file, pass); err == nil {
			return pass, nil
		}
	}
	return "", errors.New("can't guess the password.")
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
				guessed_pass, err := getPass(file)
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
 
