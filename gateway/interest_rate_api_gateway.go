package gateway

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/spf13/viper"
)

type SelicData struct {
	Data  string `json:"data"`
	Valor string `json:"valor"`
}

func PrintSelic() {
	apiURL := viper.GetString("api.url")
	resp, err := http.Get(apiURL)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var selicData []SelicData
	err = json.Unmarshal(body, &selicData)
	if err != nil {
		panic(err)
	}

	// Calcular a média das taxas diárias do mês atual
	today := time.Now()
	currentMonth := today.Month()
	totalRate := 0.0
	countRate := 0

	for _, data := range selicData {
		dateObj, err := time.Parse("2006-01-02", data.Data)
		if err != nil {
			panic(err)
		}

		if dateObj.Month() == currentMonth {
			rate, err := strconv.ParseFloat(data.Valor, 64)
			if err != nil {
				panic(err)
			}

			totalRate += rate
			countRate++
		}
	}

	averageRate := totalRate / float64(countRate)
	fmt.Printf("A taxa SELIC média do mês atual é: %.2f\n", averageRate)
}
