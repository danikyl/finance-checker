package integration

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/spf13/viper"
)

type SelicData struct {
	Date  string `json:"data"`
	Valor string `json:"valor"`
}

func PrintSelic() {
	viper.SetConfigFile("config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	apiURL := viper.GetString("api.url")
	formattedUrl := addInitialAndFinalDateVariablesToUrl(apiURL)

	resp, err := http.Get(formattedUrl)
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

	selicToday := selicData[len(selicData)-1]
	fmt.Printf("A taxa SELIC média do mês atual é: %v\n", selicToday.Valor)
}

func addInitialAndFinalDateVariablesToUrl(url string) string {
	now := time.Now()
	firstDayOfLastMonth := time.Date(now.Year(), now.Month()-1, 1, 0, 0, 0, 0, now.Location())
	initialDate := firstDayOfLastMonth.Format("02/01/2006")
	finalDate := now.Format("02/01/2006")
	formattedUrl := strings.Replace(url, "$initial_date", initialDate, 1)
	formattedUrl = strings.Replace(formattedUrl, "$final_date", finalDate, 1)
	return formattedUrl
}
