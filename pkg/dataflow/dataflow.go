package dataflow

import (
	"context"

	"github.com/rai-project/plini/pkg/dataset"
	"github.com/rai-project/plini/pkg/step"
)

type DataFlow struct {
	ctx    context.Context
	levels []Level
	err    error
}

func New() *dataset.DataFlowContext {
	return &dataset.DataFlowContext{
		ctx:    context.Background(),
		levels: []Level{},
		err:    nil,
	}
}

func (d *DataFlow) Then(steps ...step.Step) *DataFlow {
	if len(d.levels) == 0 {
		d.level = append(d.levels, levels)
		return d
	}
	lastLevel := d.levels[len(d.levels)-1]
	for ii := range steps {
		steps[ii].SetDepends(lastLevel.ID())
	}
	d.level = append(d.levels, levels)
	return d
}

func (d *DataFlow) Run(steps ...step.Step) error {
	for _, level := range d.levels {
		for _, step := range level.steps {
			err := step.Run()
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (df *DataFlow) Error() string {
	return df.err.Error()
}
