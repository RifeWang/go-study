package main

import (
	"fmt"
	"log"
	"regexp"
	"time"
)

func main() {
	// Compile the expression once, usually at init time.
	// Use raw strings to avoid having to quote the backslashes.
	var validID = regexp.MustCompile(`^[a-z]+\[[0-9]+\]$`)

	fmt.Println(validID.MatchString("adam[23]"))
	fmt.Println(validID.MatchString("eve[7]"))
	fmt.Println(validID.MatchString("Job[48]"))
	fmt.Println(validID.MatchString("snakey"))

	re := regexp.MustCompile(`foo.?`)
	fmt.Printf("%q\n", re.Find([]byte(`seafood fool`)))

	urlRegexp := regexp.MustCompile(`filename=(([a-z0-9-.]+)\_(\d{8}\_\d{2}\.\d{2})\_\d{2}\.\d{2})`)
	strs := urlRegexp.FindStringSubmatch("https://runreport.dwion.com/logdown?param=93E9D3E337D3E50D3E551C459B48CB536DD59649167048863AD0135F16241C80C9EC2D186D72FE404E7559A494D32EBFB2422187787DD056DF3673F8AF26BDCA&filename=upmov.a.yximgs.com_20200421_23.55_24.00")
	fmt.Println(strs)

	file := strs[1]
	domain := strs[2]
	date := strs[3]

	t, _ := time.Parse("20060102_15.04", date)

	fmt.Println(file, domain, date, t.Format("200601021504"))

	log.Println("=============================")
	log.Println("=============================")
	log.Println("=============================")

	cm := "https://logdownload.cmcdn.cdn.10086.cn/bj/logdown/upmov.a.yximgs.com_20200420000500_01_25_3.log.gz?path=bG9nZG93bi9pbnRpbWUvMjAyMDA0MjAvdXBtb3YuYS55eGltZ3MuY29tL3VwbW92LmEueXhpbWdzLmNvbV8yMDIwMDQyMDAwMDUwMF8wMV8yNV8zLmxvZy5neg=="
	cmRegexp := regexp.MustCompile(`(([a-z0-9-.]+)\_(\d{12})\d{2}\_[0-9_]+\.log\.gz)`)

	result := cmRegexp.FindStringSubmatch(cm)
	filename := result[1]

	doma := result[2]
	datetime := result[3]
	fmt.Println(result)
	fmt.Println(filename, doma, datetime)
}
