package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"net/http"
	"strings"
)

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
