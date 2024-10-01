package constructor_logic_entrypoint

func (w *LogicWorker) InsertTask(dir Directive) *DirResp {
	return &DirResp{
		Err:  nil,
		Data: dir.Data,
	}
}

func (w *LogicWorker) DeleteTask(dir Directive) *DirResp {
	return &DirResp{
		Err:  nil,
		Data: dir.Data,
	}
}

func (w *LogicWorker) TXTask(dir Directive) *DirResp {
	return &DirResp{
		Err:  nil,
		Data: dir.Data,
	}
}
