package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/mfrw/dumbfioparser"
)

var (
	inDir = flag.String("i", "", "Input Directory")
	hdr   = flag.Bool("h", false, "Print the header")
)

func main() {
	flag.Parse()
	if *inDir == "" {
		log.Fatal("Please provide an input directory")
	}
	if *hdr {
		fmt.Println("JobName BlockSize Rd-bw Wr-bw Rd-Iops Wr-Iops, Avg-Rd-lat, Avg-Wr-lat")
	}
	SlurpDir(*inDir)
}

func SlurpDir(dir string) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		err = ParseSingleJson(filepath.Join(dir, f.Name()))
		if err != nil {
			log.Println("Failed to parse:", f.Name())
		}
	}
}

func ParseSingleJson(filePath string) error {
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	fioData, err := SlurpFioOut(f)
	if err != nil {
		return err
	}
	PrintFancy(&fioData)
	return nil
}

func PrintFancy(f *dumbfioparser.Fio) {
	for i := range f.Jobs {
		fmt.Printf("%s %s %f %f %f %f %f %f\n", f.Jobs[i].Jobname, f.Global_options.Bs, f.Jobs[i].Read.Bw, f.Jobs[i].Write.Bw, f.Jobs[i].Read.Iops, f.Jobs[i].Write.Iops, f.Jobs[i].Read.LatNs.Mean, f.Jobs[i].Write.LatNs.Mean)
	}
}

func SlurpFioOut(r io.Reader) (dumbfioparser.Fio, error) {
	f := dumbfioparser.Fio{}
	err := json.NewDecoder(r).Decode(&f)
	return f, err
}
