package response

type (
	User struct {
		Id               string `json:"id,omitempty"`
		UserName         string `json:"username,omitempty"`
		Name             string `json:"name,omitempty"` //姓名
		Icon             string `json:"icon,omitempty"` //头像图片url
		Sex              string `json:"sex,omitempty"`  //性别
		NickName         string `json:"nickname,omitempty"`
		Email            string `json:"email,omitempty"`
		Mobile           string `json:"mobile,omitempty"`
		IsEmailConfirmed bool   `json:"emailConfirmed"`
		Birthday         *int64 `json:"birthday,omitempty"` //生日
	}
)

func (u User) TableName() string {
	return "t_user"
}
