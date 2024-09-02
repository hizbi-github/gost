package module

import (
	"bytes"
	"context"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/mmcdole/gofeed"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	models "github.com/hizbi-github/gost/new-project-core/models"
	moduleRepo "github.com/hizbi-github/gost/new-project-core/module/repo"
	utils "github.com/hizbi-github/gost/new-project-core/utils"
)

//- https://news.mc/feed/
//- https://www.monaco-tribune.com/en/category/news/feed/

type Scraper struct {
	Interval   int64
	RssUrl     string
	WebsiteUrl string
	Database   *mongo.Database
}

func NewScraper(scraperInput Scraper) *Scraper {
	//Now we can manage some default values here...
	scraper := scraperInput
	return &scraper
}

func (s Scraper) StartCronJob() {
	ticker := time.NewTicker(20 * time.Second)

	for {
		<-ticker.C
		err := s.someJob()
		if err != nil {
			logrus.Errorf("some cron job error: %+v", err)
			continue
		}
	}
}

func (s Scraper) someJob() error {
	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()

	feed, err := gofeed.NewParser().ParseURLWithContext("https://news.mc/feed/", ctx)
	if err != nil {
		return err
	}

	for _, item := range feed.Items {
		if moduleRepo.Exists(ctx, s.Database, item.GUID) {
			continue
		}

		request := &models.HttpRequest{
			Url:     "some_url",
			Headers: nil,
			Body:    nil,
		}
		response, err := utils.HttpGet(request)
		if err != nil {
			logrus.Errorln(err)
			continue
		}

		parsedHtml, err := goquery.NewDocumentFromReader(bytes.NewReader(response.Body))
		if err != nil {
			logrus.Errorln(err)
			continue
		}

		articleContent := utils.Trim(parsedHtml.Find(".post-content").Text())
		if articleContent == "" {
			logrus.Errorln(err)
			continue
		}

		someMongoDocument := models.SomeMongoDocument{
			Id:        primitive.NewObjectID(),
			SomeKey:   "some_value",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		err = moduleRepo.Save(ctx, s.Database, &someMongoDocument)
		if err != nil {
			logrus.Errorln(err)
			continue
		}

		logrus.Infof("some log: %s, some log: %s", &someMongoDocument.Id, someMongoDocument.SomeKey)
	}

	return nil
}
