package main

import (
	// "flag"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/gocolly/colly"
)

var (
	Token string
)

// func init() {
// 	flag.StringVar(&Token, "t", "", "Bot Token")
// 	flag.Parse()
// }

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

func main() {

	fmt.Println("hello")
	// Create a new Discord session using the provided bot token.
	// dg, err := discordgo.New("Bot " + Token)
	dg, err := discordgo.New("Bot " + "PLACE HOLDER TOKEN")
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	fmt.Println(m.Content)

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "!help" {
		imageURL := formatLink("7c8m6c")
		response, err := http.Get(imageURL)
		if err != nil {
			fmt.Println(err)
		}
		defer response.Body.Close()

		if response.StatusCode == 200 {
			_, err = s.ChannelFileSend(m.ChannelID, imageURL, response.Body)
			s.ChannelMessageSend(m.ChannelID, "<:hell:1073060956353089606> Forgot about Google's memegen, this is the real memegen <:hell:1073060956353089606>")
			s.ChannelMessageSend(m.ChannelID, "Commands: \n Tell the time: !time \n Get a random meme: !random \n Get the top meme rn: !topmeme \n Search a template: !template <keyword> \n Search by meme code: !code <meme code>")
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Error: oof something went wrong :-(")
		}
	}

	if m.Content == "!helmo" {
		imageURL := "https://www.mypokecard.com/my/galery/L8tjeLWgmPD2.jpg"
		response, err := http.Get(imageURL)
		if err != nil {
			fmt.Println(err)
		}
		defer response.Body.Close()

		if response.StatusCode == 200 {
			_, err = s.ChannelFileSend(m.ChannelID, imageURL, response.Body)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Error: helmo is OOO :-(")
		}
	}

	if m.Content == "!time" {
		imageURL := formatLink("7c8jyi")
		response, err := http.Get(imageURL)
		if err != nil {
			fmt.Println(err)
		}
		defer response.Body.Close()

		if response.StatusCode == 200 {
			_, err = s.ChannelFileSend(m.ChannelID, imageURL, response.Body)
			s.ChannelMessageSend(m.ChannelID, "<:yogurt:1073061148422836284> It's Yogurt time! <:yogurt:1073061148422836284>")
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Error: helmo is OOO :-(")
		}
	}

	if m.Content == "!random" {
		imageCode := randomMeme()
		imageURL := "https://i.imgflip.com/" + imageCode + ".jpg"
		response, err := http.Get(imageURL)
		if err != nil {
			fmt.Println(err)
		}
		defer response.Body.Close()

		if response.StatusCode == 200 {
			_, err = s.ChannelFileSend(m.ChannelID, imageURL, response.Body)
			s.ChannelMessageSend(m.ChannelID, "Meme code:"+imageCode)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Error: Can't get random meme! :-(")
		}
	}

	if m.Content == "!topmeme" {
		keyword := topMeme()
		imageURL := formatLink(keyword)
		response, err := http.Get(imageURL)
		if err != nil {
			fmt.Println(err)
		}
		defer response.Body.Close()

		if response.StatusCode == 200 {
			_, err = s.ChannelFileSend(m.ChannelID, imageURL, response.Body)
			s.ChannelMessageSend(m.ChannelID, "Meme code:"+keyword)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Error: Can't get top meme! :-(")
		}
	}

	if strings.HasPrefix(m.Content, "!template") {
		keyword := m.Content[10:]
		searchResult, err := searchTemplate(keyword)
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, err.Error())
		} else {
			imageURL := formatLink(searchResult)
			response, err := http.Get(imageURL)
			if err != nil {
				fmt.Println(err)
			}
			defer response.Body.Close()

			if response.StatusCode == 200 {
				_, err = s.ChannelFileSend(m.ChannelID, imageURL, response.Body)
				s.ChannelMessageSend(m.ChannelID, "Meme code:"+searchResult)
				if err != nil {
					fmt.Println(err)
				}
			} else {
				fmt.Println("Error: Can't get top meme! :-(")
			}
		}
	}

	if strings.HasPrefix(m.Content, "!code") {
		keyword := m.Content[6:]
		imageURL := formatLink(keyword)
		response, err := http.Get(imageURL)
		if err != nil {
			fmt.Println(err)
		}
		defer response.Body.Close()

		if response.StatusCode == 200 {
			_, err = s.ChannelFileSend(m.ChannelID, imageURL, response.Body)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Error: Can't get top meme! :-(")
		}

	}
}
