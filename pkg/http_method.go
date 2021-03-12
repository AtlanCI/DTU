package pkg

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func HttpGET(urlString string, urlParams url.Values, sTimeout int, header http.Header) (*http.Response, []byte, error) {
	//startTime := time.Now().UnixNano()
	client := http.Client{
		Timeout: time.Duration(sTimeout) * time.Second,
	}
	urlString = AddGetDataToUrl(urlString, urlParams)
	req, err := http.NewRequest("GET", urlString, nil)
	if err != nil {
		//Log.TagWarn(DLTagHTTPFailed, map[string]interface{}{
		//	"proc_time": float32(time.Now().UnixNano()-startTime) / 1.0e9,
		//	"args":      urlParams,
		//	"err":       err.Error(),
		//})
		//return nil, nil, err
	}
	if len(header) > 0 {
		req.Header = header
	}
	resp, err := client.Do(req)
	if err != nil {
		//Log.TagWarn(DLTagHTTPFailed, map[string]interface{}{
		//	"url":       urlString,
		//	"proc_time": float32(time.Now().UnixNano()-startTime) / 1.0e9,
		//	"method":    "GET",
		//	"args":      urlParams,
		//	"err":       err.Error(),
		//})
		//return nil, nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		//Log.TagWarn(DLTagHTTPFailed, map[string]interface{}{
		//	"url":       urlString,
		//	"proc_time": float32(time.Now().UnixNano()-startTime) / 1.0e9,
		//	"method":    "GET",
		//	"args":      urlParams, "result": Substr(string(body), 0, 1024),
		//	"err": err.Error(),
		//})
		//return nil, nil, err
	}
	//Log.TagInfo(DLTagHTTPSuccess, map[string]interface{}{
	//	"url":       urlString,
	//	"proc_time": float32(time.Now().UnixNano()-startTime) / 1.0e9,
	//	"method":    "GET",
	//	"args":      urlParams,
	//	"result":    Substr(string(body), 0, 1024),
	//})
	return resp, body, nil
}

func HttpPOST(urlString string, urlParams url.Values, sTimeout int, header http.Header, contextType string) (*http.Response, []byte, error) {
	//startTime := time.Now().UnixNano()
	client := http.Client{
		Timeout: time.Duration(sTimeout) * time.Second,
	}
	if contextType == "" {
		contextType = "application/x-www-form-urlencoded"
	}
	urlParamEncode := urlParams.Encode()
	req, err := http.NewRequest("POST", urlString, strings.NewReader(urlParamEncode))
	if len(header) > 0 {
		req.Header = header
	}
	req.Header.Set("Content-Type", contextType)
	resp, err := client.Do(req)
	if err != nil {
		//Log.TagWarn(DLTagHTTPFailed, map[string]interface{}{
		//             "url":       urlString,
		//             "proc_time": float32(time.Now().UnixNano()-startTime) / 1.0e9,
		//         "method":    "POST",
		//            "args":      Substr(urlParamEncode, 0, 1024),
		//            "err":       err.Error(),
		//        })
		//    return nil, nil, err
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// Log.TagWarn(DLTagHTTPFailed, map[string]interface{}{
		//         "url":       urlString,
		//        "proc_time": float32(time.Now().UnixNano()-startTime) / 1.0e9,
		//         "method":    "POST",
		//         "args":      Substr(urlParamEncode, 0, 1024),
		//         "result":    Substr(string(body), 0, 1024),
		//         "err":       err.Error(),
		//     })
		//return nil, nil, err
		panic(err)
	}
	//Log.TagInfo(DLTagHTTPSuccess, map[string]interface{}{
	//        "url":       urlString,
	//        "proc_time": float32(time.Now().UnixNano()-startTime) / 1.0e9,
	//        "method":    "POST",
	//        "args":      Substr(urlParamEncode, 0, 1024),
	//       "result":    Substr(string(body), 0, 1024),
	//    })
	return resp, body, nil
}
func AddGetDataToUrl(urlString string, data url.Values) string {
	if strings.Contains(urlString, "?") {
		urlString = urlString + "&"
	} else {
		urlString = urlString + "?"
	}
	return fmt.Sprintf("%s%s", urlString, data.Encode())
}

func HttpJSON(urlString string, jsonContent string, sTimeout int, header http.Header) (*http.Response, []byte, error) {
	//startTime := time.Now().UnixNano()
	client := http.Client{
		Timeout: time.Duration(sTimeout) * time.Second,
	}
	req, err := http.NewRequest("POST", urlString, strings.NewReader(jsonContent))
	if len(header) > 0 {
		req.Header = header
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		//Log.TagWarn(DLTagHTTPFailed, map[string]interface{}{
		//	"url":       urlString,
		//	"proc_time": float32(time.Now().UnixNano()-startTime) / 1.0e9,
		//	"method":    "POST",
		//	"args":      Substr(jsonContent, 0, 1024),
		//	"err":       err.Error(),
		//})
		//return nil, nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		//Log.TagWarn(DLTagHTTPFailed, map[string]interface{}{
		//	"url":       urlString,
		//	"proc_time": float32(time.Now().UnixNano()-startTime) / 1.0e9,
		//	"method":    "POST",
		//	"args":      Substr(jsonContent, 0, 1024),
		//	"result":    Substr(string(body), 0, 1024),
		//	"err":       err.Error(),
		//})
		//return nil, nil, err
	}
	//Log.TagInfo(DLTagHTTPSuccess, map[string]interface{}{
	//	"url":       urlString,
	//	"proc_time": float32(time.Now().UnixNano()-startTime) / 1.0e9,
	//	"method":    "POST",
	//	"args":      Substr(jsonContent, 0, 1024),
	//	"result":    Substr(string(body), 0, 1024),
	//})
	return resp, body, nil
}
