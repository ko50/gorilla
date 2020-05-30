package scraping

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

// GetSearchResults 検索キーワードを受け取り、Googleで検索した上位3つのURLを返します
func GetSearchResults(url string) {
	doc, _ := goquery.NewDocument(url)
	doc.Find(".CloneDocument").Each(func(i int, s *goquery.Selection) {
		s.Find("pre").Each(func(ii int, s *goquery.Selection) {
			if ii >= 3 {
				return
			}
			fmt.Printf("ウホ#%d", ii)
			fmt.Println(s.Text())
		})

		//doc.Find(".r").Each(func(i int, s *goquery.Selection) {
		if i >= 3 {
			return
		}
		//url, _ := s.Attr("href")
		//fmt.Println(s.Find("a").Attr("href"))*/
		// fmt.Println(txt)
	})
}
