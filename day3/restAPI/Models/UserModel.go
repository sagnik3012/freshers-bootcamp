package Models

type Student struct {
	ID        uint   `json:"id" gorm:"autoIncrement"`
	Stu_id    uint   `json:"stu_id" `
	FirstName string `json:"firstname" `
	LastName  string `json:"surname" `
	DOB       string `json:"dob"`
	Address   string `json:"address" `
	Subject   string `json:"subject" `
	Marks     uint   `json:"marks" `
}

func (b *Student) TableName() string {
	return "Students"
}
