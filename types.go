package main

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
	Mps  float64 `xml:"mps,attr"`
	Name string  `xml:"name,attr"`
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

var weekdays = [...]string{
	"Sunday",
	"Monday",
	"Tueday",
	"Wednesday",
	"Thursday",
	"Saturday",
	"Sunday",
}

var place = map[string]string{
	"no": "sted",
	"en": "place",
}

var degrees = map[string]string{
	"no": "grader",
	"en": "degrees",
}

var precipitation = map[string]string{
	"no": "nedbør",
	"en": "precipitation",
}

var and = map[string]string{
	"no": "og",
	"en": "and",
}

var wind = map[string]string{
	"no": "vind",
	"en": "wind",
}

var frm = map[string]string{
	"no": "fra",
	"en": "from",
}
