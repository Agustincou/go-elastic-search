package config

import "github.com/go-elastic-search/src/api/statics"

const (
	TIME_LAYOUT           string = "2006-01-02T15:04:05.999-0700"
	SEARCH_DEFAULT_LIMIT  int32  = 50
	SEARCH_DEFAULT_OFFSET int32  = 0
)

var (
	Scope             statics.Scope
	Environment       statics.Environment
	FieldsDSTypes     map[string]string
	SearchDateValues  = map[string]bool{"last_indexed": true}
	SearchItemsValues = map[string]bool{"sort": true, "criteria": true, "offset": true, "limit": true, "range": true, "end_date": true, "begin_date": true}
	RegExpSearchField = map[string]bool{"raw_row": true}
)

func init() {
	Scope = statics.BETA
	Environment = statics.DEVELOPMENT
}
