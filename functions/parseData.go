package functions

import (
	"fmt"

	"github.com/gocolly/colly"
)

/*https://www.kinopoisk.ru/index.php?level=7&from=forma&result=adv&m_act%5Bfrom%5D=forma&m_act
%5Bwhat%5D=content&m_act%5Bfind%5D=Lalalend&m_act%5Byear%5D=2017&m_act%5Bcountry%5D=1&m_act%5Bgenre%5D%5B%5D=11
&m_act%5Bactor%5D=Rayan&m_act%5Bcast%5D=Roy&m_act%5Bcontent_find%5D=film&m_act%5Bgenre_and%5D=on
*/

/*
https://www.kinopoisk.ru/index.php?level=7&from=forma&result=adv&m_act%5Bfrom%5D=forma&m_act
%5Bwhat%5D=content&m_act%5Bfind%5D=Texas&m_act%5Byear%5D=2012&m_act%5Bcountry%5D=136&m_act%5Bgenre%5D%5B%5D=11
&m_act%5Bactor%5D=Dylan&m_act%5Bcast%5D=Roy&m_act%5Bcontent_find%5D=serial&m_act%5Bgenre_and%5D=on
*/
func ParseData(movie string) []string {
	res := []string{}
	name, nextUrl, timing, info, rating := "", "", "", "", ""
	VisitUrl := "https://www.kinopoisk.ru/index.php?level=7&from=forma&result=adv&m_act%5Bfrom%5D=forma&m_act%5Bwhat%5D=content&m_act%5Bfind%5D=" + movie
	c := colly.NewCollector()

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("We are visiting this site:", r.Request.URL)
	})

	c.OnHTML(".element.most_wanted", func(h *colly.HTMLElement) {
		name = h.ChildText(".info>p.name")
		timing = h.ChildText("span:nth-child(2)")
		info = h.ChildText("span:nth-child(3)")
		rating = h.ChildText(".right>div")
		nextUrl = h.ChildAttr("a", "href")
		nextUrl = "https://www.kinopoisk.ru" + nextUrl[0:len(nextUrl)-5]
	})

	c.Visit(VisitUrl)

	res = append(res, name, timing[4:], info, rating)
	res = append(res, fmt.Sprint("for detailed information u can visit this site:", nextUrl))

	return res
}
