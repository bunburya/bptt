package waqi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"ptt/output"

	"github.com/fatih/color"
)

const baseUrl = "https://api.waqi.info"

func cityUrl(city string, apiKey string) string {
	return fmt.Sprintf("%s/feed/%s?token=%s", baseUrl, city, apiKey)
}

func latLonUrl(lat float64, lon float64, apiKey string) string {
	return fmt.Sprintf("%s/feed/geo:%f;%f/?token=%s", baseUrl, lat, lon, apiKey)
}

type aqi int

func (a *aqi) description() string {
	n := *a
	if n <= 50 {
		return "Good"
	} else if n <= 100 {
		return "Moderate"
	} else if n <= 150 {
		return "Unhealthy for Sensitive Groups"
	} else if n <= 200 {
		return "Unhealthy"
	} else if n <= 300 {
		return "Very Unhealthy"
	} else {
		return "Hazardous"
	}
}

func (a *aqi) color() *color.Color {
	n := *a
	if n <= 50 {
		return output.SafetyColors["green"]
	} else if n <= 150 {
		return output.SafetyColors["yellow"]
	} else {
		return output.SafetyColors["red"]
	}
}

func (a *aqi) toRow() output.Row {
	return output.NewRow(
		output.NewCell(fmt.Sprintf("%d", *a), nil),
		output.NewCell(a.description(), a.color().Add(color.Bold)),
	)
}

type rawWaqiResponse struct {
	Status  string          `json:"status"`
	RawData json.RawMessage `json:"data"`
}

type waqiData struct {
	Aqi aqi `json:"aqi"`
}

func parseResponse(body []byte) (aqi, error) {
	var r rawWaqiResponse
	if err := json.Unmarshal(body, &r); err != nil {
		return 0, err
	}

	if r.Status == "ok" {
		var data waqiData
		if err := json.Unmarshal(r.RawData, &data); err != nil {
			return 0, err
		}
		return data.Aqi, nil
	} else {
		var errMsg string
		if err := json.Unmarshal(r.RawData, &errMsg); err != nil {
			return 0, err
		}
		return 0, errors.New(errMsg)
	}
}

func request(url string) (aqi, error) {
	var a aqi
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return a, err
	}
	req.Header.Set("User-Agent", "ptt")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return a, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return a, errors.New(resp.Status)
	}
	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return a, err
	}
	return parseResponse(respBytes)
}

func getCityForecast(city string, apiKey string) (aqi, error) {
	return request(cityUrl(city, apiKey))
}

func getLatLonForecast(lat float64, lon float64, apiKey string) (aqi, error) {
	return request(latLonUrl(lat, lon, apiKey))
}

func CityAqiTable(city string, apiKey string, options output.Options) (output.Table, error) {
	t := output.Table{}
	f, err := getCityForecast(city, apiKey)
	if err != nil {
		return t, err
	}
	if options.Header {
		t.SetHeader(output.NewRow(
			output.NewCell("AQI", color.New(color.Bold)),
			output.NewCell("Description", color.New(color.Bold)),
		))
	}
	t.AddRow(f.toRow())
	if options.Timestamp {
		t.Timestamp()
	}
	return t, nil
}
