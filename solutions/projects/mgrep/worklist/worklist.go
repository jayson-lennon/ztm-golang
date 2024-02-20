package worklist

// Empty path indicates that there are no more jobs to be done
type Entry struct {
	Path string
}

type Worklist struct {
	jobs chan Entry
}

func (w *Worklist) Add(work Entry) {
	w.jobs <- work
}

func (w *Worklist) Next() Entry {
	j := <-w.jobs
	return j
}

func New(bufSize int) Worklist {
	return Worklist{make(chan Entry, bufSize)}
}

func NewJob(path string) Entry {
	return Entry{path}
}

// We add a `NoMoreJobs` message to the worklist for each
// worker that we are using. Once the worker receives this
// message, it will terminate. After each worker terminates,
// the program can continue.
func (w *Worklist) Finalize(numWorkers int) {
	for i := 0; i < numWorkers; i++ {
		w.Add(Entry{""})
	}
}
