package builder

import(
	"github.com/go-elastic-search/src/api/request"
)

//DSQueryBuilder - Process GetDocument struct to build elastic search query
type DSQueryBuilder interface {
	New() DSQueryBuilder

	AddSortBuilder(request.GetDocument) DSQueryBuilder
	AddLimit(request.GetDocument) DSQueryBuilder
	AddOffset(request.GetDocument) DSQueryBuilder
	AddRange(request.GetDocument) DSQueryBuilder
	AddFilterByStrings(request.GetDocument) DSQueryBuilder
	AddFilterByDates(request.GetDocument) DSQueryBuilder
	AddFilterByNumbers(request.GetDocument) DSQueryBuilder
	AddFilterByBooleans(request.GetDocument) DSQueryBuilder
}
