# gofpdf

This directory contains two samples. To run either, first install the prerequisites using `go get
github.com/jung-kurt/gofpdf`.

## generate

This is an example of how to generate a PDF using the [gofpdf](https://github.com/jung-kurt/gofpdf) library. Run using `go run main.go`. The generated file will be saved as `ppm.form`.

## image_to_pdf

This program wraps one or more images in a PDF container for ease in merging with other PDFs in a document processing pipeline. Run using  `go run main.go -input orders-p1.jpg -input orders-p2.jpg -output orders.pdf`.

# unidoc

A very preliminary spike on getting [unidoc](https://github.com/unidoc/unidoc) up and running to fill in a PDF form.

# pypdf2

This script successfully fills in a PDF form using the [PyPDF2](https://github.com/mstamy2/PyPDF2) library. To use, run `pip install PyPDF2` using Python 3. Then run `cd` into `pdf_generation/pypdf2` and run `python fill_ppm_form.py PPM_Form_.pdf out.pdf`.

# pdfcpu

This file uses []() to merge multiple PDFs into a single PDF. The process is simple and retains the original quality of the input files; another way to look at this is that no additional compression is done during the process. Install prerequisites using `go get github.com/hhrutter/pdfcpu/pkg/api` and then run with `go run main.go -input form.pdf -input orders.pdf -output merged.pdf`.
