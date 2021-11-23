package main

import (
	"testing"
)

func TestGetFullTimeEmployeeById(t *testing.T){
	table := [] struct {
		id int
		dni string
		mockFunc func()
		expectedEmployee FullTimeEmployee
	}{
		{ 
			id: 1 ,
			dni : "1",
			mockFunc: func(){
				GetEmployeeeById = func (id int ) (Employee, error){
					return Employee{
						Id: 1,
						Position: "Ceo",
					}, nil
				}
				GetPersonByDNI = func (id string) (Person , error){
					return Person {
						Name: "Jhon Doe",
						Age: 35,
						DNI: "1",
					}, nil
				}
			},
			expectedEmployee: FullTimeEmployee {
				 Person:  Person {
					 Age: 35,
					 DNI: "1",
					 Name: "Jhon Doe",
				 },
				 Employee: Employee{
					 Id: 1, 
					 Position: "CEO",
				 },
			},
		},
	}

	originalGetEmployeeById := GetEmployeeeById
	originalGetPersonByDni := GetPersonByDNI


	for _ , test := range  table{
			test.mockFunc()
			ft, err := GetFullTimeEmployeeById(test.id, test.dni)

			if err != nil {
				t.Errorf("Error when  getting Employee")
			}

			if ft.Age != test.expectedEmployee.Age{
					t.Errorf("Error, got %d expected %d", ft.Age, test.expectedEmployee.Age)
			}
	}

	GetEmployeeeById = originalGetEmployeeById
	GetPersonByDNI = originalGetPersonByDni

}

