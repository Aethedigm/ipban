package ipban

import (
	"bufio"
	"net/http"
	"os"
)

func IPBan(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		banListLoc := os.Getenv("BAN_LIST")
		banList := []string{}

		banListFile, err := os.Open(banListLoc)
		if err != nil {
			panic(err)
		}

		scanner := bufio.NewScanner(banListFile)
		for scanner.Scan() {
			banList = append(banList, scanner.Text())
		}

		if IPInList(r.RemoteAddr, banList) {
			http.Error(w, "Forbidden", http.StatusForbidden)
			panic("IP Banned")
		}
	}

	return http.HandlerFunc(fn)
}

func IPInList(ip string, list []string) bool {
	for _, v := range list {
		if v == ip {
			return true
		}
	}
	return false
}
