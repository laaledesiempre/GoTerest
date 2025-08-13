package main

import (
    "fmt"
    "github.com/gocolly/colly"
    "encoding/json"
 "html"
)

//type Pin struct {
//   Name string
//   Id string
//   ReferUrl string
//   OriginalImage string
//}
////type Download interface{
////	//TODO
////}
//
//type Idea struct {
//   Name string
//   Id string
//   ReferUrl string
//   Pins []Pin
//}
//
//func getPinDataFromScrap(htmlString string) Pin { //TODO error patern matching
//  var m interface{}
//  json.Unmarshal([]byte(html.UnescapeString(htmlString)), &m)
//  m.(map[string]interface{})["response"].(map[string]interface{})["data"].(map[string]interface{})["v3GetPinQuery"].(map[string]interface{})["data"].(map[string]interface{})["imageSpec_orig"].(map[string]interface{})["url"]
//}
//func pinFromUrl(url string) (Pin, error) {
//    var pin Pin
//    var responseCode uint16
//    c := colly.NewCollector()
//    c.OnResponse(func(r *colly.Response) {
//        responseCode = r.StatusCode
//    })
//
//    c.OnHTML("script", func(e *colly.HTMLElement) {
//	if e.Attr("data-relay-response") == "true" {
//	  innerHTML, _ := e.DOM.Html() //TODO get datapinfromscrap
//	  //fmt.Printf("%s\n",html.UnescapeString(innerHTML))
//	  fmt.Printf("%s\n",[]byte(html.UnescapeString(innerHTML)))
//	  json.Unmarshal([]byte(html.UnescapeString(innerHTML)), &m)
//	  fmt.Printf("%s\n","responsed")
//	  fmt.Println(m.(map[string]interface{})["response"].(map[string]interface{})["data"].(map[string]interface{})["v3GetPinQuery"].(map[string]interface{})["data"].(map[string]interface{})["imageSpec_orig"].(map[string]interface{})["url"])
//	}
//    })
//    c.Visit("https://ar.pinterest.com/pin/1046383294679022220/")
//
//}
func getValueFromInterfaceMap(mapToParse map[string]interface{}, valuePath []string) (string, error) {
	var stageInterface map[string]interface{} = mapToParse
	var result string
	for _, v := range valuePath {
	  switch stageInterface[v].(type){
            case map[string]interface{}:
		    stageInterface = stageInterface[v].(map[string]interface{})
            case string:
                    result = stageInterface[v].(string)
		    return result, nil
            default:
		    return "error", nil
	  }
        }
        return "error", nil
}
func main() {
 
    var m interface{}
    // instantiate a new collector object
    c := colly.NewCollector()

    c.OnResponse(func(r *colly.Response) {
        fmt.Println("Status:", r.StatusCode)
	//fmt.Printf("Response Body:\n%s\n", r.Body)
    })

    c.OnHTML("script", func(e *colly.HTMLElement) {
	if e.Attr("data-relay-response") == "true" {
	  innerHTML, _ := e.DOM.Html()
	  //fmt.Printf("%s\n",html.UnescapeString(innerHTML))
	  fmt.Printf("%s\n",[]byte(html.UnescapeString(innerHTML)))
	  json.Unmarshal([]byte(html.UnescapeString(innerHTML)), &m)
	  fmt.Printf("%s\n","responsed")
	  //fmt.Println(m.(map[string]interface{})["response"].(map[string]interface{})["data"].(map[string]interface{})["v3GetPinQuery"].(map[string]interface{})["data"].(map[string]interface{})["imageSpec_orig"].(map[string]interface{})["url"])
 fmt.Println(getValueFromInterfaceMap(m.(map[string]interface{}), []string{"response","data","v3GetPinQuery", "data", "imageSpec_orig","url"}))
	}
    })
    c.Visit("https://ar.pinterest.com/pin/1046383294679022220/")
}
