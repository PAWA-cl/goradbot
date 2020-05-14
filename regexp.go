package goradbot

import(
	"regexp"
)

//RegexMap is a map for the regular expression method
type RegexMap struct {
	MappedResult map[string]string
}

//getParams is the best way to obtain the groups for regex
//kudos to eluleci: https://stackoverflow.com/a/39635221/3700649
func getParams(regEx, str string) (paramsMap map[string]string) {

    var compRegEx = regexp.MustCompile(regEx)
    match := compRegEx.FindStringSubmatch(str)

    paramsMap = make(map[string]string)
    for i, name := range compRegEx.SubexpNames() {
        if i > 0 && i <= len(match) {
            paramsMap[name] = match[i]
        }
    }
    return
}

//Matches checks if expression match any substring and saves the mapped result.
func(r *RegexMap)Matches(regEx, str string)(bool){
	r.MappedResult = getParams(regEx, str)
	
	return len(r.MappedResult) > 0
}


