package browser

type Resource struct {
	UUID    string `json:"uuid" db:"uuid"`
	Url     string `json:"url" db:"url"`
	Element string `json:"element" db:"element"`
	Image   string `json:"image"`
}
