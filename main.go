package main

import (
	"bufio"
	"context"
	"encoding/json"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"

	twitterscraper "github.com/n0madic/twitter-scraper"
)

var _ = os.Mkdir("./bin/", os.ModePerm)

func checkDatabase(id string) bool {

	//Open ./bin/database.txt, read file line by line, check for id in string

	file, err := os.Open("./bin/database.txt")
	if err != nil {
		_, err := os.Create("./bin/database.txt")
		if err != nil {
			log.Fatalln(err)
		}
		return false
	}
	defer file.Close()

	res := false

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if id == scanner.Text() {
			res = true
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return res
}

func getProfile(creator string, scraper *twitterscraper.Scraper) {

	profile, err := scraper.GetProfile(creator)
	folder_path := getSettings().Path + "/" + creator + "/"
	if err != nil {
		log.Fatalln(err)
	}

	os.MkdirAll(folder_path, os.ModePerm)

	profilejson, err := json.MarshalIndent(profile, "", "\t")
	if err != nil {
		log.Fatalln(err)
	}

	f, err := os.Create(folder_path + creator + ".json")
	if err != nil {
		log.Fatalln(err)
	}

	_, err2 := f.WriteString(string(profilejson))
	if err2 != nil {
		log.Fatalln(err)
	}
	f.Close()

	avatar_path := profile.Avatar
	banner_path := profile.Banner

	if strings.Contains(avatar_path, "_normal") {
		avatar_path = strings.ReplaceAll(avatar_path, "_normal", "")
	}

	if strings.Contains(banner_path, "_normal") {
		banner_path = strings.ReplaceAll(banner_path, "_normal", "")
	}

	downloadMedia(avatar_path, folder_path+"/avatar.jpg", nil)
	downloadMedia(banner_path, folder_path+"/banner.jpg", nil)
}

func appendDB(msg string) {
	db, err := os.OpenFile("./bin/database.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	if _, err := db.WriteString(msg + "\n"); err != nil {
		log.Println(err)
	}
	db.Close()
}

func main() {

	settings := getSettings()

	for i := 0; i < len(settings.Creators); i++ {

		creator := settings.Creators[i]

		scraper := twitterscraper.New().WithDelay(3)

		log.Println("Scraping tweets from \"" + creator + "\"...")

		//Get user profile information

		if !(checkDatabase(creator)) {
			getProfile(creator, scraper)
			appendDB(creator)
		}

		//Get tweets

		for tweet := range scraper.WithReplies(true).GetTweets(context.Background(), creator, 3200) { //3200 is the limt of tweets in the api. Sadge.
			if tweet.Error != nil {
				log.Fatalln(tweet.Error)
			}

			//Check database to see if tweet ID exists, convert to json if not

			if checkDatabase(string(tweet.ID)) {
				log.Println("\"" + creator + "\" tweet ID \"" + tweet.ID + "\" already in database, skipping...")
				continue
			}

			log.Println("Downloading \"" + creator + "\" tweet ID \"" + tweet.ID + "\"...")

			tweetjson, err := json.MarshalIndent(tweet, "", "\t")
			if err != nil {
				log.Fatalln(err)
			}

			//Make tweet folder path then save tweet to json

			folder_path := settings.Path + "/" + creator + "/tweets/" + string(tweet.ID) + "/"

			os.MkdirAll(folder_path, os.ModePerm)

			f, err := os.Create(folder_path + string(tweet.ID) + ".json")
			if err != nil {
				log.Fatalln(err)
			}

			_, err2 := f.WriteString(string(tweetjson))
			if err2 != nil {
				log.Fatalln(err)
			}
			f.Close()

			//Change modified and created time of the tweet json to the tweet date

			err = os.Chtimes(folder_path+string(tweet.ID)+".json", tweet.TimeParsed, tweet.TimeParsed)
			if err != nil {
				log.Fatal(err.Error())
			}

			//Iterate through images and download each

			getImages(tweet, folder_path)

			//Iterate through videos and download each

			getVideos(tweet, folder_path)

			//Change modified and created time of the tweet folder to the tweet date

			err = os.Chtimes(folder_path, tweet.TimeParsed, tweet.TimeParsed)
			if err != nil {
				log.Fatal(err.Error())
			}

			//Append to database

			appendDB(string(tweet.ID))

			//Print success

			log.Println("Successfully downloaded \"" + creator + "\" tweet ID \"" + tweet.ID + "\"!")

		}
	}

	log.Println("Completed downloading all tweets successfully!")
}
