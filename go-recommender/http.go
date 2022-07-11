package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/allegro/bigcache/v3"
	"github.com/codfrm/cago/database/mysql"
	"github.com/codfrm/cago/mux"
	"github.com/codfrm/cago/pkg/httputils"
	"github.com/codfrm/cago/pkg/utils"
	"github.com/scriptscat/dz_recommend/go-recommender/es"
	"github.com/scriptscat/dz_recommend/go-recommender/synchronizer"
)

type recommendResponse struct {
	Tid     int64  `json:"tid"`
	Subject string `json:"subject"`
}

func recommend() func(ctx *mux.WebContext) {
	cache, _ := bigcache.NewBigCache(bigcache.DefaultConfig(time.Hour * 24))
	return func(ctx *mux.WebContext) {
		httputils.Handle(ctx, func() interface{} {
			tid := utils.ToNumber[int64](ctx.Query("tid"))
			entry, err := cache.Get("recommend_" + fmt.Sprintf("%d", tid))
			b := make(map[string]interface{})
			if err != nil {
				if err != bigcache.ErrEntryNotFound {
					return err
				}
				thread := &synchronizer.ForumThread{}
				err = mysql.Ctx(ctx).Where("tid = ?", tid).Find(&thread).Error
				if err != nil {
					return err
				}
				query := map[string]interface{}{
					"query": map[string]interface{}{
						"match": map[string]interface{}{
							"subject": thread.Subject,
						},
					},
					"size": 7,
				}
				var buf bytes.Buffer
				if err := json.NewEncoder(&buf).Encode(query); err != nil {
					return err
				}
				search := es.Ctx(ctx).Search
				resp, err := es.Ctx(ctx).Search(
					search.WithIndex(thread.CollectName()),
					search.WithBody(&buf),
					search.WithTrackTotalHits(true),
					search.WithPretty())
				if err != nil {
					return err
				}
				respByte, _ := io.ReadAll(resp.Body)
				if err := json.Unmarshal(respByte, &b); err != nil {
					return err
				}
				if resp.IsError() {
					return fmt.Errorf("elasticsearch error: [%s] %s: %s",
						resp.Status(),
						b["error"].(map[string]interface{})["type"],
						b["error"].(map[string]interface{})["reason"],
					)
				}
				cache.Set("recommend_"+fmt.Sprintf("%d", tid), respByte)
			} else {
				if err := json.Unmarshal(entry, &b); err != nil {
					return err
				}
			}
			ret := make([]*recommendResponse, 0)
			n := 0
			for _, v := range b["hits"].(map[string]interface{})["hits"].([]interface{}) {
				m := v.(map[string]interface{})
				source := m["_source"].(map[string]interface{})
				if n == 6 {
					break
				}
				if int64(source["tid"].(float64)) == tid {
					continue
				}
				n++
				q := &recommendResponse{
					Tid:     int64(source["tid"].(float64)),
					Subject: source["subject"].(string),
				}
				ret = append(ret, q)
			}
			return ret
		})
	}
}
