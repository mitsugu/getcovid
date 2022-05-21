package main

import (
  "encoding/csv"
  "net/http"
  "strings"
  "fmt"
)

func readCSVFromUrl(url string) ([][]string, error) {
  resp, err := http.Get(url)
  if err != nil {
    return nil, err
  }

  defer resp.Body.Close()
  reader := csv.NewReader(resp.Body)
  reader.Comma = ';'
  rows, err := reader.ReadAll()
  if err != nil {
    return nil, err
  }

  return rows, nil
}

func getIrc() (string, string, string) {
  ircUrl := "https://covid19.mhlw.go.jp/public/opendata/requiring_inpatient_care_etc_daily.csv"
	ircdata, err := readCSVFromUrl(ircUrl)
	if err != nil {
		panic(err)
	}

  strRow := ircdata[len(ircdata)-1][0]
  slice := strings.Split(strRow, ",")
  return slice[0], slice[1], slice[2]
}

func getSrv() (string) {
  srvUrl := "https://covid19.mhlw.go.jp/public/opendata/severe_cases_daily.csv"
	srvdata, err := readCSVFromUrl(srvUrl)
	if err != nil {
		panic(err)
	}

  strRow := srvdata[len(srvdata)-1][0]
  slice := strings.Split(strRow, ",")
  return slice[1]
}

func getDeath() (string) {
  deathUrl := "https://covid19.mhlw.go.jp/public/opendata/number_of_deaths_daily.csv"
	deathdata, err := readCSVFromUrl(deathUrl)
	if err != nil {
		panic(err)
	}

  strRow := deathdata[len(deathdata)-1][0]
  slice := strings.Split(strRow, ",")
  return slice[1]
}

func getDeathAccumulation() (string) {
  accUrl := "https://covid19.mhlw.go.jp/public/opendata/deaths_cumulative_daily.csv"
	accdata, err := readCSVFromUrl(accUrl)
	if err != nil {
		panic(err)
	}

  strRow := accdata[len(accdata)-1][0]
  slice := strings.Split(strRow, ",")
  return slice[1]
}

func main() {
  d, irc, released := getIrc()  // 日付, 要入院者数, 退院数
  srv := getSrv()               // 重症者数
  death := getDeath()           // 死亡者数
  acc := getDeathAccumulation() // 死亡者数累計

  fmt.Printf("%v, %v, %v, %v, %v, %v\n", d, irc, srv, death, acc, released )
}

