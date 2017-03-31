package inference

var genSym = make(chan int)

func init() {
	go func() {
		ii := 0
		for {
			genSym <- ii
			ii++
		}
	}()
}
