package utils

import (
	"fmt"
	"net"
	"net/url"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

var (
	instagramRegex = regexp.MustCompile(`^https:\/\/(www\.)?instagram\.com\/([A-Za-z0-9._]{1,30})\/?$`)
	linkedinRegex  = regexp.MustCompile(`^https:\/\/(www\.)?linkedin\.com\/in\/([A-Za-z0-9-_%]+)\/?$`)
	githubRegex    = regexp.MustCompile(`^https:\/\/(www\.)?github\.com\/([A-Za-z0-9-]+)\/?$`)

	blockShortenedURL = []string{
		"bit.ly",
		"tinyurl.com",
		"t.co",
		"goo.gl",
		"is.gd",
		"buff.ly",
		"rebrand.ly",
		"cutt.ly",
		"ow.ly",
	}
)

type SocialPlatform struct {
	Name  string `json:"name"`
	Regex *regexp.Regexp
}

var SocialPlatforms = []SocialPlatform{
	{
		Name:  "instagram",
		Regex: instagramRegex,
	},
	{
		Name:  "linkedin",
		Regex: linkedinRegex,
	},
	{
		Name:  "github",
		Regex: githubRegex,
	},
}

type SocialMediaInfo struct {
	Platform string `json:"platform"`
	Username string `json:"username"`
	URL      string `json:"url"`
}

func ExtractSocialMedia(rawURL string) (*SocialMediaInfo, error) {
	for _, platform := range SocialPlatforms {
		match := platform.Regex.FindStringSubmatch(rawURL)

		if len(match) > 2 {
			return &SocialMediaInfo{
				Platform: platform.Name,
				Username: match[2],
				URL:      rawURL,
			}, nil
		}
	}

	return nil, fmt.Errorf("Unsuported URL social media")
}

func ExtractSocialMediaBatch(
	urlSosmed []string,
) ([]*SocialMediaInfo, error) {

	socialMediaInfo := make([]*SocialMediaInfo, 0, len(urlSosmed))

	for i, url := range urlSosmed {
		info, err := ExtractSocialMedia(url)
		if err != nil {
			return nil, fmt.Errorf("url ke %d tidak valid %w", i+1, err)
		}

		socialMediaInfo = append(socialMediaInfo, info)
	}

	return socialMediaInfo, nil
}

func isPrivateOrLoopbackIP(host string) bool {
	ip := net.ParseIP(host)
	if ip == nil {
		return false // hostname normal not an ip
	}

	if ip.IsPrivate() || ip.IsLoopback() {
		return true
	}
	return false
}

func ValidateURL(fl validator.FieldLevel) bool {
	rawURL := fl.Field().String()

	//parse url
	u, err := url.Parse(rawURL)
	if err != nil {
		return false
	}

	// just https
	if u.Scheme != "https" {
		return false
	}

	host := u.Hostname()
	// block url shortener
	for _, s := range blockShortenedURL {
		if strings.EqualFold(host, s) {
			return false
		}
	}

	// bock private ip
	if isPrivateOrLoopbackIP(host) {
		return false
	}

	// path traversal
	path := strings.ToLower(u.Path)
	if strings.Contains(path, "../") || strings.Contains(path, "%2e%2e%2f") {
		return false
	}

	// batasi panjang maksimal
	if len(rawURL) > 300 {
		return false
	}

	// harus regex yang di pastikan
	if instagramRegex.MatchString(rawURL) ||
		linkedinRegex.MatchString(rawURL) ||
		githubRegex.MatchString(rawURL) {
		return true
	}

	return false
}
