package dataset

import (
	"github.com/rai-project/plini/pkg/instruction"
	"github.com/rai-project/uuid"
)

type Dataset interface {
	ID() string
	Split(int) Dataset
	Apply(instruction.Instruction) Dataset
	Run() Dataset
}

type dataset struct {
	id           string
	data         []int
	size         int
	instructions []instruction.Instruction
}

type datasets struct {
	id       string
	datasets []*dataset
}

func (d *dataset) ID() string {
	return d.id
}

func (d *dataset) Apply(f instruction.Instruction) Dataset {
	return &dataset{
		id:           uuid.NewV4(),
		data:         d.data,
		size:         d.size,
		instructions: append(d.instructions, f),
	}
}

func (d *dataset) Run() Dataset {
	return nil
}

func (d *dataset) Split(n int) Dataset {
	ds := make([]*dataset, n)
	for ii := 0; ii < n; ii++ {
		sz := min(ii*n, d.size)
		ds[ii] = &dataset{
			id:           d.id, // TODO:: unsure if this is a good idea
			data:         d.data[ii*n : ii*n+sz],
			size:         sz,
			instructions: d.instructions,
		}
	}
	return &datasets{
		id:       uuid.NewV4(),
		datasets: ds,
	}
}

func (d *datasets) ID() string {
	return d.id
}

func (d *datasets) Apply(f instruction.Instruction) Dataset {
	return nil
}

func (d *datasets) Split(n int) Dataset {
	return nil
}

func (d *datasets) Run() Dataset {
	return nil
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
