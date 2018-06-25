package main

import (
	"./pdf"
	"fmt"
)

func main() {
	/*
	  err := pdf.FillPPM("PPM_Form.pdf", "ppm.pdf")
	  if err != nil {
	    fmt.Printf("error: %#v\n", err)
	  }
	  err = pdf.ReadPDF("ppm.pdf")
	  if err != nil {
	    fmt.Printf("error: %#v\n", err)
	  }
	*/
	data := pdf.PPMData{
		Name:    "Lastname, Firstname",
		Phone:   "123-456-7890",
		DODID:   "123456789",
		Branch:  "Air Force",
		Rank:    "E-10",
		Email:   "user@example.com",
		Address: "1234 Main St, Washington DC 20001",
	}
	form := pdf.NewMyMoveForm(data)
	err := form.GeneratePPM()
	if err != nil {
		fmt.Printf("error: %#v\n", err)
	}
}
