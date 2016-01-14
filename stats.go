package neat

type Stats struct {
	Ok       int
	Changed  int
	Deferred int
	Skipped  int
}

func (s *Stats) Record(status ModuleStatus) {
	switch status {
	case ModuleOk:
		s.Ok++
	case ModuleChanged:
		s.Changed++
	case ModuleDeferred:
		s.Deferred++
	case ModuleSkipped:
		s.Skipped++
	}
}
