package goose

import (
	"time"
)

const defaultUserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_2) AppleWebKit/534.52.7 (KHTML, like Gecko) Version/5.1.2 Safari/534.52.7"

// Configuration is a wrapper for various config options
type Configuration struct {
	LocalStoragePath        string `json:"localStoragePath"` //not used in this version
	ImagesMinBytes          int    `json:"imagesMinBytes"`   //not used in this version
	TargetLanguage          string `json:"targetLanguage"`
	ImageMagickConvertPath  string `json:"imageMagickConvertPath"`  //not used in this version
	ImageMagickIdentifyPath string `json:"imageMagickIdentifyPath"` //not used in this version
	BrowserUserAgent        string `json:"browserUserAgent"`
	Debug                   bool   `json:"debug"`
	ExtractPublishDate      bool   `json:"extractPublishDate"`
	AdditionalDataExtractor bool   `json:"additionalDataExtractor"`
	EnableImageFetching     bool   `json:"enableImageFetching"`
	UseMetaLanguage         bool   `json:"useMetaLanguage"`

	//path to the stopwords folder
	stopWordsPath string
	stopWords     StopWords
	parser        *Parser

	timeout time.Duration
}

// GetDefaultConfiguration returns safe default configuration options
func GetDefaultConfiguration(args ...string) Configuration {
	if len(args) == 0 {
		return Configuration{
			LocalStoragePath:        "",   //not used in this version
			ImagesMinBytes:          4500, //not used in this version
			EnableImageFetching:     true,
			UseMetaLanguage:         true,
			TargetLanguage:          "en",
			ImageMagickConvertPath:  "/usr/bin/convert",  //not used in this version
			ImageMagickIdentifyPath: "/usr/bin/identify", //not used in this version
			BrowserUserAgent:        defaultUserAgent,
			Debug:                   false,
			ExtractPublishDate:      true,
			AdditionalDataExtractor: false,
			stopWordsPath:           "resources/stopwords",
			stopWords:               NewStopwords(), //TODO with path
			parser:                  NewParser(),
			timeout:                 time.Duration(5 * time.Second),
		}
	}
	return Configuration{
		LocalStoragePath:        "",   //not used in this version
		ImagesMinBytes:          4500, //not used in this version
		EnableImageFetching:     true,
		UseMetaLanguage:         true,
		TargetLanguage:          "en",
		ImageMagickConvertPath:  "/usr/bin/convert",  //not used in this version
		ImageMagickIdentifyPath: "/usr/bin/identify", //not used in this version
		BrowserUserAgent:        defaultUserAgent,
		Debug:                   false,
		ExtractPublishDate:      true,
		AdditionalDataExtractor: false,
		stopWordsPath:           "resources/stopwords",
		stopWords:               NewStopwords(), //TODO with path
		parser:                  NewParser(),
		timeout:                 time.Duration(5 * time.Second),
	}
}
