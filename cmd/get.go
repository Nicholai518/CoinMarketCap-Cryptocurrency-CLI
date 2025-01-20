/*
Copyright © 2025 Nicholas Pazienza
*/
package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

type Cryptocurrency struct {
	Name     string  `json:"name"`
	Symbol   string  `json:"symbol"`
	Price    float64 `json:"price"`
	Quantity float64 `json:"quantity"`
	Value    float64 `json:"value"`
}

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Reads user input and sends GET request to CoinMarketCap.",
	Long:  `Application sends GET request to CoinMarketCap. Displays displays Cryptocurrency Name, Symbol, USD price, Quantity and Value.`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		quantity, _ := cmd.Flags().GetFloat64("quantity")

		cryptocurrency, err := getCryptocurrencyData(name, quantity)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println()
		time := time.Now().Format("2006-01-02 15:04:05")
		fmt.Println("Date & Time:", time)
		fmt.Println("Cryptocurrency Name:", cryptocurrency.Name)
		fmt.Println("Symbol:", cryptocurrency.Symbol)
		fmt.Println("Price:", fmt.Sprintf("$%.2f", cryptocurrency.Price))
		fmt.Println("Quantity:", cryptocurrency.Quantity)
		fmt.Println("Value:", fmt.Sprintf("$%.2f", cryptocurrency.Value))
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.PersistentFlags().String("name", "", "Cryptocurrency name.")
	getCmd.PersistentFlags().Float64("quantity", 0, "Quantity owned for Cryptocurrency name.")
}

func getCryptocurrencyData(name string, quantity float64) (Cryptocurrency, error) {
	client := &http.Client{}
	apiKey := os.Getenv("API_KEY")
	cryptocurrency := Cryptocurrency{}

	// crypto name / slug has to be all lowercase to make successful GET request
	cryptoName := strings.ToLower(name)

	if apiKey == "" {
		fmt.Println("API_KEY not set!. Please get your API key from CoinMarketCap and try again. Now Exiting.")
		os.Exit(1)
	}

	req, err := http.NewRequest("GET", "https://pro-api.coinmarketcap.com/v2/cryptocurrency/quotes/latest", nil)
	if err != nil {
		log.Print("GET request failed: ", err)
		os.Exit(1)
	}

	q := url.Values{}
	q.Add("slug", cryptoName)

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", apiKey)
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request to server. Now Exiting.")
		os.Exit(1)
	}

	// name is valid - return crypto
	if resp.StatusCode == 200 {

		respBody, _ := ioutil.ReadAll(resp.Body)

		// Unmarshall JSON into a map[string]interface{}
		var data map[string]interface{}
		err = json.Unmarshal([]byte(respBody), &data)
		if err != nil {
			panic(err)
		}

		// JSON is heavily nested
		// data > id > name, symbol
		dataSection := data["data"].(map[string]interface{})

		// every crypto has a different id. Cannot just use "id"
		// Unable to determine how CoinMarketCap is assigning id's. Not numerical order. Appears random
		// iterate once and stop at first key which is the id
		var id string
		for key, _ := range dataSection {
			id = key
			break
		}
		idForCryptoName := dataSection[id].(map[string]interface{})
		cryptocurrency.Name = idForCryptoName["name"].(string)
		cryptocurrency.Symbol = idForCryptoName["symbol"].(string)

		// data > id > quote > USD > price
		quote := idForCryptoName["quote"].(map[string]interface{})
		usd := quote["USD"].(map[string]interface{})
		cryptocurrency.Price = usd["price"].(float64)

		cryptocurrency.Quantity = quantity
		cryptocurrency.Value = cryptocurrency.Price * cryptocurrency.Quantity
		return cryptocurrency, nil
	} else {
		// User either misspelled crypto name or crypto does not exist on CoinMarketCap
		fmt.Println()
		fmt.Println("WARNING: Double check spelling and make sure this Cryptocurrency is available on CoinMarketCap.")
		fmt.Println()
	}

	// if you make it to this point, there was a problem. Return an error
	return Cryptocurrency{Name: name}, errors.New("unable to find Cryptocurrency on CoinMarketCap")
}
