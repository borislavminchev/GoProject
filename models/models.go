package models


type Person struct {
	
	Id int `gorm:"primaryKey" gorm:"id"`
	Name string `gorm:"name"`
	Age int `gorm:"age"`
	Address string `gorm:"address"`
	Email string `gorm:"email"`
	Phone string `gorm:"phone"`
	Wallet string `gorm:"wallet"`
	Friends []*Person `gorm:"many2many:friends;foreignKey:id;joinForeignKey:person_id;References:id;joinReferences:friend_id"`
}

func (Person) TableName() string {
	return "person"
  }

