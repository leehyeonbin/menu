package dto

type Content struct {
	SeqNo          int        `json:"seqNo"`
	PortalSeqNo    int        `json:"portalSeqNo"`
	NoticeType     int        `json:"noticeType"`
	BoardID        string     `json:"boardId"`
	BoardType      int        `json:"boardType"`
	Title          string     `json:"title"`
	Contents       string     `json:"contents"`
	UserName       string     `json:"userName"`
	CreateUserID   string     `json:"createUserId"`
	UpdateUserID   string     `json:"updateUserId"`
	ViewCount      int        `json:"viewCount"`
	Status         int        `json:"status"`
	UserIP         string     `json:"userIp"`
	CreateDateTime string     `json:"createDateTime"`
	UpdateDateTime string     `json:"updateDateTime"`
	ReplyCount     int        `json:"replyCount"`
	CreateUser     CreateUser `json:"createUser"`
}
