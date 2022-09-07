package domains

import "github.com/chocogem/keycloak-authen-golang-simple-clean-arch/models"

type ITestRepository interface {
	Get(id int) (models.Test, error)
}
