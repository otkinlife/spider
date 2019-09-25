package queue

import (
	"../doSite"
	"container/list"
	"errors"
)

type Queue struct {
	l        *list.List
	limitLen int
}

//初始化队列
func New() *Queue {
	return &Queue{
		l:        list.New(),
		limitLen: 0,
	}
}

//入队列
func (q *Queue) PutIn(obj doSite.SiteDownload) error {
	if q.l.Len() >= q.limitLen {
		err := errors.New("队列已满")
		return err
	}
	e := list.Element{
		Value: obj,
	}
	q.l.PushBack(e)
	return nil
}

//出队列
func (q *Queue) PutOut() doSite.SiteDownload {
	if q.l.Len() == 0 {
		return nil
	}
	outRes := q.l.Front()
	q.l.Remove(q.l.Front())
	if v, ok := outRes.Value.(doSite.SiteDownload); ok {
		return v
	}
	return nil
}
