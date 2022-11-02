package gdb

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"strings"
	"sync"

	"github.com/elastic/go-elasticsearch/esapi"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/eoe2005/g/gconf"
)

var (
	_esConf  = map[string]*gconf.GDbYaml{}
	_localEs = map[string]*elasticsearch.Client{}
)

type EsObject struct {
	Id   string
	Data any
}

func initEs(item *gconf.GDbYaml) {
	_esConf[item.Name] = item
}

func getEs(name string) *elasticsearch.Client {
	if r, ok := _localEs[name]; ok {
		return r
	}
	if conf, ok := _esConf[name]; ok {
		clis, err := elasticsearch.NewClient(elasticsearch.Config{
			Addresses: strings.Split(conf.Host, ","),
			Username:  conf.UserName,
			Password:  conf.UserPass,
		})
		if err != nil {
			panic("创建链接失败")
		}
		_localEs[name] = clis
		return clis
	}
	panic("配置不存在")
}
func EsSave(ctx context.Context, name, index string, data ...EsObject) []EsObject {
	unsaves := []EsObject{}
	client := getEs(name)
	wg := sync.WaitGroup{}
	for _, i := range data {
		go func(item EsObject, c *elasticsearch.Client, unsave []EsObject) {
			defer wg.Done()
			wg.Add(1)
			b, e := json.Marshal(item.Data)
			if e != nil {
				unsave = append(unsave, item)
				return
			}
			res, e := esapi.IndexRequest{
				Index:      index,
				Body:       bytes.NewReader(b),
				DocumentID: item.Id,
				Refresh:    "true",
			}.Do(ctx, c)
			if e != nil {
				unsave = append(unsave, item)
				return
			}
			defer res.Body.Close()
			if res.IsError() {
				unsave = append(unsave, item)
				return
			}
		}(i, client, unsaves)

	}
	wg.Wait()
	return unsaves

}
func EsDelete(ctx context.Context, name, index string, ids ...string) {
	// es := getEs(name)
	// es.DeleteByQuery
}
func EsQuery(ctx context.Context, name, index string, query map[string]any, target any) error {
	es := getEs(name)
	var buf bytes.Buffer
	e := json.NewEncoder(&buf).Encode(query)
	if e != nil {
		return e
	}
	res, err := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex(index),
		es.Search.WithBody(&buf),
		es.Search.WithTrackTotalHits(true),
		es.Search.WithPretty(),
	)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.IsError() {
		return errors.New("请求失败")
	}
	return json.NewDecoder(res.Body).Decode(target)
}
