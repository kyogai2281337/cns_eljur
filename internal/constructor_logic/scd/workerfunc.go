package constructor_logic_entrypoint

import (
	"github.com/gofiber/fiber/v2/log"
)

func (w *LogicWorker) InsertTask(dir Directive, schedule *CacheItem) *DirResp {
	req := UpdateInsertRequest{}
	req = dir.Data.(UpdateInsertRequest)

	if err := schedule.Schedule.Insert(req.Day, req.Pair, schedule.Schedule.RecoverLectureData(&struct {
		Groups  []string
		Teacher string
		Cabinet string
		Subject string
	}{
		Groups:  req.Lecture.Groups,
		Teacher: req.Lecture.Teacher,
		Cabinet: req.Lecture.Cabinet,
		Subject: req.Lecture.Subject,
	})); err != nil {
		return &DirResp{
			Err:  err,
			Data: nil,
		}
	}

	return &DirResp{
		Err:  nil,
		Data: dir.Data,
	}
}

func (w *LogicWorker) DeleteTask(dir Directive, schedule *CacheItem) *DirResp {
	//Parsing request
	req := UpdateDeleteRequest{}
	req = dir.Data.(UpdateDeleteRequest)

	// Doing case
	if err := schedule.Schedule.Delete(req.Day, req.Pair, schedule.Schedule.RecoverObject(req.Name, req.Type)); err != nil {
		return &DirResp{
			Err:  err,
			Data: nil,
		}
	}

	return &DirResp{
		Err:  nil,
		Data: dir.Data,
	}
}

func (w *LogicWorker) TXTask(dir Directive, sch *CacheItem) *DirResp {
	dirArr := dir.Data.([]Directive)

	for _, directive := range dirArr {
		switch directive.Type {
		case DirInsert:
			resp := w.InsertTask(directive, sch)
			if resp.Err != nil {
				return resp
			}
		case DirDelete:
			resp := w.DeleteTask(directive, sch)
			if resp.Err != nil {
				return resp
			}
		default:
			log.Error("Unknown directive type: ", directive.Type)
		}
	}
	err := sch.Schedule.MakeReview()
	if err != nil {
		return &DirResp{
			Err:  err,
			Data: nil,
		}
	}

	return &DirResp{
		Err:  nil,
		Data: dir.ID,
	}
}
