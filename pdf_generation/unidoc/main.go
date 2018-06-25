package main

import (
  "fmt"
  "pdf/pdf"
)

func main() {
  err := pdf.FillPPM("PPM_Form.pdf", "ppm.pdf")
  if err != nil {
    fmt.Printf("error: %#v\n", err)
  }
  err = pdf.ReadPDF("ppm.pdf")
  if err != nil {
    fmt.Printf("error: %#v\n", err)
  }
}
