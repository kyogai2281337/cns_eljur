package constructor_logic_entrypoint

import (
	"fmt"
)

func (w *LogicWorker) InsertTask(dir Directive, schedule *CacheItem) *DirResp {
	insertReq := dir.Data.(UpdateInsertRequest).Data

	lectureData := struct {
		Groups  []string
		Teacher string
		Cabinet string
		Subject string
	}{
		Groups:  insertReq.Lecture.Groups,
		Teacher: insertReq.Lecture.Teacher,
		Cabinet: insertReq.Lecture.Cabinet,
		Subject: insertReq.Lecture.Subject,
	}
	schedule.mu.RLock()
	defer schedule.mu.RUnlock()
	lecData, err := schedule.Schedule.RecoverLectureData(&lectureData)
	if err != nil {
		return &DirResp{
			Err:  err,
			Data: dir.ScheduleID,
		}
	}
	if err := schedule.Schedule.Insert(insertReq.Day, insertReq.Pair, lecData); err != nil {
		return &DirResp{
			Err:  err,
			Data: dir.ScheduleID,
		}
	}

	return &DirResp{
		Err:  nil,
		Data: dir.ScheduleID,
	}
}

func (w *LogicWorker) DeleteTask(dir Directive, schedule *CacheItem) *DirResp {
	//Parsing request
	req := UpdateDeleteRequest{}.Data
	req = dir.Data.(UpdateDeleteRequest).Data

	schedule.mu.RLock()
	defer schedule.mu.RUnlock()
	// Doing case
	if err := schedule.Schedule.Delete(req.Day, req.Pair, schedule.Schedule.RecoverObject(req.Name, req.Type)); err != nil {
		return &DirResp{
			Err:  err,
			Data: dir.ScheduleID,
		}
	}

	return &DirResp{
		Err:  nil,
		Data: dir.ScheduleID,
	}
}

func (w *LogicWorker) RenameTask(dir Directive, schedule *CacheItem) *DirResp {
	//Parsing request
	req := UpdateRenameRequest{}.Data
	req = dir.Data.(UpdateRenameRequest).Data

	schedule.mu.RLock()
	defer schedule.mu.RUnlock()
	// Doing case
	if err := schedule.Schedule.Rename(req.Name); err != nil {
		return &DirResp{
			Err:  err,
			Data: dir.ScheduleID,
		}
	}

	return &DirResp{
		Err:  nil,
		Data: dir.ScheduleID,
	}
}

func (w *LogicWorker) TXTask(dir Directive, sch *CacheItem) *DirResp {
	dirArr := dir.Data.(UpdateTXRequest).Data

	for idx, directive := range dirArr {
		switch directive.Type {
		case DirInsert:
			resp := w.InsertTask(directive, sch)
			if resp.Err != nil {
				return &DirResp{
					Data: resp.Data,
					Err:  fmt.Errorf("instruction %d: %s", idx, resp.Err.Error()),
				}
			}
		case DirDelete:
			resp := w.DeleteTask(directive, sch)
			if resp.Err != nil {
				return &DirResp{
					Data: resp.Data,
					Err:  fmt.Errorf("instruction %d: %s", idx, resp.Err.Error()),
				}
			}
		case DirRename:
			resp := w.RenameTask(directive, sch)
			if resp.Err != nil {
				return &DirResp{
					Data: resp.Data,
					Err:  fmt.Errorf("instruction %d: %s", idx, resp.Err.Error()),
				}
			}
		default:
			return &DirResp{
				Err: fmt.Errorf("unknown instruction type in instruction %d: %d", idx, directive.Type),
			}
		}
	}
	sch.mu.RLock()
	defer sch.mu.RUnlock()
	err := sch.Schedule.MakeReview()
	if err != nil {
		return &DirResp{
			Err:  err,
			Data: dir.ScheduleID,
		}
	}

	return &DirResp{
		Err:  nil,
		Data: dir.ScheduleID,
	}
}
