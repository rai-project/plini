package dataset

import "io"

func (df *DataFlowContext) Input(f func(io.Writer) error) *Dataset {
}

func (df *DataFlowContext) Textfile(path string) *Dataset {
	return nil
}
