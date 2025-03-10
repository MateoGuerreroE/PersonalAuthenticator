package typings

type RegisterRequest struct {
	AppName     string `json:"appName"`
	AccountName string `json:"accountName"`
	Secret      string `json:"secret"`
}

type GenerateRequest struct {
	AppName string `json:"appName"`
}

type ControllerResponse struct {
	Data string `json:"data"`
}
