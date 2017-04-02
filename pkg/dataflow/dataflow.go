package dataflow

import "github.com/rai-project/plini/pkg/dataset"

type DataFlow struct {
	err error
}

func New() *dataset.DataFlowContext {
	return &dataset.DataFlowContext{}
}

// func (df *dataflow) Input(locs ...string) *dataflow {
// 	for _, loc := range locs {
// 		fs.
// 	}
// }

func (df *DataFlow) Error() string {
	return df.err.Error()
}
