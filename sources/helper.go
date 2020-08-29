package sources

import (
	"bufio"
	"github.com/PuerkitoBio/goquery"
	"github.com/anacrolix/torrent"
	"golang.org/x/net/html/charset"
	"github.com/AlecAivazis/survey/v2"
	"io"
	"log"
	"net/http"
)

func detectContentCharset(body io.Reader) string {
	r := bufio.NewReader(body)
	if data, err := r.Peek(1024); err == nil {
		if _, name, ok := charset.DetermineEncoding(data, ""); ok {
			return name
		}
	}
	return "utf-8"
}


func DocumentFromURL(url string) *goquery.Document {
	resp, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()


	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return doc
}

func DisplayMenu(items []string) (int,error) {
	prompt := &survey.Select{
		Message: "Choose a torrent:",
		Options: items,
	}
	var choice string

	survey.AskOne(prompt, &choice)

	var index int
	for i, v := range items {
		if v == choice {
			index = i
			break
		}
	}

	return index, nil
}

func DownloadFile(url string) {
	c, _ := torrent.NewClient(nil)
	defer c.Close()
	t, _ := c.AddMagnet(url)
	<-t.GotInfo()
	t.DownloadAll()
	c.WaitAll()
}