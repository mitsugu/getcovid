package main

import (
  "encoding/json"
  "net/http"
  "fmt"
  "io/ioutil"
  "github.com/mattn/go-jsonpointer"
)

func main() {
  url := "https://raw.githubusercontent.com/stop-covid19-hyogo/covid19/development/data/main_summary.json"
  resp, _ := http.Get(url)
  defer resp.Body.Close()
  byteArray, _ := ioutil.ReadAll(resp.Body)

  var positives interface{}
  json.Unmarshal(byteArray, &positives)

  // 年月日
  d, _ := jsonpointer.Get(positives, "/last_update")

  // 軽症・中等症
  kei, _ := jsonpointer.Get(positives,
                            "/children/0/children/0/children/0/value")
  // 重症
  jyuu, _ := jsonpointer.Get(positives,
                              "/children/0/children/0/children/1/value")
  // 宿泊療養
  syuku, _ := jsonpointer.Get(positives, "/children/0/children/1/value")
  // 入院・宿泊療養調整等
  cyou, _ := jsonpointer.Get(positives, "/children/0/children/2/value")
  // 自宅療養
  home, _ := jsonpointer.Get(positives, "/children/0/children/3/value")
  // その他医療機関福祉施設等
  fukushi, _ := jsonpointer.Get(positives, "/children/0/children/4/value")
  // 死亡
  shibou, _ := jsonpointer.Get(positives, "/children/0/children/5/value")
  // 退院
  taiin, _ := jsonpointer.Get(positives, "/children/0/children/6/value")

  fmt.Printf("%v, %v, %v, %v, %v, %v, %v,, %v, %v\n", d, syuku, kei, cyou, home, fukushi, jyuu, shibou, taiin)

}

