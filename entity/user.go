package entity

import "time"

type User struct {
	Id                                        uint32 `gorm:"primary_key;auto_increment:true"`
	Age                                       uint32
	FirstName, LastName, City, State, Country string
	Email                                     string `gorm:"type:varchar(100);unique_index"`
	CreatedAt, ModifiedAt                     time.Time
}

func (user User) IsValid() {
}
