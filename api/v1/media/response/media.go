package response

//Login response payload
type MediaResponse struct {
	Filename string `json:"filename"`
}

//NewMediaResponse construct MediaResponse
func NewMediaResponse(filename string) *MediaResponse {
	var MediaResponse MediaResponse

	MediaResponse.Filename = filename

	return &MediaResponse
}
