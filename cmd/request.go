package cmd

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"

    "github.com/spf13/cobra"
)

type response2 struct {
    Height     string
    Mass       string
    Hair_color string
    Name       string
    Films      []string
}

var flat bool
var id int

func init() {
    rootCmd.AddCommand(requestCmd)
    rootCmd.PersistentFlags().BoolVarP(&flat, "verbose", "v", false, "print output (verbosely)")
    rootCmd.PersistentFlags().IntVarP(&id, "id", "i", 1, "ID of person to get")
}

var requestCmd = &cobra.Command{
    Use:   "request",
    Short: "Print the request number of go-cli",
    Long:  `Make a request to SWAPI`,
    Run: func(cmd *cobra.Command, args []string) {
        var url = fmt.Sprintf("https://swapi.co/api/people/%d/", id)

        fmt.Println("Starting the request...")
        response, err := http.Get(url)
        if err != nil {
            fmt.Printf("The HTTP request failed with error %s\n", err)
        } else {
            data, _ := ioutil.ReadAll(response.Body)

            res := response2{}
            json.Unmarshal([]byte(string(data)), &res)

            if flat {
                fmt.Println(res)
                fmt.Println(res.Films[0])
            }

        }
        // jsonData := map[string]string{"firstname": "Nic", "lastname": "Raboy"}
        // jsonValue, _ := json.Marshal(jsonData)
        // response, err = http.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer(jsonValue))
        // if err != nil {
        //     fmt.Printf("The HTTP request failed with error %s\n", err)
        // } else {
        //     data, _ := ioutil.ReadAll(response.Body)
        //     fmt.Println(string(data))
        // }
        fmt.Println("Terminating requests...")

    },
}