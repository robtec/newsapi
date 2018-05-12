package api

// Options can be used when making requests to the newsapi
type Options struct {

	// The 2-letter ISO 3166-1 code of the country you want to get headlines for.
	// Possible options:
	// ae ar at au be bg br ca ch cn co cu cz de eg fr gb gr
	// hk hu id ie il in it jp kr lt lv ma mx my ng nl no nz
	// ph pl pt ro rs ru sa se sg si sk th tr tw ua us ve za
	// Note: you can't mix this param with the sources param.
	Country string

	// The category you want to get headlines for. Possible options:
	// business entertainment general health science sports technology
	// Note: you can't mix this param with the sources param.
	Category string

	// A comma-separated string of domains
	// (eg bbc.co.uk, techcrunch.com, engadget.com) to restrict the search to
	Domains string

	// A date and optional time for the oldest article allowed. This should be in ISO 8601 format
	// (e.g. 2018-05-12 or 2018-05-12T15:29:05) Default: the oldest according to your
	// plan.
	From string

	// A date and optional time for the newest article allowed. This should be in ISO 8601 format
	// (e.g. 2018-05-12 or 2018-05-12T15:29:05) Default: the newest according to your
	// plan.
	To string

	// The 2-letter ISO-639-1 code of the language you want to get headlines for. Possible options:
	// ar de en es fr he it nl no pt ru se ud zh. Default: all languages returned.
	Language string

	// The order to sort the articles in. Possible options:
	// relevancy, popularity, publishedAt.
	// relevancy = articles more closely related to q come first.
	// popularity = articles from popular sources and publishers come first.
	// publishedAt = newest articles come first.
	// Default: publishedAt
	SortBy string

	// A comma-separated string of identifiers for the news sources or blogs you want headlines from.
	// Use the /sources endpoint to locate these programmatically or look at the sources index.
	// Note: you can't mix this param with the country or category params.
	Sources string

	// The number of results to return per page (request).
	// 20 is the default, 100 is the maximum.
	PageSize int

	// Use this to page through the results if the total results found is greater than the page size.
	Page int

	// Keywords or a phrase to search for.
	Q string
}
