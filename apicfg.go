package guerrillamail

import (
	"strings"
)

const (
	GuerrillamailAPI = "https://api.guerrillamail.com/ajax.php"
)

const (
	LANGUAGE_EN      = "en"
	LANGUAGE_FR      = "fr"
	LANGUAGE_NL      = "nl"
	LANGUAGE_RU      = "ru"
	LANGUAGE_TR      = "tr"
	LANGUAGE_UK      = "uk"
	LANGUAGE_AR      = "ar"
	LANGUAGE_KO      = "ko"
	LANGUAGE_JP      = "jp"
	LANGUAGE_ZH      = "zh"
	LANGUAGE_ZH_HANT = "zh-hant"
	LANGUAGE_DE      = "de"
	LANGUAGE_ES      = "es"
	LANGUAGE_IT      = "it"
	LANGUAGE_PT      = "pt"
)

var (
	SUPPORT_LANG = []string{
		"en",
		"fr",
		"nl",
		"ru",
		"tr",
		"uk",
		"ar",
		"ko",
		"jp",
		"zh",
		"zh-hant",
		"de",
		"es",
		"it",
		"pt",
	}
)

func checkLanguage(lang string) bool {
	lang = strings.ToLower(lang)

	for _, l := range SUPPORT_LANG {
		if l == lang {
			return true
		}
	}

	return false
}
