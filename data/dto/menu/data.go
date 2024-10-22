package menu

type Data struct {
	Pagination Pagination `json:"pagination"`
	Contents   []Content  `json:"contents"`
}
