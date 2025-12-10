package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/developerasun/SignalDash/server/models"
	"github.com/developerasun/SignalDash/server/sderror"
	"github.com/gocolly/colly/v2"
	"gorm.io/gorm"
)

type indicator struct {
	crawler *colly.Collector
}

type Indicator interface {
	ScrapeDollarIndex() (dxy string, err error)
}

func NewIndicator(domains []string, botHeader string) Indicator {
	_crawler := NewCrawler(domains, botHeader)

	return &indicator{
		crawler: _crawler,
	}
}

func (i indicator) ScrapeDollarIndex() (string, error) {
	c := i.crawler

	var dxy string
	var hasError error = nil
	c.OnHTML("section[data-an-section-id=symbol-overview-page-section]", func(e *colly.HTMLElement) {
		substringToReplace := "The current value of U.S. Dollar Index is"
		expression := `(\d+\.\d+)`
		full := fmt.Sprintf("%s %s", substringToReplace, expression)
		target := regexp.MustCompile(full)

		match := target.FindString(e.Text)
		extracted := strings.Replace(match, substringToReplace, "", 1)
		log.Println("match: ", match, "extracted: ", extracted)

		dxy = extracted
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Error:", err)
		hasError = err
	})

	if vErr := c.Visit("https://www.tradingview.com/symbols/TVC-DXY/"); vErr != nil {
		hasError = vErr
		return "", hasError
	}

	// @dev wait for jobs to anchor values without error
	c.Wait()

	return dxy, hasError
}

func FindLatestDollarIndex(db *gorm.DB) (*models.Indicator, error) {
	ctx := context.Background()
	indicator, lErr := gorm.G[models.Indicator](db).Last(ctx)
	if errors.Is(lErr, gorm.ErrRecordNotFound) {
		return nil, sderror.EmptyStorage
	}

	return &indicator, nil
}

// ================================================================== //
// ============================== deps ============================== //
// ================================================================== //

func NewCrawler(domains []string, botHeader string) *colly.Collector {
	return colly.NewCollector(
		colly.AllowedDomains(domains...),
		colly.UserAgent(botHeader),
		colly.IgnoreRobotsTxt(),
	)
}
