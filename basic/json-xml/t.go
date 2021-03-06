package main

import (
	"fmt"

	"github.com/tidwall/gjson"
)

const json = `{
    "result": {
        "bc6df421ed9c43b57b74c5978185c75b": {
            "organization": "",
            "organizationUnit": "",
            "commonName": "qiaohaohuo.com",
            "validity": {
                "start": 1577667911000,
                "end": 1585443911000
            },
            "config_domain": 4,
            "crt_type": "custom",
            "payment_type": "free",
            "brand": "",
            "crt_type_pay": "",
            "pro": false,
            "domain_type": "",
            "created_at": 1577677112
        },
        "1ae590de09bce57cd99e6f92c4ab6277": {
            "organization": "",
            "organizationUnit": "",
            "commonName": "qiaohaohuo.com",
            "validity": {
                "start": 1577690060000,
                "end": 1585466060000
            },
            "config_domain": 0,
            "crt_type": "custom",
            "payment_type": "free",
            "brand": "",
            "crt_type_pay": "",
            "pro": false,
            "domain_type": "",
            "created_at": 1577693747
        },
        "8a779cbdde7df45bdac91ca59f7b225a": {
            "organization": "",
            "organizationUnit": "",
            "commonName": "qiaohaohuo.com",
            "validity": {
                "start": 1578292269000,
                "end": 1586068269000
            },
            "config_domain": 0,
            "crt_type": "custom",
            "payment_type": "free",
            "brand": "",
            "crt_type_pay": "",
            "pro": false,
            "domain_type": "",
            "created_at": 1578295876
        }
    },
    "status": 0
}`

const esjson = `
{
    "took": 142,
    "timed_out": false,
    "_shards": {
      "total": 3,
      "successful": 3,
      "skipped": 0,
      "failed": 0
    },
    "hits": {
      "total": 42538909,
      "max_score": 1,
      "hits": [
        {
          "_index": "albums",
          "_type": "doc",
          "_id": "20308943",
          "_score": 1
        },
        {
          "_index": "albums",
          "_type": "doc",
          "_id": "20308947",
          "_score": 1
        },
        {
          "_index": "albums",
          "_type": "doc",
          "_id": "20308957",
          "_score": 1
        },
        {
          "_index": "albums",
          "_type": "doc",
          "_id": "20308970",
          "_score": 1
        },
        {
          "_index": "albums",
          "_type": "doc",
          "_id": "20309002",
          "_score": 1
        },
        {
          "_index": "albums",
          "_type": "doc",
          "_id": "20309012",
          "_score": 1
        },
        {
          "_index": "albums",
          "_type": "doc",
          "_id": "20309015",
          "_score": 1
        },
        {
          "_index": "albums",
          "_type": "doc",
          "_id": "20309030",
          "_score": 1
        },
        {
          "_index": "albums",
          "_type": "doc",
          "_id": "20309051",
          "_score": 1
        },
        {
          "_index": "albums",
          "_type": "doc",
          "_id": "20309053",
          "_score": 1
        }
      ]
    }
  }
`

func main() {
	result := gjson.Get(string(json), `result`).Map()

	fmt.Println(len(result))
	var deleteCertID string
	var oldest int64

	for k, v := range result {
		createdAt := gjson.Get(v.Raw, "created_at").Int()
		if oldest == 0 || oldest > createdAt {
			deleteCertID = k
			oldest = createdAt
		}

		configDomain := gjson.Get(v.Raw, "config_domain").Int()
		if configDomain > 0 {
			println(configDomain)
		}
	}

	fmt.Println(deleteCertID, oldest)

	total := gjson.Get(esjson, "hits.total").Int()
	fmt.Println("es total: ", total)

	_ids := gjson.Get(esjson, "hits.hits.#._id").Array()
	for _, v := range _ids {
		fmt.Println(v)
	}
}
