package main

import "fmt"

type FullTimeEmployee struct {
	Id           int
	NationalCode string
	FullName     string
	ExtraHours   float64
}

type PartTimeEmployee struct {
	Id           int
	NationalCode string
	FullName     string
	PartHours    float64
}

type ShiftTimeEmployee struct {
	Id           int
	NationalCode string
	FullName     string
	ShiftHours   float64
}

const (
	BaseSalary       = 80000
	ExtraHourRate    = 125459
	HourlySalaryRate = 547878
	ShiftSalaryRate  = 12459974
	TaxRate          = 0.09
)

type Employee interface {
	SalaryCalculator(hours float64) (salary float64, tax float64)
}

// if We Need To Use Encapsulation .. Write Camel Case Function Or ...

func (employee FullTimeEmployee) SalaryCalculator(extraHours float64) (salary float64, tax float64) {
	salary = BaseSalary + (ExtraHourRate*extraHours)*1.4
	tax = salary * TaxRate
	salary += tax

	return
}

func (employee PartTimeEmployee) SalaryCalculator(hours float64) (salary float64, tax float64) {
	salary = HourlySalaryRate * hours
	tax = salary * TaxRate
	salary += tax

	return
}

func main() {

	fullTimePerson1 := FullTimeEmployee{
		Id:           1,
		ExtraHours:   234.2,
		FullName:     "pouya",
		NationalCode: "3432",
	}

	fullTimeSalary, taxSalary := fullTimePerson1.SalaryCalculator(fullTimePerson1.ExtraHours)
	fmt.Printf("Full time salary is: %.2f and Tax is: %v", fullTimeSalary, taxSalary)

	partTimeEmployee := PartTimeEmployee{
		Id:           2,
		FullName:     "hossein",
		NationalCode: "1213213",
		PartHours:    124.4,
	}

	salaryPartTime, TaxPartTime := partTimeEmployee.SalaryCalculator(partTimeEmployee.PartHours)
	fmt.Printf("Part time salary is: %.2f and Tax is: %v", salaryPartTime, TaxPartTime)

}
