package curl

import (
	"io/ioutil"
	"net/http"
	"net/url"
)


func Get(url string , parms url.Values) string{
	resp, err := http.Get(url+ "?"+parms.Encode())
	if err != nil {
		// handle error
		return ""
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		return ""
	}

	return string(body)
	//fmt.Println(string(body))
	//fmt.Printf("%s",body)
}

func Post(url string,parms url.Values) string{
	resp, err := http.PostForm(url,parms)

	if err != nil {
		// handle error
		return ""
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		return ""
	}
	return string(body)
	//fmt.Println(string(body))
	//fmt.Printf("%s",body)
}


