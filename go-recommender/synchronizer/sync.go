package synchronizer

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/codfrm/cago"
	"github.com/codfrm/cago/config"
	"github.com/codfrm/cago/database/mysql"
	"github.com/codfrm/cago/pkg/logger"
	"github.com/go-mysql-org/go-mysql/canal"
	mysql2 "github.com/go-mysql-org/go-mysql/mysql"
	"github.com/go-mysql-org/go-mysql/replication"
	"github.com/scriptscat/dz_recommend/go-recommender/es"
	"go.uber.org/zap"
)

type syncToEs struct {
	ctx    context.Context
	cancel context.CancelFunc
	config *config.Config
	canal  *canal.Canal
}

// 同步到es
func (s *syncToEs) start() error {
	s.canal.SetEventHandler(s)
	go func() {
		if err := s.canal.Run(); err != nil {
			logger.Ctx(cago.Background()).Error("synchronizer error", zap.Error(err))
			s.cancel()
		}
	}()
	go func() {
		<-s.ctx.Done()
		s.canal.Close()
	}()
	var num int64
	thread := &ForumThread{}
	if err := mysql.Ctx(nil).Model(thread).Count(&num).Error; err != nil {
		return err
	}
	// 简单点 es与mysql的数量对不上就直接全量同步
	count := es.Ctx(nil).Count
	resp, err := es.Ctx(nil).Count(count.WithIndex(thread.CollectName()))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNotFound {
		return errors.New("es count error")
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	m := make(map[string]interface{})
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	esNum, _ := m["count"].(float64)
	if resp.StatusCode == http.StatusNotFound || esNum != float64(num) {
		go s.syncAll()
	}
	return nil
}

func (s *syncToEs) syncAll() {
	logger.Ctx(nil).Info("全量同步")
	page := 0
	forumThread := &ForumThread{}
	if _, err := es.Ctx(nil).Indices.Delete([]string{forumThread.CollectName()}); err != nil {
		logger.Ctx(nil).Error("delete index error", zap.Error(err))
	}
	thread := &ForumThread{}
	for {
		list := make([]*ForumThread, 0, 100)
		if err := mysql.Ctx(nil).Model(thread).Offset(page * 100).Limit(100).Scan(&list).Error; err != nil {
			logger.Ctx(nil).Error("synchronizer error", zap.Error(err))
			s.cancel()
			return
		}
		if len(list) == 0 {
			break
		}
		for _, v := range list {
			b, _ := json.Marshal(v)
			func() {
				resp, err := es.Ctx(nil).Create(
					thread.CollectName(), strconv.FormatInt(int64(v.Tid), 10),
					bytes.NewReader(b),
				)
				if err != nil {
					logger.Ctx(nil).Error("insert forum_thread error", zap.Uint("tid", v.Tid),
						zap.Error(err))
					return
				}
				defer resp.Body.Close()
				if resp.StatusCode != http.StatusCreated {
					b, _ := io.ReadAll(resp.Body)
					logger.Ctx(nil).Error("insert forum_thread error", zap.Uint("tid", v.Tid),
						zap.ByteString("body", b),
						zap.Int("status", resp.StatusCode))
					return
				}
				logger.Ctx(nil).Info("insert forum_thread success", zap.Uint("tid", v.Tid))
			}()
		}
		page++
	}
}

func (s *syncToEs) OnRotate(roateEvent *replication.RotateEvent) error {
	return nil
}

func (s *syncToEs) OnTableChanged(schema string, table string) error {
	return nil
}

func (s *syncToEs) OnDDL(nextPos mysql2.Position, queryEvent *replication.QueryEvent) error {
	return nil
}

func (s *syncToEs) OnRow(e *canal.RowsEvent) error {
	return nil
}

func (s *syncToEs) OnXID(nextPos mysql2.Position) error {
	return nil
}

func (s *syncToEs) OnGTID(gtid mysql2.GTIDSet) error {
	return nil
}

func (s *syncToEs) OnPosSynced(pos mysql2.Position, set mysql2.GTIDSet, force bool) error {
	return nil
}

func (s *syncToEs) String() string {
	return "synchronizer"
}
