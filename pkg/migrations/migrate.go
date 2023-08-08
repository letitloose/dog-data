package migrations

import (
	"github.com/letitloose/dog-data/pkg/models"
)

type MigrateModel struct {
	LegacyModel models.LegacyModel
	DogModel    models.DogModel
}

func (model *MigrateModel) migrateDogs() error {

	legacyTollers, err := model.LegacyModel.GetTollers()
	if err != nil {
		return err
	}

	for _, currentToller := range legacyTollers {
		currentDog := models.Dog{}
		currentDog.Regnum = currentToller.Regnum
		currentDog.Nsdtrcregnum = currentToller.Nsdtrcregnum
		currentDog.Sequencenum = currentToller.Sequencenum
		currentDog.Name = currentToller.Name
		currentDog.Callname = currentToller.Callname
		currentDog.Whelpdate = currentToller.Whelpdate
		currentDog.Callname = currentToller.Callname
		currentDog.Nba = yesNoToBool(currentToller.Nba)
		currentDog.Alive = yesNoToBool(currentToller.Alive)
		currentDog.Intact = yesNoToBool(currentToller.Intact)
		currentDog.Litterid = 1
		currentDog.Sex = "SD"
		currentDog.Sire = 1
		currentDog.Dam = 1
		//sex lookup
		// Sex          string
		//parent lookup
		// Sire         int
		// Dam          int

		err := model.DogModel.Insert(currentDog)
		if err != nil {
			return err
		}
	}
	return nil
}

func yesNoToBool(yesNo string) bool {
	return yesNo == "Y"
}
