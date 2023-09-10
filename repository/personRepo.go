package repository

import (
	"example/database"
	"example/models"

	"gorm.io/gorm/clause"
)

type personRepo struct {}
func PersonRepository() *personRepo{ return &personRepo{} }

func (r *personRepo) GetAllPeople() []*models.Person {
	var p []*models.Person
	database.PersonDB.Preload(clause.Associations).Find(&p)
	return p
}

func (r *personRepo) GetPerson(id int) *models.Person {
	var p *models.Person
	database.PersonDB.Preload(clause.Associations).First(&p, id)
	return p
}

