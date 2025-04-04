package modules

type Test struct {
	Test string
}

func (Test) TableName() string {
	return "test"
}
