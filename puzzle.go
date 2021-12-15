package common

import "log"

const (
	sampleFileName = "sample"
	inputFileName  = "input"
)

type Puzzle struct {
	Sample []byte
	Input  []byte
}

func Load() (p *Puzzle, err error) {
	p = &Puzzle{}
	if p.Sample, err = Ingest(sampleFileName); err != nil {
		return nil, err
	}
	if p.Input, err = Ingest(inputFileName); err != nil {
		return nil, err
	}
	return p, nil
}

type Solver func(b []byte) (string, error)

func Solve(p *Puzzle, solvers ...Solver) {
	for i, solver := range solvers {
		sample, err := solver(p.Sample)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("%d.sample\t%s", i+1, sample)
		input, err := solver(p.Input)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("%d.input\t%s", i+1, input)
	}
}
