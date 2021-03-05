package pkg

import (
	cro "github.com/robfig/cron/v3"
	"sync"
)

var (
	CronManager *Crontab
	once        sync.Once
)

func init() {
	once.Do(func() {
		CronManager = NewCrontab()
	})
}

// Crontab crontab manager
type Crontab struct {
	inner *cro.Cron
	ids   map[string]cro.EntryID
	mutex sync.Mutex
}

// Crontab 构造函数
func NewCrontab() *Crontab {
	return &Crontab{
		inner: cro.New(),
		ids:   make(map[string]cro.EntryID),
	}

}

// IDs ... 返回ID列表
func (c *Crontab) IDs() []string {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	validIDs := make([]string, 0, len(c.ids))
	invalidIDs := make([]string, 0)
	for sid, eid := range c.ids {
		if e := c.inner.Entry(eid); e.ID != eid {
			invalidIDs = append(invalidIDs, sid)
			continue
		}
		validIDs = append(validIDs, sid)
	}
	for _, id := range invalidIDs {
		delete(c.ids, id)
	}
	return validIDs
}

// 启动Crontab引擎
func (c *Crontab) Start() {
	c.inner.Start()
}

// 停止crontab引擎
func (c *Crontab) Stop() {
	c.inner.Stop()
}

// 通过ID移除定时任务
func (c *Crontab) DelByID(id string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	eid, ok := c.ids[id]
	if !ok {
		return
	}
	c.inner.Remove(eid)
	delete(c.ids, id)
}

// 添加一个作业
// id是唯一的
// spec 是表达式 "* * * * *"
func (c *Crontab) AddByID(id string, spec string, cmd cro.Job) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if _, ok := c.ids[id]; ok {
		return nil
		//return errors.Errorf("crontab id exists")
	}
	eid, err := c.inner.AddJob(spec, cmd)
	if err != nil {
		return err
	}
	c.ids[id] = eid
	return nil
}

// 添加一个函数作业
func (c *Crontab) AddByFunc(id string, spec string, f func()) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if _, ok := c.ids[id]; ok {
		return nil
	}
	eid, err := c.inner.AddFunc(spec, f)
	if err != nil {
		return err
	}
	c.ids[id] = eid
	return nil
}

// 根据作业ID判断作业是否存在
func (c *Crontab) IsExists(jid string) bool {
	_, exist := c.ids[jid]
	return exist
}
