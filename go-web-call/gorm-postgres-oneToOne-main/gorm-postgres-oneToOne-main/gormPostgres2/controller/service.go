package controller

import (
	"demo/connection"
	"demo/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func CreateEmployee(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	db := connection.DatabaseConnection()
	var Employee models.Employee
	json.NewDecoder(request.Body).Decode(&Employee)
	//var department = Employee.Department
	//db.Create(&department)
	//Employee.DepartmentId = department.ID
	db.Create(&Employee)
	json.NewEncoder(response).Encode(Employee.ID)
	return

}
func GetAllEmployees(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	db := connection.DatabaseConnection()
	var Employees []models.Employee
	data := db.Model(models.Employee{}).Preload("Department").Find(&Employees)
	log.Println(data)
	json.NewEncoder(response).Encode(Employees)
	return

}
func GetEmployeeById(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	db := connection.DatabaseConnection()
	params := mux.Vars(request)
	var Employee models.Employee
	db.Model(models.Employee{}).Preload("Department").First(&Employee, params["employeeId"])
	json.NewEncoder(response).Encode(Employee)
	return

}
func UpdateEmployee(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	db := connection.DatabaseConnection()
	params := mux.Vars(request)
	var Employee models.Employee
	db.First(&Employee, params["employeeId"])
	var department models.Department
	department.ID = Employee.DepartmentId
	json.NewDecoder(request.Body).Decode(&Employee)
	department = Employee.Department
	db.Save(&department)
	db.Save(&Employee)
	json.NewEncoder(response).Encode(Employee)
	return

}
func DeleteEmployeeById(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	db := connection.DatabaseConnection()
	params := mux.Vars(request)
	var Employee models.Employee
	db.Model(models.Employee{}).Preload("Department").First(&Employee, params["employeeId"])
	//db.First(&Employee, params["employeeId"])
	var Department = Employee.Department
	db.Delete(&Department)
	db.Delete(&Employee)
	json.NewEncoder(response).Encode("deleted successfully")
	return

}
