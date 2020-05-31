package module

import "net/http"

type Response struct {
	httpResponse *http.Response
	depth        uint32
}

func NewResponse(httpResponse *http.Response, depth uint32) *Response {
	return &Response{
		httpResponse: httpResponse,
		depth:        depth,
	}
}

func (resp *Response) HTTPResponse() *http.Response {
	return resp.httpResponse
}

func (resp *Response) Depth() uint32 {
	return resp.depth
}

func (resp *Response) Valid() bool {
	return resp.httpResponse != nil && resp.httpResponse.Body != nil
}
