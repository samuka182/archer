package archttp

import (
	"archer/structs"
	"bytes"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
	"time"
)

//Request ...
func Request(request *structs.Request) (r *structs.Response, err error) {

	url := getURL(request)
	client := &http.Client{}
	response := &structs.Response{}

	startTime := time.Now()
	req, err := http.NewRequest(request.Method, url, bytes.NewBufferString(request.Payload))
	setHEADERS(request, req)

	resp, err := client.Do(req)

	response.ExecutionTime = fmt.Sprintf("%s", time.Since(startTime))

	dispatch(resp, response)

	return response, err
}

func getURL(request *structs.Request) (url string) {
	url = request.URL

	for i := 0; i < len(request.Path); i++ {
		path := request.Path[i]
		re := regexp.MustCompile("{" + path.Key + "}")
		url = re.ReplaceAllString(url, path.Value)
	}

	return url
}

func setHEADERS(request *structs.Request, req *http.Request) {

	for i := 0; i < len(request.Header); i++ {
		header := request.Header[i]
		req.Header.Set(header.Key, header.Value)
	}

}

func dispatch(resp *http.Response, output *structs.Response) {

	buf := new(bytes.Buffer)
	io.Copy(buf, resp.Body)
	resp.Body.Close()
	payload := strings.TrimSpace(buf.String())

	output.HTTPStatus = resp.Status
	output.HTTPStatusCode = resp.StatusCode
	output.Payload = payload

	for k, v := range resp.Header {
		header := structs.Header{}
		header.Chave = k
		header.Valor = v[0]
		output.Headers = append(output.Headers, header)
	}

}
