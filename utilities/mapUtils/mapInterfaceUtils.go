package mapInterface

import "errors"
func GetValueFromInterfaceMap(mapToParse map[string]interface{}, valuePath []string) (string, error) {
	var stageInterface map[string]interface{} = mapToParse
	for _, v := range valuePath {
	  switch stageInterface[v].(type){
            case map[string]interface{}:
		    stageInterface = stageInterface[v].(map[string]interface{})
            case string:
                    return stageInterface[v].(string), nil
            default:
                    return "", errors.New("Error, interface element is not either a string or another map interface, or is not present")
	  }
        }
        return "", errors.New("Error, key finding exhausted without string match")
}
