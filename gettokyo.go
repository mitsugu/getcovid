package main

import (
  "encoding/json"
  "net/http"
  "fmt"
  "io/ioutil"
  "github.com/mattn/go-jsonpointer"
)

func main() {
  url := "https://raw.githubusercontent.com/tokyo-metropolitan-gov/covid19/development/data/data.json"
  resp, _ := http.Get(url)
  defer resp.Body.Close()
  byteArray, _ := ioutil.ReadAll(resp.Body)

  var positives interface{}
  json.Unmarshal(byteArray, &positives)

  // 年月日
  d, _ := jsonpointer.Get(positives, "/lastUpdate")
  // 軽症・中等症
  keisyou, _ := jsonpointer.Get(positives,
              "/main_summary/children/0/children/0/children/0/value")
  // 重症
  jyuusyou, _ := jsonpointer.Get(positives,
              "/main_summary/children/0/children/0/children/1/value")
  // 退院
  taiin, _ := jsonpointer.Get(positives,
              "/main_summary/children/0/children/1/value")
  // 死亡
  shibou, _ := jsonpointer.Get(positives,
              "/main_summary/children/0/children/2/value")
  // 宿泊療養
  syukuhaku, _ := jsonpointer.Get(positives,
              "/main_summary/children/0/children/3/value")
  // 自宅療養
  jitaku, _ := jsonpointer.Get(positives,
              "/main_summary/children/0/children/4/value")
  // 調整中 ( 調査中 )
  cyousei, _ := jsonpointer.Get(positives,
              "/main_summary/children/0/children/5/value")

  fmt.Printf("%v, %v, %v, %v, %v, %v,, %v, %v\n", d, cyousei, jitaku, syukuhaku, keisyou, jyuusyou, shibou, taiin)
}

