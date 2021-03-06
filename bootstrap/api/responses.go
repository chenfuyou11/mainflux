//
// Copyright (c) 2018
// Mainflux
//
// SPDX-License-Identifier: Apache-2.0
//

package api

import (
	"fmt"
	"net/http"

	"github.com/mainflux/mainflux/bootstrap"

	"github.com/mainflux/mainflux"
)

var (
	_ mainflux.Response = (*removeRes)(nil)
	_ mainflux.Response = (*configRes)(nil)
	_ mainflux.Response = (*stateRes)(nil)
	_ mainflux.Response = (*viewRes)(nil)
	_ mainflux.Response = (*listRes)(nil)
)

type removeRes struct{}

func (res removeRes) Code() int {
	return http.StatusNoContent
}

func (res removeRes) Headers() map[string]string {
	return map[string]string{}
}

func (res removeRes) Empty() bool {
	return true
}

type configRes struct {
	id      string
	created bool
}

func (res configRes) Code() int {
	if res.created {
		return http.StatusCreated
	}

	return http.StatusOK
}

func (res configRes) Headers() map[string]string {
	if res.created {
		return map[string]string{
			"Location": fmt.Sprintf("/configs/%s", res.id),
		}
	}

	return map[string]string{}
}

func (res configRes) Empty() bool {
	return true
}

type viewRes struct {
	MFThing     string          `json:"mainflux_id,omitempty"`
	MFKey       string          `json:"mainflux_key,omitempty"`
	Channels    []string        `json:"channels,omitempty"`
	ExternalID  string          `json:"external_id"`
	ExternalKey string          `json:"external_key,omitempty"`
	Content     string          `json:"content,omitempty"`
	State       bootstrap.State `json:"state,omitempty"`
}

func (res viewRes) Code() int {
	return http.StatusOK
}

func (res viewRes) Headers() map[string]string {
	return map[string]string{}
}

func (res viewRes) Empty() bool {
	return false
}

type listRes struct {
	Configs []viewRes `json:"configs"`
}

func (res listRes) Code() int {
	return http.StatusOK
}

func (res listRes) Headers() map[string]string {
	return map[string]string{}
}

func (res listRes) Empty() bool {
	return false
}

type stateRes struct{}

func (res stateRes) Code() int {
	return http.StatusOK
}

func (res stateRes) Headers() map[string]string {
	return map[string]string{}
}

func (res stateRes) Empty() bool {
	return true
}
