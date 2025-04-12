package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"services-api/models"
	"services-api/utils"

	"github.com/google/uuid"
)

// In-memory storage can be replace with DB in prod
var Services []models.Service
var ServiceVersions []models.ServiceVersion

func InitializeData() {
	Services = []models.Service{
		{ID: "1", Name: "Priority Services", Description: "Priority service for user management", Version: 1},
		{ID: "2", Name: "Security", Description: "Service for managing user security", Version: 2},
		{ID: "3", Name: "Reporting", Description: "Service to handle reportings", Version: 3},
		{ID: "4", Name: "Analytics", Description: "Service for data analytics", Version: 4},
		{ID: "5", Name: "Monitoring", Description: "Service for system monitoring", Version: 5},
	}

	ServiceVersions = []models.ServiceVersion{
		{ServiceID: "1", Version: 1},
		{ServiceID: "2", Version: 1},
		{ServiceID: "2", Version: 2},
		{ServiceID: "3", Version: 1},
		{ServiceID: "3", Version: 2},
		{ServiceID: "3", Version: 3},
		{ServiceID: "4", Version: 1},
		{ServiceID: "4", Version: 2},
		{ServiceID: "4", Version: 3},
		{ServiceID: "4", Version: 4},
		{ServiceID: "5", Version: 1},
		{ServiceID: "5", Version: 2},
		{ServiceID: "5", Version: 3},
		{ServiceID: "5", Version: 4},
		{ServiceID: "5", Version: 5},
	}
}

// CreateService creates a new service
func CreateService(w http.ResponseWriter, r *http.Request) {
	var service models.Service
	if err := json.NewDecoder(r.Body).Decode(&service); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	service.ID = uuid.New().String()
	Services = append(Services, service)
	utils.RespondWithJSON(w, http.StatusCreated, service)
}

// GetServices retrieves all services with filtering, sorting, and pagination
func GetServices(w http.ResponseWriter, r *http.Request) {
	// Default values set to 1 and 10
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	if page == 0 {
		page = 1
	}
	if limit == 0 {
		limit = 10
	}

	// Filtering
	filter := strings.TrimSpace(strings.ToLower(r.URL.Query().Get("filter")))

	// Sorting
	sortBy := strings.TrimSpace(r.URL.Query().Get("sort_by"))
	sortOrder := strings.TrimSpace(strings.ToLower(r.URL.Query().Get("sort_order")))
	if sortOrder == "" {
		sortOrder = "asc"
	}

	// Apply filtering
	filteredServices := applyFiltering(Services, filter)

	// Apply sorting
	sortedServices := applySorting(filteredServices, sortBy, sortOrder)

	// Apply pagination
	paginatedServices, totalPages := applyPagination(sortedServices, page, limit)

	// Prepare response
	response := map[string]interface{}{
		"page":        page,
		"limit":       limit,
		"total":       len(sortedServices),
		"total_pages": totalPages,
		"services":    paginatedServices,
	}

	utils.RespondWithJSON(w, http.StatusOK, response)
}

// GetService by id
func GetService(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/api/services/")

	for _, item := range Services {
		if item.ID == id {
			utils.RespondWithJSON(w, http.StatusOK, item)
			return
		}
	}
	utils.RespondWithError(w, http.StatusNotFound, "Service not found")
}

// UpdateService by id
func UpdateService(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/api/services/")

	var updatedService models.Service
	if err := json.NewDecoder(r.Body).Decode(&updatedService); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	for index, item := range Services {
		if item.ID == id {
			updatedService.ID = item.ID
			Services[index] = updatedService
			utils.RespondWithJSON(w, http.StatusOK, updatedService)
			return
		}
	}
	utils.RespondWithError(w, http.StatusNotFound, "Service not found")
}

// DeleteService by id
func DeleteService(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/api/services/")

	for index, item := range Services {
		if item.ID == id {
			Services = append(Services[:index], Services[index+1:]...)
			utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "Service deleted successfully"})
			return
		}
	}
	utils.RespondWithError(w, http.StatusNotFound, "Service not found")
}

// GetServiceVersions get service versions by ServiceID
func GetServiceVersions(w http.ResponseWriter, r *http.Request) {
	serviceID := strings.TrimPrefix(r.URL.Path, "/api/services/")
	serviceID = strings.TrimSuffix(serviceID, "/versions")

	var versions []models.ServiceVersion
	for _, version := range ServiceVersions {
		if version.ServiceID == serviceID {
			versions = append(versions, version)
		}
	}

	if len(versions) == 0 {
		utils.RespondWithError(w, http.StatusNotFound, "No versions found for this service")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, versions)
}

// GetServiceByName get a service by name
func GetServiceByName(w http.ResponseWriter, r *http.Request) {
	name := strings.TrimPrefix(r.URL.Path, "/api/services/name/")
	name = strings.TrimSpace(strings.ToLower(name))

	for _, item := range Services {
		if strings.TrimSpace(strings.ToLower(item.Name)) == name {
			utils.RespondWithJSON(w, http.StatusOK, item)
			return
		}
	}
	utils.RespondWithError(w, http.StatusNotFound, "Service not found")
}

// Apply filtering
func applyFiltering(services []models.Service, filter string) []models.Service {
	if filter == "" {
		return services
	}

	filtered := []models.Service{}
	for _, service := range services {
		if strings.Contains(strings.ToLower(service.Name), filter) ||
			strings.Contains(strings.ToLower(service.Description), filter) {
			filtered = append(filtered, service)
		}
	}
	return filtered
}

// Apply sorting
func applySorting(services []models.Service, sortBy string, sortOrder string) []models.Service {
	if sortBy == "" {
		return services
	}

	less := func(i, j int) bool {
		if sortBy == "name" {
			if sortOrder == "asc" {
				return strings.ToLower(services[i].Name) < strings.ToLower(services[j].Name)
			}
			return strings.ToLower(services[i].Name) > strings.ToLower(services[j].Name)
		}
		return true
	}

	SortServices(services, less)
	return services
}

// SortServices function to sort services
func SortServices(services []models.Service, less func(i, j int) bool) {
	for i := 0; i < len(services); i++ {
		for j := i + 1; j < len(services); j++ {
			if less(i, j) {
				services[i], services[j] = services[j], services[i]
			}
		}
	}
}

// Apply pagination
func applyPagination(services []models.Service, page int, limit int) ([]models.Service, int) {
	start := (page - 1) * limit
	if start >= len(services) {
		return []models.Service{}, 0
	}

	end := start + limit
	if end > len(services) {
		end = len(services)
	}

	totalPages := (len(services) + limit - 1) / limit
	return services[start:end], totalPages
}
