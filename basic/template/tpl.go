package main

import (
	"bytes"
	"log"
	"text/template"
)

func main() {
	name := "world"
	tmpl, err := template.New("test").Parse("hello, {{.}}") //建立一个模板，内容是"hello, {{.}}"
	if err != nil {
		panic(err)
	}

	buf := new(bytes.Buffer)
	if err := tmpl.Execute(buf, name); err != nil { //将string与模板合成，变量name的内容会替换掉{{.}}
		panic(err)
	}
	log.Println(buf.String(), "=====")

	const tpl = `{
		"query": {
            "bool": {
            	"filter": {
					"bool": {
						"must": {
							"term": { "userId": {{.UserID}} }
						},
						"should": [
							{ "term": { "phash": "{{.Phash0}}" } },
							{ "term": { "phash": "{{.Phash1}}" } },
							{ "term": { "phash": "{{.Phash2}}" } },
							{ "term": { "phash": "{{.Phash3}}" } },
							{ "term": { "phash": "{{.Phash4}}" } },
							{ "term": { "phash": "{{.Phash5}}" } },
							{ "term": { "phash": "{{.Phash6}}" } },
							{ "term": { "phash": "{{.Phash7}}" } }
						],
						"minimum_should_match": 5
					}
				}
            }
        },
		"from": {{.From}},
		"size": {{.Size}}
	}`

	esreq, err := template.New("tpl").Parse(tpl)
	if err != nil {
		panic(err)
	}
	s := struct {
		UserID int64
		Phash0 string
		Phash1 string
		Phash2 string
		Phash3 string
		Phash4 string
		Phash5 string
		Phash6 string
		Phash7 string
		From   int64
		Size   int64
	}{
		UserID: 123,
		Phash0: "1_1",
		Phash1: "2_2",
		Phash2: "3_3",
		Phash3: "4_4",
		Phash4: "5_5",
		Phash5: "",
		Phash6: "",
		From:   0,
		Size:   10,
	}

	buff := new(bytes.Buffer)
	if err := esreq.Execute(buff, s); err != nil { //将string与模板合成，变量name的内容会替换掉{{.}}
		panic(err)
	}
	log.Println(buff.String())

}
