package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/olivere/elastic/v7"
)

var ctx = context.Background()
var esUrl = "http://192.168.72.128:9201/"
var EsClient *elastic.Client

// 初始化es连接
func InitEs() {
	// 连接es客户端
	client, err := elastic.NewClient(
		elastic.SetURL(esUrl),
	)
	if err != nil {
		log.Fatal("es 连接失败:", err)
	}
	info, code, err := client.Ping(esUrl).Do(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println("Elasticsearch returned with code>: %d and version %s\n", code, info.Version.Number)
	EsClient = client
}
func init() {
	InitEs()
}

// 定义数据结构体
type data struct {
	Id    string  `json:"id"`
	Icon  string  `json:"icon"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// 定义一些变量，mapping为定制的index字段类型
const mapping = `
  {
  	"settings":{
  		"number_of_shards": 1,
  		"number_of_replicas": 0
  	},
  	"mappings":{
  		"properties":{
  			"name":{
  				"type":"keyword"
  			},
  			"icon":{
  				"type":"text"
  			},
  			"price":{
  				"type":"double"
  			},
  			"id":{
  				"type":"text"
  			}
  		}
  	}
  }`

// 添加文档
func Addindex(table string) (bool, error) {
	// 创建index前，先查看es引擎中是否存在自己想要创建的索引index
	exists, err := EsClient.IndexExists(table).Do(ctx)
	if err != nil {
		fmt.Println("存在索引:", err)
		return true, nil
	}
	if !exists {
		// 如果不存在，就创建
		createIndex, err := EsClient.CreateIndex(table).BodyString(mapping).Do(ctx)
		if err != nil {
			return false, err
		}
		if !createIndex.Acknowledged {
			return false, err
		}
	}
	return true, nil
}

// 添加数据
func Add(table string, data interface{}) (bool, error) {
	// 添加索引
	_, err := Addindex(table)
	if err != nil {
		log.Fatal("创建索引失败", err)
	}
	// 添加文档
	res, err := EsClient.Index().
		Index(table).
		BodyJson(data).
		Do(ctx)
	if err != nil {
		return false, err
	}
	fmt.Println("添加数据成功:", res)
	return true, nil
}
func BulkAdd(table string, d []data) (bool, error) {
	// 添加索引
	_, err := Addindex(table)
	if err != nil {
		log.Fatal("创建索引失败", err)
	}
	bulkReq := EsClient.Bulk()
	for _, v := range d {
		req := elastic.NewBulkIndexRequest().
			Index(table).
			Doc(v)
		bulkReq = bulkReq.Add(req)
	}
	res, err := bulkReq.Do(ctx)
	if err != nil {
		return false, err
	}
	fmt.Println("添加数据成功:", res)
	return true, nil
}
func Query(table string, field string, filter elastic.Query, sort string, page int, limit int) (*elastic.SearchResult, error) {
	// 分页数据处理
	isAsc := true
	if sort != "" {
		sortSlice := strings.Split(sort, " ")
		sort = sortSlice[0]
		if sortSlice[1] == "desc" {
			isAsc = false
		}
	}
	// 查询位置处理
	if page <= 1 {
		page = 1
	}

	fsc := elastic.NewFetchSourceContext(true)
	// 返回字段处理
	if field != "" {
		fieldSlice := strings.Split(field, ",")
		if len(fieldSlice) > 0 {
			for _, v := range fieldSlice {
				fsc.Include(v)
			}
		}
	}

	// 开始查询位置
	fromStart := (page - 1) * limit
	res, err := EsClient.Search().
		Index(table).
		FetchSourceContext(fsc).
		Query(filter).
		Sort(sort, isAsc).
		From(fromStart).
		Size(limit).
		Pretty(true).
		Do(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func UpWhere(table string, filter elastic.Query, data map[string]interface{}) (bool, error) {
	// 修改数据组装
	if len(data) < 0 {
		return false, errors.New("修改参数不正确")
	}
	scriptStr := ""
	for k := range data {
		scriptStr += "ctx._source." + k + " = params." + k + ";"
	}
	script := elastic.NewScript(scriptStr).Params(data)
	res, err := EsClient.UpdateByQuery(table).
		Query(filter).
		Script(script).
		Do(ctx)
	if err != nil {
		return false, err
	}
	fmt.Println("更新数据成功:", res)
	return true, nil
}
func Del(table string, filter elastic.Query) (bool, error) {
	res, err := EsClient.DeleteByQuery().
		Query(filter).
		Index(table).
		Do(ctx)
	if err != nil {
		return false, err
	}
	fmt.Println("删除信息：", res)
	return true, nil
}

func main() {
	//添加索引
	//res, err := Addindex("product")
	//fmt.Println(res, err)

	//添加数据
	//data := data{
	//	Id:    "1",
	//	Icon:  "头像",
	//	Name:  "名称",
	//	Price: 9.99,
	//}
	//res, err := Add("product", data)
	//fmt.Println(res, err)

	//批量添加
	//var d []data
	//for i := 0; i < 100; i++ {
	//	iStr := strconv.Itoa(i)
	//	v := data{
	//		Id:    iStr,
	//		Icon:  "icon " + iStr,
	//		Name:  "name " + iStr,
	//		Price: float64(i),
	//	}
	//	d = append(d, v)
	//}
	//res, err := BulkAdd("product", d)

	// 精准查询
	//filter := elastic.NewTermQuery("name", "李白")

	// 分词匹配查询
	// match_all 查询所有
	//filter := elastic.NewMatchAllQuery()
	// match 单个字符匹配
	//filter := elastic.NewMatchQuery("name", "名称")
	// multi_match 多个字段进行匹配查询
	//filter := elastic.NewMultiMatchQuery("白日", "name", "icon")
	// match_phrase 短语匹配
	//filter := elastic.NewMatchPhraseQuery("icon", "途穷反遭俗眼白")

	// fuzzy模糊查询 fuzziness是可以允许纠正错误拼写的个数
	//filter := elastic.NewFuzzyQuery("icon", "夜").Fuzziness(1)

	// wildcard 通配符查询
	//filter := elastic.NewWildcardQuery("name", "静夜*")

	//table := "product"
	//sort := "price asc"
	//page := 0
	//limit := 10
	//field := "name,icon,price"
	//res, err := Query(table, field, filter, sort, page, limit)
	//strD, _ := json.Marshal(res)
	//fmt.Println(string(strD), err)

	//更新数据
	//table := "product"
	//filter := elastic.NewTermQuery("name", "名称")
	//data := make(map[string]interface{})
	//data["icon"] = "hello world"
	//data["name"] = "name 00"
	//res, err := UpWhere(table, filter, data)
	//
	//fmt.Println(res, err)

	//删除数据
	filter := elastic.NewTermQuery("name", "name 1")
	res, err := Del("product", filter)
	fmt.Println(res, err)
}

/***
elastic.SetSniff(false)
no active connection found: no Elasticsearch node available
*/
