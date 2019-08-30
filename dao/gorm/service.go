/*
@Time : 2019-07-09 14:27
@Author : zr
*/
package gorm

import (
	"camdig/server/errors"
	"camdig/server/model"
	"github.com/sirupsen/logrus"
	"strings"
)

type ServiceDAO struct {
	*BaseSqlDAO
}

func GetServiceDAO() ServiceDAO {
	m := ServiceDAO{
		BaseSqlDAO: newBase(),
	}
	return m
}
func (s *ServiceDAO) GetRolesServiceNecessary(serviceUrl string) (roles []model.Role, err error) {
	services, err := s.getAllServices()
	if err != nil {
		return roles, err
	}
	servicesNeed := s.findServicesNeed(services, serviceUrl)
	if len(servicesNeed) == 0 {
		return roles, err
	}
	roles = s.findRolesNeeded(servicesNeed)
	return roles, nil
}

func (s *ServiceDAO) getAllServices() (services []model.Service, err error) {
	if err := s.getDB().Find(&services).Error; err != nil {
		err = errors.DatabaseError(err)
	}
	return services, err
}

func (s *ServiceDAO) findServicesNeed(services []model.Service, serviceUrl string) (servicesNeed []model.Service) {
	serverStrs := strings.Split(serviceUrl, "/")
	for _, serviceStr := range serverStrs {
		found := false
		for _, service := range services {
			if serviceStr == service.Name {
				servicesNeed = append(servicesNeed, service)
				found = true
				break
			}
		}
		if !found {
			break
		}
	}
	return servicesNeed
}

func appendRoleNoDuplication(roles []model.Role, role model.Role) []model.Role {
	for _, v := range roles {
		if v.Id == role.Id {
			return roles
		}
	}
	return append(roles, role)
}

func (s *ServiceDAO) findRolesNeeded(services []model.Service) (roles []model.Role) {
	for _, service := range services {
		var roles_ []model.Role
		if err := s.getDB().Model(&service).Related(&roles_, "Roles").Error; err != nil {
			logrus.Warn(err)
		}

		for _, role := range roles_ {
			roles = appendRoleNoDuplication(roles, role)
		}
	}
	return roles
}
