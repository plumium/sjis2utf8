package main

import (
	"fmt"
	"io"
	"os"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

func main() {
	args := os.Args
	if len(args) != 3 {
		fmt.Printf("Usage: %s <input_file_path> <output_file_path>\n", os.Args[0])
		return
	}

	in, err := os.Open(args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	defer in.Close()

	out, err := os.Create(args[2])
	if err != nil {
		fmt.Println(err)
		return
	}
	defer out.Close()

	err = Conversion(in, out)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Conversion successfull")
}

func Conversion(in io.Reader, out io.Writer) error {
	r := transform.NewReader(in, japanese.ShiftJIS.NewDecoder())
	_, err := io.Copy(out, r)
	if err != nil {
		return err
	}
	return nil
}
