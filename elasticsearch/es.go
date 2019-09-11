package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"gopkg.in/olivere/elastic.v5"
)

type esdoc struct {
	UserID int64  `json:"userId"`
	Path   string `json:"path"`
	Phash  string `json:"phash"`
}

func main() {
	client, err := elastic.NewClient(
		elastic.SetURL("http://10.0.6.49:9201"),
	)
	if err != nil {
		log.Fatal(err)
	}

	exists, err := client.IndexExists("dypic").Do(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	if !exists {
		// Index does not exist yet.
	}
	log.Println(exists, err)

	// index 覆盖创建索引
	// 使用 JSON 序列化，配合 BodyJson
	doc := esdoc{UserID: 123, Path: "/qwer/tyui", Phash: "1_1 2_2 3_3 4_4 5_5 6_6 7_7 8_8 9_9"}
	log.Println("=============", doc, &doc)

	fc := func(d esdoc) {
		put1, err := client.Index().Index("dypic").Type("doc").Id("1").BodyJson(d).Do(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Indexed id %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)
	}
	fc(doc)

	// 使用 string ，配合 BodyString
	doc2 := `{"userId": 123, "path": "/image/qqqqqqqqqq", "phash": "1_1 2_2 3_3 4_4 5_5 6_6 7_7 8_8 9_9"}`
	put2, err := client.Index().Index("dypic").Type("doc").Id("2").BodyString(doc2).Do(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Indexed id %s to index %s, type %s\n", put2.Id, put2.Index, put2.Type)

	// Search with a term query
	// termQuery := elastic.NewTermQuery("userId", "123")
	// searchResult, err := client.Search().
	// 	Index("twitter").        // search in index "twitter"
	// 	Query(termQuery).        // specify the query
	// 	Sort("user", true).      // sort by "user" field, ascending
	// 	From(0).Size(10).        // take documents 0-9
	// 	Pretty(true).            // pretty print request and response JSON
	// 	Do(context.Background()) // execute
	// if err != nil {
	// 	// Handle error
	// 	panic(err)
	// }

	// es 包的方法适用于简单请求，对于复杂的请求不好处理，建议直接使用 http 请求
	httpclient := &http.Client{}

	// _source: false 不返回源数据
	reqBody := []byte(`{
		"_source": false,
		"query": {
            "bool": {
            	"filter": {
					"bool": {
						"must": {
							"term": { "userId": 123 }
						},
						"should": [
							{ "term": { "phash": "1_1" } },
							{ "term": { "phash": "2_2" } },
							{ "term": { "phash": "3_3" } },
							{ "term": { "phash": "4_4" } },
							{ "term": { "phash": "5_5" } },
							{ "term": { "phash": "6_6" } },
							{ "term": { "phash": "7" } },
							{ "term": { "phash": "8" } }
						],
						"minimum_should_match": 5
					}
				}
            }
        },
		"from": 0,
		"size": 10
	}`)

	req, err := http.NewRequest("GET", "http://10.0.6.49:9201/dypic/doc/_search", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	// 客户端发起请求
	res, err := httpclient.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close() // 必须关闭，避免内存泄漏
	log.Println("response: ", res, err)

	result := ESResponse{}

	bts, _ := ioutil.ReadAll(res.Body)
	if err := json.Unmarshal(bts, &result); err != nil { // unmarshal 将 json 数据解码
		fmt.Println(err)
	}

	fmt.Println("response Body:", result)
	fmt.Println("response Body:", result.Hits.Hits[0])

}

// ESResponse ... ES search 返回的数据结构
type ESResponse struct {
	Took     int64 `json:"took"`
	TimedOut bool  `json:"timed_out"`
	Shards   struct {
		Total      int64 `json:"total"`
		Successful int64 `json:"successful"`
		Skipped    int64 `json:"skipped"`
		Failed     int64 `json:"failed"`
	}
	Hits struct {
		Total    int64   `json:"total"`
		MaxScore float32 `json:"max_score"`
		Hits     []struct {
			Index  string      `json:"_index"`
			Type   string      `json:"_type"`
			ID     string      `json:"_id"`
			Score  float32     `json:"_score"`
			Source interface{} `json:"_source"`
		}
	}
}
