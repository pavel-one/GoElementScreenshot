package browser

type ScreenshotRequest struct {
	Element string `json:"element"`
	Name    string `json:"name"`
	Url     string `json:"url"`
}
