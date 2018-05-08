package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func main() {

	var text = flag.Bool("text", false, "text forecast. default false")
	var location = flag.String("location", "/Norge/Troms/Tromsø/Tromsø", "the location you want the weather forecast from")
	var language = flag.String("language", "no", "language you want the forecast in (supports 'no' and 'en'")
	flag.Parse()

	resp, err := http.Get("http://www.yr.no/" + place[*language] + "/" + *location + "/varsel.xml")

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
		wd.PrintText(*language)
	} else {
		wd.PrintTabular(*language)
	}
}

func getWeekdays(language string) [7]string {
	if language == "no" {
		return ukedager
	} else {
		return weekdays
	}
}

func (wd WeatherData) PrintTabular(language string) error {
	tab := wd.Forecasts.Tabular

	var forrigeDag = ""

	days := getWeekdays(language)

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

		dag := days[from.Weekday()]

		if forrigeDag == "" || forrigeDag != dag {
			forrigeDag = dag
			fmt.Println(dag, from.Format("02.01.2006"))
		}

		windSpeed := strconv.FormatFloat(fc.WindSpeed.Mps, 'f', -1, 64)

		fmt.Print("\t")
		fmt.Print(from.Format("15:04"), "-", to.Format("15:04"), " ")
		fmt.Print(fc.Symbol.Name)
		fmt.Print(" " + and[language] + " ")
		fmt.Print(fc.Temperature.Value)
		fmt.Print(" " + degrees[language] + " ")

		fmt.Print(fc.Percipitation.Value, " mm "+precipitation[language])

		fmt.Print(". " + windSpeed + " m/s " + wind[language] + " ")
		fmt.Print(frm[language] + " " + fc.WindDirection.Name + ".")

		fmt.Print("\n")
	}
	return nil
}

func (wd WeatherData) PrintText(language string) error {
	forecasts := wd.Forecasts.Text.Location.Forecasts

	days := getWeekdays(language)

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

		fromDay := days[from.Weekday()]
		toDay := days[to.Weekday()]

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
