package pinUtils

import (
 "github.com/gocolly/colly"
 "encoding/json"
 "html"
 "errors"
 "example/hello/utilities/mapUtils"
 "example/hello/pinTypes"
 "example/hello/scrappingConstants"
)

func GetPinDataFromScrap(htmlString string) pinTypes.Pin { 

  baseString := "https://pinterest.com/pin/"
  var m interface{}
  json.Unmarshal([]byte(html.UnescapeString(htmlString)), &m)
  r := pinTypes.Pin {Name: "", Id: "",ReferUrl: "",OriginalImage: ""}
  r.Name , _ =  mapInterface.GetValueFromInterfaceMap(m.(map[string]interface{}), relationSearchQueries.Name)
  r.Id , _ =  mapInterface.GetValueFromInterfaceMap(m.(map[string]interface{}), relationSearchQueries.Id)
  r.OriginalImage , _ =  mapInterface.GetValueFromInterfaceMap(m.(map[string]interface{}), relationSearchQueries.OriginalImage)
  r.ReferUrl = baseString + r.Id
  return r
}

func PinFromUrl(url string) (pinTypes.Pin, error, int) {
    var pin pinTypes.Pin 
    var responseCode int
    var errorValue = errors.New("Error, no data to be scrapped on this link")
    c := colly.NewCollector()
    c.OnResponse(func(r *colly.Response) {
        responseCode = r.StatusCode
    })

    c.OnHTML("script", func(e *colly.HTMLElement) {
	if e.Attr("data-relay-response") == "true" {
	  innerHTML, _ := e.DOM.Html() //TODO get datapinfromscrap
	  pin = GetPinDataFromScrap(innerHTML)
	  errorValue = nil
	  return 
	}
    })
    c.Visit(url)
    return pin , errorValue , responseCode
}
