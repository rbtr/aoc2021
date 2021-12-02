package common

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
