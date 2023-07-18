package lang

const (
	AccountNotFound = "Account.NotFound"
)

var AllLanguageMap = map[string]map[string]string{
	"cn":    CnLanguageMap,
	"zh-cn": CnLanguageMap,
	"en":    EnLanguageMap,
}
