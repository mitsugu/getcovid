package main

import (
  "encoding/json"
  "net/http"
  "fmt"
  "io/ioutil"
)

func main() {
  url := "https://raw.githubusercontent.com/stop-covid19-hyogo/covid19/development/data/main_summary.json"
  resp, _ := http.Get(url)
  defer resp.Body.Close()
  byteArray, _ := ioutil.ReadAll(resp.Body)

  var positives interface{}
  json.Unmarshal(byteArray, &positives)
  // 年月日
  d := positives.(map[string]interface{})["last_update"].(string);
  // 軽症・中等症
  kei := positives.(map[string]interface{})["children"].([]interface {})[0].(map[string]interface {})["children"].([]interface {})[0].(map[string]interface {})["children"].([]interface {})[0].(map[string]interface{})["value"]
  // 重症
  jyuu := positives.(map[string]interface{})["children"].([]interface {})[0].(map[string]interface {})["children"].([]interface {})[0].(map[string]interface {})["children"].([]interface {})[1].(map[string]interface{})["value"]
  // 宿泊療養
  syuku := positives.(map[string]interface{})["children"].([]interface {})[0].(map[string]interface {})["children"].([]interface {})[1].(map[string]interface{})["value"]
  // 入院・宿泊療養調整等
  cyou := positives.(map[string]interface{})["children"].([]interface {})[0].(map[string]interface {})["children"].([]interface {})[2].(map[string]interface {})["value"]
  // 自宅療養
  home := positives.(map[string]interface{})["children"].([]interface {})[0].(map[string]interface {})["children"].([]interface {})[3].(map[string]interface {})["value"]
  // その他医療機関福祉施設等
  fukushi := positives.(map[string]interface{})["children"].([]interface {})[0].(map[string]interface {})["children"].([]interface {})[4].(map[string]interface {})["value"]
  // 死亡
  shibou := positives.(map[string]interface{})["children"].([]interface {})[0].(map[string]interface {})["children"].([]interface {})[5].(map[string]interface {})["value"]
  // 退院
  taiin := positives.(map[string]interface{})["children"].([]interface {})[0].(map[string]interface {})["children"].([]interface {})[6].(map[string]interface {})["value"]

  fmt.Printf("%v, %v, %v, %v, %v, %v, %v, %v, %v\n", d, syuku, kei, cyou, home, fukushi, jyuu, shibou, taiin)

}


/*
map[string]interface {}{
  "attr":"検査実施人数",
  "children":[]interface {}{
    map[string]interface {}{
      "attr":"陽性患者数",
      "children":[]interface {}{
        map[string]interface {}{
          "attr":"入院中",
          "children":[]interface {}{
            map[string]interface {}{
              "attr":"軽症・中等症",
              "value":276
            },
            map[string]interface {}{
              "attr":"重症",
              "value":15
            }
          },
          "value":291
        },
        map[string]interface {}{
          "attr":"宿泊療養",
          "value":327
        },
        map[string]interface {}{
          "attr":"入院・宿泊療養調整等",
          "children":[]interface {}{
            map[string]interface {}{
              "attr":"入院調整",
              "value":40
            }
          },
          "value":64
        },
        map[string]interface {}{
          "attr":"自宅療養",
          "value":204
        },
        map[string]interface {}{
          "attr":"その他医療機関福祉施設等",
          "value":2
        },
        map[string]interface {}{
          "attr":"死亡",
          "value":1315
        },
        map[string]interface {}{
          "attr":"退院", "value":40322
        }
      },
      "value":42525
    }
  },
  "last_update":"2021-07-26T00:00:00+09:00",
  "value":544463
}
*/

