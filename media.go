package main

import (
	"os"
	"strings"

	req "github.com/imroc/req"
	twitterscraper "github.com/n0madic/twitter-scraper"
	log "github.com/sirupsen/logrus"
)

func downloadMedia(link string, dest string, tweet *twitterscraper.TweetResult) {
	r, _ := req.Get(link)
	if r == nil {
		log.Fatalf("Error, unable to download \"%v\".\n", link)
	}
	r.ToFile(dest)
	if tweet != nil {
		err := os.Chtimes(dest, tweet.TimeParsed, tweet.TimeParsed)
		if err != nil {
			log.Fatal(err.Error())
		}
	}
}

func getImages(tweet *twitterscraper.TweetResult, tweet_folder_path string) {
	all_photos := []string{}

	//Get photos from normal, reply, quotes, and retweeted tweets

	if tweet.Photos != nil {
		for item := range tweet.Photos {
			all_photos = append(all_photos, tweet.Photos[item])
		}
	}

	if tweet.InReplyToStatus != nil {
		if tweet.InReplyToStatus.Photos != nil {
			for item := range tweet.InReplyToStatus.Photos {
				all_photos = append(all_photos, tweet.InReplyToStatus.Photos[item])
			}
		}
	}

	if tweet.QuotedStatus != nil {
		if tweet.QuotedStatus.Photos != nil {
			for item := range tweet.QuotedStatus.Photos {
				all_photos = append(all_photos, tweet.QuotedStatus.Photos[item])
			}
		}
	}

	if tweet.RetweetedStatus != nil {
		if tweet.RetweetedStatus.Photos != nil {
			for item := range tweet.RetweetedStatus.Photos {
				all_photos = append(all_photos, tweet.RetweetedStatus.Photos[item])
			}
		}
	}

	// Get file names and download images

	for item := range all_photos {
		filename := all_photos[item]
		filename = filename[strings.LastIndex(filename, "/")+1:]
		log.Println("Downloading tweet ID \"" + string(tweet.ID) + "\" image \"" + filename + "\"...")
		downloadMedia(all_photos[item], tweet_folder_path+"/"+filename, tweet)
	}

}

func getVideos(tweet *twitterscraper.TweetResult, tweet_folder_path string) {
	all_videos := []string{}

	//Get videos from normal, reply, quotes, and retweeted tweets

	if tweet.Videos != nil {
		for item := range tweet.Videos {
			all_videos = append(all_videos, tweet.Videos[item].URL)
		}
	}

	if tweet.InReplyToStatus != nil {
		if tweet.InReplyToStatus.Videos != nil {
			for item := range tweet.InReplyToStatus.Videos {
				all_videos = append(all_videos, tweet.InReplyToStatus.Videos[item].URL)
			}
		}
	}

	if tweet.QuotedStatus != nil {
		if tweet.QuotedStatus.Videos != nil {
			for item := range tweet.QuotedStatus.Videos {
				all_videos = append(all_videos, tweet.QuotedStatus.Videos[item].URL)
			}
		}
	}

	if tweet.RetweetedStatus != nil {
		if tweet.RetweetedStatus.Videos != nil {
			for item := range tweet.RetweetedStatus.Videos {
				all_videos = append(all_videos, tweet.RetweetedStatus.Videos[item].URL)
			}
		}
	}

	// Get file names and download images

	for item := range all_videos {
		filename := all_videos[item]
		filename = filename[strings.LastIndex(filename, "/")+1:]
		if strings.Contains(filename, "?") {
			filename = filename[:strings.Index(filename, "?")]
		}
		if strings.Contains(filename, "#") {
			filename = filename[:strings.Index(filename, "#")]
		}

		if !(strings.Contains(all_videos[item], "youtu.be")) && !(strings.Contains(all_videos[item], "youtube")) {
			log.Println("Downloading tweet ID \"" + string(tweet.ID) + "\" video \"" + filename + "\"...")
			downloadMedia(all_videos[item], tweet_folder_path+"/"+filename, tweet)
		}
	}

}
