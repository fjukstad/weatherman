package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type WeatherData struct {
	Location     Location     `xml:"location"`
	Credit       Credit       `xml:"credit"`
	Links        []Link       `xml:"links"`
	Meta         Meta         `xml:"meta"`
	Sun          Sun          `xml:"sun"`
	Forecasts    Forecasts    `xml:"forecast"`
	Observations Observations `xml:"observations"`
}

type Credit struct {
	Link Link `xml:"link"`
}

type Link struct {
	Text string `xml:"text,attr"`
	Url  string `xml:"url,attr"`
}

type Location struct {
	Name       string
	Type       string
	Country    string
	Timezone   string     `xml:"id,attr"`
	AltLongLat AltLongLat `xml:"location"`
}

type AltLongLat struct {
	Alt  string `xml:"altitude,attr"`
	Lat  string `xml:"longitude,attr"`
	Long string `xml:"latitude,attr"`

	geobase   string `xml:"geobase,attr"`
	geobaseId string `xml:"geobaseid,attr"`
}

type Meta struct {
	LastUpdate string
	NextUpdate string
}

type Sun struct {
	Rise string `xml:"rise,attr"`
	Set  string `xml:"set,attr"`
}

type Forecasts struct {
	Text    Text    `xml:"text"`
	Tabular Tabular `xml:"tabular"`
}

type Text struct {
	Location ForecastLocation `xml:"location"`
}

type ForecastLocation struct {
	Forecasts []Forecast `xml:"time"`
}

type Forecast struct {
	From          string        `xml:"from,attr"`
	To            string        `xml:"to,attr"`
	Day           string        `xml:"title"`
	Forecast      string        `xml:"body"`
	Symbol        Symbol        `xml:"symbol"`
	Percipitation Percipitation `xml:"precipitation"`
	WindDirection WindDirection `xml:"windDirection"`
	WindSpeed     WindSpeed     `xml:"windSpeed"`
	Temperature   Temperature   `xml:"temperature"`
	Pressure      Pressure      `xml:"pressure"`
}

type Symbol struct {
	Number   int    `xml:"number,attr"`
	NumberEx int    `xml:"numberEx,attr"`
	Name     string `xml:"name,attr"`
	Var      string `xml:"var,attr"`
}

type Percipitation struct {
	Value float64 `xml:"value,attr"`
	Max   float64 `xml:"maxvalue,attr"`
	Min   float64 `xml:"minvalue,attr"`
}

type WindDirection struct {
	Deg  string `xml:"deg,attr,float64"`
	Code string `xml:"code,attr"`
	Name string `xml:"name,attr"`
}
type WindSpeed struct {
	Mps  float64 `xml:"mps, attr"`
	Name string  `xml:"name"`
}
type Temperature struct {
	Unit  string  `xml:"unit, attr"`
	Value float64 `xml:"value,attr"`
}
type Pressure struct {
	Unit  string  `xml:"unit, attr"`
	value float64 `xml:"value,attr"`
}

type Tabular struct {
	Forecasts []Forecast `xml:"time"`
}

type Observations struct {
}

var ukedager = [...]string{
	"Søndag ",
	"Mandag",
	"Tirsdag",
	"Onsdag",
	"Torsdag",
	"Fredag",
	"Lørdag",
}

func main() {

	var text = flag.Bool("text", false, "text forecast. default false")
	flag.Parse()

	resp, err := http.Get("http://www.yr.no/sted/Norge/Troms/Troms%C3%B8/Troms%C3%B8/varsel.xml")

	if err != nil {
		fmt.Println("Could not fetch data from yr.no")
		return
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Could not read body ", err)
		return
	}

	wd := WeatherData{}

	err = xml.Unmarshal(body, &wd)

	if *text {
		wd.PrintText()
	} else {
		wd.PrintTabular()
	}

}

func (wd WeatherData) PrintTabular() error {
	tab := wd.Forecasts.Tabular
	var forrigeDag = ""

	for _, fc := range tab.Forecasts {
		from, err := time.Parse("2006-01-02T15:04:05", fc.From)
		if err != nil {
			fmt.Println("Could not parse time", err)
			return err
		}

		to, err := time.Parse("2006-01-02T15:04:05", fc.To)
		if err != nil {
			fmt.Println("Could not parse time", err)
			return err
		}

		dag := ukedager[from.Weekday()]

		if forrigeDag == "" || forrigeDag != dag {
			forrigeDag = dag
			fmt.Println(dag, from.Format("02.01.2006"))
		}

		fmt.Print("\t")
		fmt.Print(from.Format("15:04"), "-", to.Format("15:04"), " ")
		fmt.Print(fc.Symbol.Name)
		fmt.Print(" og ")
		fmt.Print(fc.Temperature.Value)
		fmt.Print(" grader. ")

		fmt.Print(fc.Percipitation.Value, " mm nedbør")

		fmt.Print("\n")
	}
	return nil
}

func (wd WeatherData) PrintText() error {
	forecasts := wd.Forecasts.Text.Location.Forecasts
	for _, fc := range forecasts {

		from, err := time.Parse("2006-01-02", fc.From)
		if err != nil {
			fmt.Println("Could not parse time", err)
			return err
		}

		to, err := time.Parse("2006-01-02", fc.To)
		if err != nil {
			fmt.Println("Could not parse time", err)
			return err
		}

		fromDay := ukedager[from.Weekday()]
		toDay := ukedager[to.Weekday()]

		if fromDay != toDay {
			fmt.Println(fromDay, "til", toDay)
		} else {
			fmt.Println(fromDay)
		}

		txt := strings.Split(fc.Forecast, "</strong>")
		forecast := txt[1]

		fmt.Print("\t")
		fmt.Println(forecast)

	}
	return nil

}
