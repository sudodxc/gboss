package parse

import (
	"city/engine"
	"regexp"
)

const cityList = `<a ka=".+" href="(/.+/)">([^<]+)</a>`

//    <a href="/c101020100-p100103/?page=1" ka="page-prev" class="prev"></a>
func ParseCityList(contents []byte) engine.ParseResult {

	re := regexp.MustCompile(cityList)
	match := re.FindAllSubmatch(contents, -1)
	request := engine.ParseResult{}
	for _, m := range match {
		//fmt.Printf("%s", m[1]) //wangzhi
		//name := string(m[2])
		// profile := model.Profile{}
		// profile.Name = name
		// profile.Url = "https://www.zhipin.com" + string(m[1])
		// request.Items = append(request.Items, profile) //job
		request.Requests = append(request.Requests, engine.Request{
			Url: "https://www.zhipin.com" + string(m[1]), //url
			ParseFunc: func(c []byte) engine.ParseResult { //job
				return ParseJob(c)
			}, //job
			//                         //joblist
		})
		//fmt.Printf("job:%s ,URL:https://www.zhipin.com%s\n", m[2], m[1])
	}
	return request
}
