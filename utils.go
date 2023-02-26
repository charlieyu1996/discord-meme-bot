package main

import (
	"errors"
	"fmt"
	"github.com/gocolly/colly"
	"math/rand"
	"strings"
)

func topMeme() string {
	var imageURL string
	c := colly.NewCollector()

	c.OnHTML("a.base-img-link", func(e *colly.HTMLElement) {
		imageURL = e.Attr("href")[3:]
	})

	c.Wait()
	startUrl := "https://imgflip.com"
	c.Visit(startUrl)

	return imageURL
}

func randomMeme() string {
	var imageURL string
	c := colly.NewCollector()

	c.OnHTML("body", func(e *colly.HTMLElement) {
		imageURL = e.Text[3:]
	})

	c.Wait()
	startUrl := "https://imgflip.com/ajax_img_flip"
	c.Visit(startUrl)

	return imageURL
}

func searchTemplate(search string) (string, error) {
	memes := make([]string, 0, 60)
	specificMemes := make([]string, 0, 15)
	c := colly.NewCollector()
	c2 := colly.NewCollector()

	c.OnHTML("a.s-result", func(e *colly.HTMLElement) {
		meme := e.Attr(("href"))
		if strings.HasPrefix(meme, "/meme/") {
			memes = append(memes, meme)
		}
	})

	c2.OnHTML("a.base-img-link", func(e *colly.HTMLElement) {
		meme := e.Attr(("href"))
		specificMemes = append(specificMemes, meme)
	})

	startUrl := "https://imgflip.com/search?q=" + search
	c.Visit(startUrl)

	// keep trying until there is a template with content
	for len(specificMemes) == 0 && len(memes) != 0 {
		randomIndex := rand.Intn(len(memes))
		selectedMemeTemplate := memes[randomIndex]
		if randomIndex+1 > len(memes) {
			memes = memes[:randomIndex]
		} else {
			memes = append(memes[:randomIndex], memes[randomIndex+1:]...)
		}
		fmt.Println("Selected meme template: " + selectedMemeTemplate)
		newURL := "https://imgflip.com" + selectedMemeTemplate
		c2.Visit(newURL)
	}

	if len(memes) == 0 {
		return "", errors.New("no templates found for this keyword :(")
	}

	selectedMeme := specificMemes[rand.Intn(len(specificMemes))][3:]

	fmt.Println("Selected meme: " + selectedMeme)

	return selectedMeme, nil
}

// this function checks if the link is a GIF or an image
func formatLink(link string) string {
	if strings.HasPrefix(link, "f/") {
		// it is a GIF
		return "https://i.imgflip.com/" + link[2:] + ".mp4"
	}

	return "https://i.imgflip.com/" + link + ".jpg"
}
