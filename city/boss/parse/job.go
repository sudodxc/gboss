package parse

import (
	"city/engine"
	"regexp"
)

var urlall = regexp.MustCompile(`<a href="(/job_detail/[\w|~]+[\w\W]*?)data-jid="[\w|~]+[\w\W]*?data-itemid="\d+"[\w\W]*?lid="[\w|.]+[\w\W]*?jobid="\d+"[\w\W]*?class="job-title">(.*)<[\w\W]*?class="red">([\w|-]+)</span>`)
var nexturl = regexp.MustCompile(`<a href="(/c101020100-p100103/?.*)" ka="page-[0-9]">[0-9]</a>`)

// const urlall = `<p>(.*?)<em class="vline"></em>(.*?)<em class="vline"></em>(.*?)</p>`
func ParseJob(contents []byte) engine.ParseResult {
	//hao

	match := urlall.FindAllSubmatch(contents, -1)
	request := engine.ParseResult{}
	for _, m := range match {
		//fmt.Printf("is urls:%s", m[1]) //url
		name := string(m[2])
		url := "https://www.zhipin.com" + string(m[1])
		postSalary := string(m[3])
		// profile := model.Profile{}
		// profile.Url = "https://www.zhipin.com" + string(m[1])
		// profile.PostSalary = string(m[3])
		// profile.Name = name
		//request.Items = append(request.Items, profile) //jobname
		request.Requests = append(request.Requests, engine.Request{
			Url: "https://www.zhipin.com" + string(m[1]), //url
			ParseFunc: func(c []byte) engine.ParseResult {
				//jobs
				return ParseJobs(c, name, url, postSalary)
			}, //
		})

		//fmt.Printf("jobname:%s ,nextURL:https://www.zhipin.com%s\n", m[2], m[1])
	}
	matchs := nexturl.FindAllSubmatch(contents, -1)
	for _, m := range matchs {
		request.Items = append(request.Items, "nexturl:"+"https://www.zhipin.com"+string(m[1]))
		request.Requests = append(request.Requests,
			engine.Request{
				Url: "https://www.zhipin.com" + string(m[1]),
				ParseFunc: func(c []byte) engine.ParseResult { //jobs
					return ParseJob(c)
				}, //
			})
	}

	return request

}
