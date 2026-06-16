package main

import (
	"fmt"
	"net/http"
)

type Department interface {
	GetInfo() string
}

type Tech struct {
	DeptName string
	Phone    string
}

type Billing struct {
	DeptName string
	Phone    string
}

func (t Tech) GetInfo() string {
	return "Welcome to " + t.DeptName + ". Call line: " + t.Phone
}

func (b Billing) GetInfo() string {
	return "Welcome to " + b.DeptName + ". Call line: " + b.Phone
}

type RouterHandler struct {
	Target Department
}

func (r RouterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, r.Target.GetInfo())
}

func main() {
	techDept := Tech{DeptName: "Tech Support", Phone: "101"}
	billingDept := Billing{DeptName: "Billing Department", Phone: "102"}

	techPage := RouterHandler{Target: techDept}
	billingPage := RouterHandler{Target: billingDept}

	http.Handle("/tech", &techPage)
	http.Handle("/billing", &billingPage)

	http.ListenAndServe(":8080", nil)
}
