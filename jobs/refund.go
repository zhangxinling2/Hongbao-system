package jobs

import (
	"fmt"
	"github.com/go-redsync/redsync"
	"github.com/gomodule/redigo/redis"
	log "github.com/sirupsen/logrus"
	"github.com/tietang/go-utils"
	"github.com/zhangxinling2/infra"
	"github.com/zhangxinling2/infra/base"
	"resk/core/envelopes"
	services "resk/services"
	"time"
)

type RefundExpiredJobStarter struct {
	infra.BaseStarter
	ExpiredGoods []services.RedEnvelopeGoodsDTO
	ticker       *time.Ticker
	mutex        *redsync.Mutex
}

func (r *RefundExpiredJobStarter) Init(ctx infra.StarterContext) {
	dur := base.Props().GetDurationDefault("jobs.refund.internal", time.Minute)
	r.ticker = time.NewTicker(dur)
	maxIdel := base.Props().GetIntDefault("redis.maxIdel", 2)
	maxActive := base.Props().GetIntDefault("redis.maxActive", 4)
	timeout := base.Props().GetDurationDefault("redis.timeout", time.Minute)
	addr := base.Props().GetDefault("redis.addr", "127.0.0.1:6379")
	pools := make([]redsync.Pool, 1)
	pool := &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", addr)
		},
		MaxIdle:     maxIdel,
		MaxActive:   maxActive,
		IdleTimeout: timeout,
	}
	pools = append(pools, pool)
	rsync := redsync.New(pools)
	ip, err := utils.GetExternalIP()
	if err != nil {
		ip = "127.0.0.1"
	}
	r.mutex = rsync.NewMutex("lock:RefundExpired",
		redsync.SetExpiry(50*time.Second),
		redsync.SetRetryDelay(3),
		redsync.SetGenValueFunc(func() (string, error) {
			now := time.Now()
			log.Infof("节点%s正在执行过期红包的退款任务", ip)
			return fmt.Sprintf("%d:%s", now.Unix(), ip), nil
		}))
}
func (r *RefundExpiredJobStarter) Start(ctx infra.StarterContext) {
	go func() {
		for {
			c := <-r.ticker.C
			err := r.mutex.Lock()
			if err == nil {
				log.Debug("红包退款开始", c)
				domain := envelopes.ExpiredEnvelopeDomain{}
				domain.Expire()
			} else {
				log.Info("已经有节点开始退款")
			}
			r.mutex.Unlock()
		}
	}()
}
func (r *RefundExpiredJobStarter) Stop(ctx infra.StarterContext) {
	r.ticker.Stop()
}
