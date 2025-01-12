package dto

type CreateUser struct {
	SequenceNumber         int       `json:"sequenceNumber"`
	UserID                 string    `json:"userId"`
	UserName               string    `json:"userName"`
	SequenceNumberOfPortal *int      `json:"sequenceNumberOfPortal"`
	ManagementPortal       int       `json:"managementPortal"`
	RoleGroup              RoleGroup `json:"roleGroup"`
	Email                  *string   `json:"email"`
	PhoneNumber            string    `json:"phoneNumber"`
	BirthDate              *string   `json:"birthDate"`
	Gender                 *string   `json:"gender"`
	Picture                *string   `json:"picture"`
	CompanyCode            *string   `json:"companyCode"`
	EasyLogin              *bool     `json:"easyLogin"`
	PasswordChangeDate     *string   `json:"passwordChangeDate"`
	LastLoginDateTime      string    `json:"lastLoginDateTime"`
	DormancyDate           *string   `json:"dormancyDate"`
	WithdrawDate           *string   `json:"withDrawDate"`
	FailureCount           int       `json:"failureCount"`
	LockDateTime           *string   `json:"lockDateTime"`
	IsLock                 bool      `json:"isLock"`
	IsDelete               bool      `json:"isDelete"`
	CreateDateTime         string    `json:"createDateTime"`
	UpdateDateTime         string    `json:"updateDateTime"`
	Interest               *string   `json:"interest"`
	KakaoID                *string   `json:"kakaoId"`
	NaverID                *string   `json:"naverId"`
	GoogleID               *string   `json:"googleId"`
}
