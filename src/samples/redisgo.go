package samples

import (
	"bytes"
	"flag"
	"fmt"
	"sync"
	"time"

	"github.com/garyburd/redigo/redis"
)

var (
	rdsMu sync.Mutex
	Rds   = make(map[string]*RedisDriver)
)

type RedisDriver struct {
	//Pool *redis.Pool
	Pool RedisPool
}

type RedisPool interface {
	Get() redis.Conn
}

var useTraceConn = flag.Bool("use_traceconn", false, "Use trace redis connection")

type traceConn struct {
	redis.Conn
}

func fmtArgs(args []interface{}) string {
	var b bytes.Buffer
	for _, a := range args {
		if b.Len() > 0 {
			b.WriteByte(' ')
		}
		fmt.Fprintf(&b, "%v", a)
	}
	return b.String()
}

func (c traceConn) Do(commandName string, args ...interface{}) (reply interface{}, err error) {
	return c.Conn.Do(commandName, args...)
}

func (c traceConn) Send(commandName string, args ...interface{}) error {
	return c.Conn.Send(commandName, args...)
}

func (r *RedisDriver) Connect(host, pwd string, db, maxidle, maxactive int) {
	//spew.Dump("register redis")
	r.Pool = &redis.Pool{
		MaxIdle:     maxidle,
		MaxActive:   maxactive,
		IdleTimeout: 240 * time.Second,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", host, redis.DialConnectTimeout(time.Second), redis.DialPassword(pwd), redis.DialDatabase(db))
			if err != nil {
				return nil, err
			}
			if *useTraceConn {
				return traceConn{c}, nil
			}
			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
}

func RedisRegister(name string, rDriver *RedisDriver) {
	rdsMu.Lock()
	defer rdsMu.Unlock()
	if _, dup := Rds[name]; dup {
		panic("Register called twice for driver " + name)
	}
	Rds[name] = rDriver
}
