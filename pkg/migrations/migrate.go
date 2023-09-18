package migrations

import (
	"log"
	"strings"

	"github.com/letitloose/dog-data/pkg/models"
)

type MigrateModel struct {
	LegacyModel    models.LegacyModel
	DogModel       models.DogModel
	CodetableModel models.CodetableModel
	LitterModel    models.LitterModel
	HealthModel    models.HealthModel
}

func (model *MigrateModel) MigrateDogs() error {

	legacyTollers, err := model.LegacyModel.GetTollers()
	if err != nil {
		return err
	}

	siresNF := map[string]int{}
	damsNF := map[string]int{}

	err = model.importSiresAndDams(legacyTollers)
	if err != nil {
		return err
	}

	male := model.CodetableModel.GetByCode("D", models.CategorySex)
	female := model.CodetableModel.GetByCode("B", models.CategorySex)
	if err != nil {
		return err
	}

	if err != nil {
		return err
	}
	for _, currentToller := range legacyTollers {

		if currentToller.Regnum == "" {
			continue
		}
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

		//sex lookup
		if currentToller.Sex == "D" {
			currentDog.Sex = male.Id
		} else {
			currentDog.Sex = female.Id
		}

		//parent lookup (by name since regnums can be dupes)
		sire, err := model.DogModel.GetByName(currentToller.Sirename)
		if err != nil {
			sire, err = model.DogModel.GetByRegnum(currentToller.Sireregnum)
			if err != nil {
				log.Printf("sire %s not found.  adding default for dog %s \n", currentToller.Sirename, currentToller.Regnum)
				siresNF[currentToller.Sireregnum] += 1
			}
		}
		currentDog.Sire = sire.Id

		dam, err := model.DogModel.GetByName(currentToller.Damname)
		if err != nil {
			dam, err = model.DogModel.GetByRegnum(currentToller.Damregnum)
			if err != nil {
				log.Printf("dam %s not found.  adding default for dog %s \n", currentToller.Damname, currentToller.Regnum)
				damsNF[currentToller.Damregnum] += 1
			}
		}

		currentDog.Dam = dam.Id

		//litter
		if currentToller.Litterregnum == "" {
			currentToller.Litterregnum = "no-litter"
		}
		litter, err := model.LitterModel.GetByRegnum(currentToller.Litterregnum)
		if err != nil {
			//litter not found, so insert it
			litter = models.Litter{Regnum: currentToller.Litterregnum}
			err = model.LitterModel.Insert(litter)
			if err != nil {
				return err
			}

			litter, err = model.LitterModel.GetByRegnum(currentToller.Litterregnum)
			if err != nil {
				return err
			}
		}

		currentDog.Litterid = litter.Id

		existingDog, err := model.DogModel.GetByRegnum(currentToller.Regnum)
		if err == nil {
			//we found the dog, so update it
			currentDog.Id = existingDog.Id
			err = model.DogModel.Update(currentDog)
			if err != nil {
				return err
			}
		} else {
			err = model.DogModel.Insert(currentDog)
			if err != nil {
				return err
			}
		}

		//insert health checks
		err = model.migrateHealthInfo(currentToller)
		if err != nil {
			return err
		}

	}
	log.Println("sires not found:", siresNF)
	log.Println("dams not found:", damsNF)
	return nil
}

func (model *MigrateModel) migrateHealthInfo(toller *models.Toller) error {

	dog, err := model.DogModel.GetByRegnum(toller.Regnum)
	if err != nil {
		return err
	}

	//hips
	if toller.HipClear != "" {
		health := models.Health{
			Dogid:      dog.Id,
			HealthType: "HHIP",
			CertId:     toller.HipClear,
		}

		err = model.HealthModel.Insert(health)
		if err != nil {
			return err
		}
	}

	//eyes
	if toller.EyeClear != "" {
		health := models.Health{
			Dogid:      dog.Id,
			HealthType: "HEYE",
			CertId:     toller.EyeClear,
		}

		err = model.HealthModel.Insert(health)
		if err != nil {
			return err
		}
	}

	//heart
	if toller.HeartClear != "" {
		health := models.Health{
			Dogid:      dog.Id,
			HealthType: "HHRT",
			CertId:     toller.HeartClear,
		}

		err = model.HealthModel.Insert(health)
		if err != nil {
			return err
		}
	}

	//elbows
	if toller.ElbowClear != "" {
		health := models.Health{
			Dogid:      dog.Id,
			HealthType: "HELB",
			CertId:     toller.ElbowClear,
		}

		err = model.HealthModel.Insert(health)
		if err != nil {
			return err
		}
	}

	return nil
}
func yesNoToBool(yesNo string) bool {
	return yesNo == "Y"
}

func (model *MigrateModel) importSiresAndDams(tollers []*models.Toller) error {

	male := model.CodetableModel.GetByCode("D", models.CategorySex)
	female := model.CodetableModel.GetByCode("B", models.CategorySex)

	sires := map[string]bool{}
	dams := map[string]bool{}
	log.Println("total tollers:", len(tollers))
	for _, currentToller := range tollers {
		if !sires[strings.ToUpper(currentToller.Sirename)] {
			currentSire := models.Dog{}
			currentSire.Regnum = currentToller.Sireregnum
			if currentSire.Regnum == "" {
				currentSire.Regnum = "UNKNOWN"
			}
			currentSire.Name = currentToller.Sirename
			currentSire.Sex = male.Id
			currentSire.Litterid = 1

			err := model.DogModel.Insert(currentSire)
			if err != nil {
				return err
			}

			sires[strings.ToUpper(currentToller.Sirename)] = true
		}

		if !dams[strings.ToUpper(currentToller.Damname)] {
			currentDam := models.Dog{}
			currentDam.Regnum = currentToller.Damregnum
			if currentDam.Regnum == "" {
				currentDam.Regnum = "UNKNOWN"
			}
			currentDam.Name = currentToller.Damname
			currentDam.Sex = female.Id
			currentDam.Litterid = 1

			err := model.DogModel.Insert(currentDam)
			if err != nil {
				return err
			}
			dams[strings.ToUpper(currentToller.Damname)] = true
		}

	}

	log.Println("sires inserted:", len(sires))
	log.Println("dams inserted:", len(dams))

	return nil
}
