package main

import (
  "encoding/json"
  "net/http"
  "fmt"
  "io/ioutil"
  "github.com/mattn/go-jsonpointer"
)

func main() {
  url := "https://raw.githubusercontent.com/codeforosaka/covid19/master/data/data.json"
  resp, _ := http.Get(url)
  defer resp.Body.Close()
  byteArray, _ := ioutil.ReadAll(resp.Body)

  var positives interface{}
  json.Unmarshal(byteArray, &positives)

  // 年月日
  d, _ := jsonpointer.Get(positives, "/lastUpdate")
  // 重症
  jyuu, _ := jsonpointer.Get(positives,
    "/main_summary/children/0/children/0/children/1/value")
  // 軽症・中等症 ( 入院・入院調整中 - 重症 )
  tmp, _ := jsonpointer.Get(positives,
              "/main_summary/children/0/children/0/value")
  kei := tmp.(float64) - jyuu.(float64)
  // 症状不明
  //     自宅療養
  tmp, _ = jsonpointer.Get(positives,
              "/main_summary/children/0/children/3/value")
  fumei := tmp.(float64)
  //     宿泊療養
  tmp, _ = jsonpointer.Get(positives,
              "/main_summary/children/0/children/4/value")
  fumei += tmp.(float64)
  //     療養等調整中
  tmp, _ = jsonpointer.Get(positives,
              "/main_summary/children/0/children/5/value")
  fumei += tmp.(float64)
  //     入院調整中
  tmp, _ = jsonpointer.Get(positives,
              "/main_summary/children/0/children/6/value")
  fumei += tmp.(float64)
  //     府外健康観察
  tmp, _ = jsonpointer.Get(positives,
              "/main_summary/children/0/children/7/value")
  fumei += tmp.(float64)
  // 死亡
  shibou, _ := jsonpointer.Get(positives,
              "/main_summary/children/0/children/2/value")  // 自宅療養
  // 退院
  taiin, _ := jsonpointer.Get(positives,
              "/main_summary/children/0/children/1/value")

  fmt.Printf("%v, %v, %v, %v,, %v, %v\n", d, kei, fumei, jyuu, shibou, taiin)

}

