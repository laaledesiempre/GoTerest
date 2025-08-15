package main

import (
    "fmt"
    "github.com/gocolly/colly"
    "encoding/json"
 "html"
 //"errors"
 "example/hello/mapUtilities"
 "example/hello/pinTypes"

)

func getPinDataFromScrap(htmlString string) { //TODO error patern matching //TODO not yet return type
  baseString := "https://pinterest.com/pin/"
  var m interface{}
  json.Unmarshal([]byte(html.UnescapeString(htmlString)), &m)
  r := pinTypes.Pin {Name: "", Id: "",ReferUrl: "",OriginalImage: ""}
  var relationSearchQueries = map[string][]string{}
  relationSearchQueries["Name"] = []string{"response","data","v3GetPinQuery", "data", "seoTitle"}
  relationSearchQueries["Id"] = []string{"variables","pinId"}
  relationSearchQueries["OriginalImage"] = []string{"response","data","v3GetPinQuery", "data", "imageSpec_orig","url"}

  r.Name , _ =  mapInterface.GetValueFromInterfaceMap(m.(map[string]interface{}), relationSearchQueries["Name"])
  r.Id , _ =  mapInterface.GetValueFromInterfaceMap(m.(map[string]interface{}), relationSearchQueries["Id"])
  r.OriginalImage , _ =  mapInterface.GetValueFromInterfaceMap(m.(map[string]interface{}), relationSearchQueries["OriginalImage"])
  r.ReferUrl = baseString + r.Id
  fmt.Printf("v%",r)
}

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
          getPinDataFromScrap(innerHTML)

	  //fmt.Printf("%s\n",html.UnescapeString(innerHTML))
	  fmt.Printf("%s\n",[]byte(html.UnescapeString(innerHTML)))
	  json.Unmarshal([]byte(html.UnescapeString(innerHTML)), &m)
	  fmt.Printf("%s\n","responsed")
	  //fmt.Println(m.(map[string]interface{})["response"].(map[string]interface{})["data"].(map[string]interface{})["v3GetPinQuery"].(map[string]interface{})["data"].(map[string]interface{})["imageSpec_orig"].(map[string]interface{})["url"])
 fmt.Println(mapInterface.GetValueFromInterfaceMap(m.(map[string]interface{}), []string{"response","data","v3GetPinQuery", "data", "imageSpec_orig","url"}))
 fmt.Println(mapInterface.GetValueFromInterfaceMap(m.(map[string]interface{}), []string{"response","data","v3GetPinQuery", "data", "imageSpec_orig","asdlkjasdlkjalsd"}))
 fmt.Println(mapInterface.GetValueFromInterfaceMap(m.(map[string]interface{}), []string{"response","data","v3GetPinQuery", "data", "imageSpec_orig"}))
	}
    })
    c.Visit("https://ar.pinterest.com/pin/1046383294679022220/")
}
