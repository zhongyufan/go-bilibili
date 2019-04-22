// Code generated by $GOPATH/src/go-common/app/tool/cache/mc. DO NOT EDIT.

/*
  Package archive is a generated mc cache package.
  It is generated from:
  type _mc interface {
		// mc: -key=staffKey
		CacheStaffData(c context.Context, key int64) ([]*arcMdl.Staff, error)
		// 这里也支持自定义注释 会替换默认的注释
		// mc: -key=staffKey -expire=3 -encode=json|gzip
		AddCacheStaffData(c context.Context, key int64, value []*arcMdl.Staff) error
		// mc: -key=staffKey
		DelCacheStaffData(c context.Context, key int64) error
		//mc: -key=viewPointCacheKey -expire=_viewPointExp -encode=json
		AddCacheViewPoint(c context.Context, aid int64, vp *arcMdl.ViewPointRow, cid int64) (err error)
		//mc: -key=viewPointCacheKey
		CacheViewPoint(c context.Context, aid int64, cid int64) (vp *arcMdl.ViewPointRow, err error)
	}
*/

package archive

import (
	"context"
	"fmt"

	arcMdl "go-common/app/interface/main/creative/model/archive"
	"go-common/library/cache/memcache"
	"go-common/library/log"
	"go-common/library/stat/prom"
)

var _ _mc

// CacheStaffData get data from mc
func (d *Dao) CacheStaffData(c context.Context, id int64) (res []*arcMdl.Staff, err error) {
	conn := d.mc.Get(c)
	defer conn.Close()
	key := staffKey(id)
	reply, err := conn.Get(key)
	if err != nil {
		if err == memcache.ErrNotFound {
			err = nil
			return
		}
		prom.BusinessErrCount.Incr("mc:CacheStaffData")
		log.Errorv(c, log.KV("CacheStaffData", fmt.Sprintf("%+v", err)), log.KV("key", key))
		return
	}
	res = []*arcMdl.Staff{}
	err = conn.Scan(reply, &res)
	if err != nil {
		prom.BusinessErrCount.Incr("mc:CacheStaffData")
		log.Errorv(c, log.KV("CacheStaffData", fmt.Sprintf("%+v", err)), log.KV("key", key))
		return
	}
	return
}

// AddCacheStaffData 这里也支持自定义注释 会替换默认的注释
func (d *Dao) AddCacheStaffData(c context.Context, id int64, val []*arcMdl.Staff) (err error) {
	if len(val) == 0 {
		return
	}
	conn := d.mc.Get(c)
	defer conn.Close()
	key := staffKey(id)
	item := &memcache.Item{Key: key, Object: val, Expiration: 3, Flags: memcache.FlagJSON | memcache.FlagGzip}
	if err = conn.Set(item); err != nil {
		prom.BusinessErrCount.Incr("mc:AddCacheStaffData")
		log.Errorv(c, log.KV("AddCacheStaffData", fmt.Sprintf("%+v", err)), log.KV("key", key))
		return
	}
	return
}

// DelCacheStaffData delete data from mc
func (d *Dao) DelCacheStaffData(c context.Context, id int64) (err error) {
	conn := d.mc.Get(c)
	defer conn.Close()
	key := staffKey(id)
	if err = conn.Delete(key); err != nil {
		if err == memcache.ErrNotFound {
			err = nil
			return
		}
		prom.BusinessErrCount.Incr("mc:DelCacheStaffData")
		log.Errorv(c, log.KV("DelCacheStaffData", fmt.Sprintf("%+v", err)), log.KV("key", key))
		return
	}
	return
}

// AddCacheViewPoint Set data to mc
func (d *Dao) AddCacheViewPoint(c context.Context, id int64, val *arcMdl.ViewPointRow, cid int64) (err error) {
	if val == nil {
		return
	}
	conn := d.mc.Get(c)
	defer conn.Close()
	key := viewPointCacheKey(id, cid)
	item := &memcache.Item{Key: key, Object: val, Expiration: _viewPointExp, Flags: memcache.FlagJSON}
	if err = conn.Set(item); err != nil {
		prom.BusinessErrCount.Incr("mc:AddCacheViewPoint")
		log.Errorv(c, log.KV("AddCacheViewPoint", fmt.Sprintf("%+v", err)), log.KV("key", key))
		return
	}
	return
}

// CacheViewPoint get data from mc
func (d *Dao) CacheViewPoint(c context.Context, id int64, cid int64) (res *arcMdl.ViewPointRow, err error) {
	conn := d.mc.Get(c)
	defer conn.Close()
	key := viewPointCacheKey(id, cid)
	reply, err := conn.Get(key)
	if err != nil {
		if err == memcache.ErrNotFound {
			err = nil
			return
		}
		prom.BusinessErrCount.Incr("mc:CacheViewPoint")
		log.Errorv(c, log.KV("CacheViewPoint", fmt.Sprintf("%+v", err)), log.KV("key", key))
		return
	}
	res = &arcMdl.ViewPointRow{}
	err = conn.Scan(reply, res)
	if err != nil {
		prom.BusinessErrCount.Incr("mc:CacheViewPoint")
		log.Errorv(c, log.KV("CacheViewPoint", fmt.Sprintf("%+v", err)), log.KV("key", key))
		return
	}
	return
}