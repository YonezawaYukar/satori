package satori

import (
	"satori/auth"
	"time"
)

type satori struct {
	//request's basic info
	url    string
	method int
	data   interface{}
	//request's header
	header      map[string]string
	cookie      map[string]string
	user_agent  string
	basic_auth  auth.Basic_auth
	bearer_auth auth.Bearer_auth
	//request's other configure
	proxy    string
	time_out time.Duration
	ssl      struct {
		verify bool
	}
}

func New() *satori {
	return new(satori).TimeOut(10 * time.Second)
}

func (r *satori) TimeOut(time time.Duration) *satori {
	r.time_out = time
	return r
}
