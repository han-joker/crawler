package errors

type ErrorType = string

type CrawlerError interface {
	Type() ErrorType
	Error() string
}

