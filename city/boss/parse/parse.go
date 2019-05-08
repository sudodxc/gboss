package parse

import (
	"city/engine"
	"model"
	"regexp"
)

//var cityre = regexp.MustCompile(`<div class="text">([\w\W]*?)</div>`)

// func ParseJobs(contents []byte, name string) engine.ParseResult {
// 	profile := model.Profile{}
// 	profile.Name = name
// 	// profile.Duty = exString(contents, dutyre)
// 	//fmt.Sprintf("duty id %s",profile.Duty)
// 	profile.CityYear = exString(contents, cityre)
// 	//	profile.Url = exString(contents, urlre)
// 	// profile.PostSalary = exString(contents, postre)
// 	// profile.Status = exString(contents, statusre)
// 	// profile.Welfare = exString(contents, welfarere)

// 	result := engine.ParseResult{
// 		Items: []interface{}{profile},
// 	}
// 	return result

// }

// func exString(contents []byte, re *regexp.Regexp) string {
// 	match := re.FindSubmatch(contents)

// 	return string(match[1])

// }

const cityLists = `<div class="text">([\w\W]*?)</div>`

func ParseJobs(contents []byte, name string, url string, postSalary string) engine.ParseResult {
	re := regexp.MustCompile(cityLists)
	match := re.FindAllSubmatch(contents, -1)

	request := engine.ParseResult{}
	for _, m := range match {
		profile := model.Profile{}
		profile.Duty = string(m[1])
		profile.Url = url
		profile.PostSalary = postSalary
		profile.Name = name
		request.Items = append(request.Items, profile) //job

		// request.Requests = append(request.Requests, engine.Request{
		// 	Url: "https://www.zhipin.com" + string(m[1]), //url
		// 	// ParseFunc: func(c []byte) engine.ParseResult { //job
		// 	// 	return ParseJob(c)
		// 	// }, //job
		// 	//                         //joblist
		// })
		//fmt.Printf("job:%s ,URL:https://www.zhipin.com%s\n", m[2], m[1])
	}
	return request
}
