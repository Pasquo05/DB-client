package funcclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type client struct {
	url string
	cl  *http.Client
}

func New(url string) client {

	c := client{
		url: url,
		cl:  &http.Client{},
	}
	return c
}

func (c client) PostRequest(path string, numberInput funcdbserver.PhoneNumber /* meglio interface{}*/) {

	number, err := json.Marshal(numberInput)

	if err != nil {
		log.Fatal(err)
		return
	}

	header := fmt.Sprintf("%s%s", c.url, path)

	resp, err := http.Post(header, "application/json", bytes.NewBuffer(number))
	if err != nil {
		fmt.Println("Errore durante l'esecuzione della richiesta:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Errore durante la lettura della risposta:", err)
		return
	}

	fmt.Println(string(body))

}

func (c client) Delete(path string, id string) {
	header := fmt.Sprintf("%s%s%s", c.url, path, id)
	fmt.Println(header)
	req, err := http.NewRequest("DELETE", header, nil)
	if err != nil {
		log.Fatal(err)
	}

	c.cl.Do(req)
}

func (c client) GetRequest(path string) {

	resp, err := http.Get(fmt.Sprintf("%v%v", c.url, path))
	if err != nil {
		fmt.Println("Errore durante l'esecuzione della richiesta:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Errore durante la lettura della risposta:", err)
		return
	}

	fmt.Println(string(body))
}
