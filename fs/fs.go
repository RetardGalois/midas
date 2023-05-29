package fs

import (
	"strings"
	"regexp"
	"errors"
)

var file_regex = regexp.MustCompile(`(?i)\bpass.*\.txt\b`)

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

func File(file_name string) ([]string, error) {
	if strings.HasSuffix(file_name, "zip") {
		return ListZip(file_name)
	} else if strings.HasSuffix(file_name, "rar") {
		return ListRar(file_name)
	}
	return nil, errors.New("not suported file type.")
}

