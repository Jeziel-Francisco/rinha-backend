package dto

type DefaultResponse struct {
	ResponseBody    interface{}
	ResponseHeaders map[string]string
	ResponseCode    int
}
