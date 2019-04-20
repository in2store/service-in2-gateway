package constants

type PaginationInfo struct {
	NextPage  int32 `json:"nextPage"`
	PrevPage  int32 `json:"prevPage"`
	FirstPage int32 `json:"firstPage"`
	LastPage  int32 `json:"lastPage"`
}
