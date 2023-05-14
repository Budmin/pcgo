package pcgo

import (
	"log"
	"net/http"
	"strconv"
	"time"
)

// Makes a query
func Query(
	req *http.Request,
) (*http.Response, error) {

	client := &http.Client{}

	res, err := client.Do(req)

	if err != nil {
		log.Println("Error fulfilling http request")
		log.Fatalln(err)
	}

	if res.StatusCode == 429 {

		retryAfterStr := res.Header.Get("Retry-After")

		intVar, err := strconv.Atoi(retryAfterStr)

		if err != nil {
			log.Println("Error parsing Retry-After. Setting wait time to fail-safe value of 20")
			log.Println(err)

			// 19 here, then the safety second is added to total 20
			intVar = 19
		}

		// add an extra second for safety
		intVar++

		// fmt.Printf("429 error, waiting %d seconds\n", intVar)
		time.Sleep(time.Second * time.Duration(intVar))

		// recursive process has a possibility to loop infinitely
		// in extreme cases this could be problematic (multiple 429 error loops leading to extra ram usage)
		// should be fine in most normal use cases, but still something to fix later
		return Query(
			req,
		)

	}

	return res, err
}
