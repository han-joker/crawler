package errors

import (
	"bytes"
	"fmt"
)

const (
	ERROR_TYPE_DOWNLOADER ErrorType = "downloader errors"
	ERROR_TYPE_ANALYZER   ErrorType = "analyzer errors"
	ERROR_TYPE_PIPELINE   ErrorType = "pipeline errors"
	ERROR_TYPE_SCHEDULER  ErrorType = "scheduler errors"
)

type myCrawlerError struct {
	errType    ErrorType
	errMsg     string
	fullErrMsg string
}

func (ce *myCrawlerError) Type() ErrorType {
	return ce.errType
}

func (ce *myCrawlerError) Error() string {
	if ce.fullErrMsg == "" {
		ce.genFullErrMsg()
	}
	return ce.fullErrMsg
}

func (ce *myCrawlerError) genFullErrMsg() {
	var buffer bytes.Buffer
	buffer.WriteString("crawler error: ")
	if ce.errType != "" {
		buffer.WriteString(string(ce.errType))
		buffer.WriteString(":")
	}
	buffer.WriteString(ce.errMsg)
	ce.fullErrMsg = fmt.Sprintf("%s", buffer.String())
	return
}

func NewCrawlerError(errType ErrorType, errMsg string) CrawlerError {
	return &myCrawlerError{
		errType:    errType,
		errMsg:     errMsg,
		fullErrMsg: "",
	}
}
