package language

import (
	"ccps.com/internal/utils"
	"fmt"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strconv"
	"strings"
)

type LangMap = map[string]map[string]string
type Loader struct {
	lang        string
	defaultLang string
	langMap     LangMap
}

func Load(langMap LangMap) *Loader {
	return &Loader{langMap: langMap}
}

func (l *Loader) SetDefaultLang(lang string) *Loader {
	l.defaultLang = lang
	return l
}

func (l *Loader) SetLang(lang string) *Loader {
	l.lang = lang
	return l
}

func (l *Loader) GetLang() string {
	return l.lang
}

func (l *Loader) SetLangByRequest(req *http.Request) *Loader {
	lang := req.URL.Query().Get("lang")

	if lang == "" {
		quality := 0.0
		headerAcceptLanguage := req.Header.Get("Accept-Language")
		acceptLanguages := strings.Split(headerAcceptLanguage, ",")
		for _, s := range acceptLanguages {
			var a, q = "", 1.0
			arr := strings.Split(s, ";q=")
			a = utils.ArrayItem(arr, 0, "")
			q, _ = strconv.ParseFloat(utils.ArrayItem(arr, 1, "1.0"), 64)

			if q > quality || utils.FloatEqual(q, quality) {
				quality = q
				lang = a
			}
		}
	}

	l.SetLang(utils.EmptyString(lang, l.defaultLang))
	return l
}

func (l *Loader) Trans(key string) *Entry {
	languageMap := l.langMap[strings.ToLower(l.lang)]

	return NewEntry(languageMap[key])
}

func (l *Loader) TransByValidator(val validator.ValidationErrors) *Entry {
	errors := val
	e := errors[0]

	key := "Validator." + e.Namespace() + "." + e.Tag()
	return l.Trans(key).Replace(map[string]string{
		"field": e.Field(), "value": fmt.Sprintf("%v", e.Value()),
		"param": e.Param(), "type": e.Type().String(),
	})
}
