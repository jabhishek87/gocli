/*
Copyright © 2021 Abhishek Jaiswal <abhishekjaiswal.kol@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

// jokeCmd represents the joke command
var jokeCmd = &cobra.Command{
	Use:   "joke",
	Short: "command to fetch joke from dad api",
	Long: `This command fetches a random dad joke from the
	https://icanhazdadjoke.com/ api's
	`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("joke called")
		getRandomJoke()
	},
}

func init() {
	rootCmd.AddCommand(jokeCmd)
}

type Joke struct {
	ID     string `json:"id"`
	Joke   string `json:"joke"`
	Status int    `json:"status"`
}

func getRandomJoke() {
	url := "https://icanhazdadjoke.com/"
	resBytes := getJokeAPI(url)

	joke := Joke{}
	if err := json.Unmarshal(resBytes, &joke); err != nil {
		log.Printf("could not unmarshall responsee - %v", err)
	}
	fmt.Println(string(joke.Joke))
}

func getJokeAPI(baseAPI string) []byte {
	req, err := http.NewRequest(
		http.MethodGet, baseAPI, nil,
	)
	if err != nil {
		log.Printf("could not request a dad joke - %v", err)
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("User-agent", "gocli")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Printf("could notmake a request - %v", err)
	}

	resBytes, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Printf("could not read reponse - %v", err)
	}

	return resBytes
}
