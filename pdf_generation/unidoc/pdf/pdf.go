package pdf

import (
  "fmt"
  "os"

  //unidocCore "github.com/unidoc/unidoc/pdf/core"
  unidoc "github.com/unidoc/unidoc/pdf/model"
)

func FillPPM(inputName string, outputName string) error {
	f, err := os.Open(inputName)
	if err != nil {
		return err
	}

	defer f.Close()

	pdfReader, err := unidoc.NewPdfReader(f)
	if err != nil {
		return err
	}

	acroForm := pdfReader.AcroForm
	if acroForm == nil {
		fmt.Printf(" No formdata present\n")
		return nil
	}
/*

	data := map[string]string{
		"Date_Prepared":       "2018-03-01",
		"Last, First, Middle": "Chen, Aileen",
		"Phone Number":        "123-321-3212",
	}

	for i, field := range *acroForm.Fields {
		value, ok := data[field.T.String()]
		if ok {
			fmt.Println(value)
			(*acroForm.Fields)[i].V = unidocCore.MakeString(value)
		}

		(*acroForm.Fields)[i].KidsF = []unidoc.PdfModel{}
		//(*acroForm.Fields)[i].KidsA = []*unidoc.PdfAnnotation{}

		       for _, child := range field.KidsF {
		   			switch c := child.(type) {
		   			case *unidoc.PdfField:
		           childField := child.(*unidoc.PdfField)
		           if childField.T != nil {
		             value, ok := data[childField.T.String()]
		             if ok {
		               fmt.Println(value)
		               childField.V = unidocCore.MakeString(value)
		             }
		           }
		   			case *unidoc.PdfAnnotationWidget:
		   				fmt.Printf(" --Widget: %+v\n", *c)
		   			default:
		   				fmt.Printf(" f--UNKNOWN\n")
		   			}
		   		}
/*
		for _, child := range field.KidsA {
			child.PdfObject
			if child.T != nil {
				value, ok := data[child.T.String()]
				if ok {
					fmt.Println(value)
					child.V = unidocCore.MakeString(value)
				}
			}
		}

	}

	fmt.Printf("before writing\n")
	for idx, field := range *acroForm.Fields {
		if idx > 3 {
			break
		}
		fmt.Printf(" -Field %d (%p): %+v\n", idx+1, field, *field)
	}
*/
	pdfWriter := unidoc.NewPdfWriter()

	numPages, err := pdfReader.GetNumPages()
	if err != nil {
		return err
	}

	for i := 1; i < numPages+1; i++ {
		page, err := pdfReader.GetPage(i)
		if err != nil {
			return err
		}

		err = pdfWriter.AddPage(page)
		if err != nil {
			return err
		}
	}

	pdfWriter.SetForms(acroForm)

	fWrite, err := os.Create(outputName)
	if err != nil {
		return err
	}

	defer fWrite.Close()

	err = pdfWriter.Write(fWrite)
	if err != nil {
		return err
	}

	return nil
}

func ReadPDF(fileName string) error {
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}

	defer f.Close()

	pdfReader, err := unidoc.NewPdfReader(f)
	if err != nil {
		return err
	}

	acroForm := pdfReader.AcroForm
	if acroForm == nil {
		fmt.Printf(" No formdata present\n")
		return nil
	}

	fmt.Printf("reading\n")
	for idx, field := range *acroForm.Fields {
		if idx > 3 {
			break
		}
		fmt.Printf(" -Field %d (%p): %+v\n", idx+1, field, *field)
		for _, child := range field.KidsF {
			switch c := child.(type) {
			case *unidoc.PdfField:
				fmt.Printf(" --Field: %+v\n", *c)
			case *unidoc.PdfAnnotationWidget:
				fmt.Printf(" --Widget: %+v\n", *c)
			default:
				fmt.Printf(" f--UNKNOWN\n")
			}
		}
	}
	return nil
}
