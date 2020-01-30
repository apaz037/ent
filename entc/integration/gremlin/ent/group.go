// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/facebookincubator/ent/dialect/gremlin"
	"github.com/facebookincubator/ent/entc/integration/gremlin/ent/groupinfo"
)

// Group is the model entity for the Group schema.
type Group struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Active holds the value of the "active" field.
	Active bool `json:"active,omitempty"`
	// Expire holds the value of the "expire" field.
	Expire time.Time `json:"expire,omitempty"`
	// Type holds the value of the "type" field.
	Type *string `json:"type,omitempty"`
	// MaxUsers holds the value of the "max_users" field.
	MaxUsers int `json:"max_users,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the GroupQuery when eager-loading is set.
	Edges GroupEdges `json:"edges"`
}

// GroupEdges holds the relations/edges for other nodes in the graph.
type GroupEdges struct {
	// Files holds the value of the files edge.
	Files []*File
	// Blocked holds the value of the blocked edge.
	Blocked []*User
	// Users holds the value of the users edge.
	Users []*User
	// Info holds the value of the info edge.
	Info *GroupInfo
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// FilesErr returns the Files value or an error if the edge
// was not loaded in eager-loading.
func (e GroupEdges) FilesErr() ([]*File, error) {
	if e.loadedTypes[0] {
		return e.Files, nil
	}
	return nil, &NotLoadedError{edge: "files"}
}

// BlockedErr returns the Blocked value or an error if the edge
// was not loaded in eager-loading.
func (e GroupEdges) BlockedErr() ([]*User, error) {
	if e.loadedTypes[1] {
		return e.Blocked, nil
	}
	return nil, &NotLoadedError{edge: "blocked"}
}

// UsersErr returns the Users value or an error if the edge
// was not loaded in eager-loading.
func (e GroupEdges) UsersErr() ([]*User, error) {
	if e.loadedTypes[2] {
		return e.Users, nil
	}
	return nil, &NotLoadedError{edge: "users"}
}

// InfoErr returns the Info value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e GroupEdges) InfoErr() (*GroupInfo, error) {
	if e.loadedTypes[3] {
		if e.Info == nil {
			// The edge info was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: groupinfo.Label}
		}
		return e.Info, nil
	}
	return nil, &NotLoadedError{edge: "info"}
}

// FromResponse scans the gremlin response data into Group.
func (gr *Group) FromResponse(res *gremlin.Response) error {
	vmap, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	var scangr struct {
		ID       string  `json:"id,omitempty"`
		Active   bool    `json:"active,omitempty"`
		Expire   int64   `json:"expire,omitempty"`
		Type     *string `json:"type,omitempty"`
		MaxUsers int     `json:"max_users,omitempty"`
		Name     string  `json:"name,omitempty"`
	}
	if err := vmap.Decode(&scangr); err != nil {
		return err
	}
	gr.ID = scangr.ID
	gr.Active = scangr.Active
	gr.Expire = time.Unix(0, scangr.Expire)
	gr.Type = scangr.Type
	gr.MaxUsers = scangr.MaxUsers
	gr.Name = scangr.Name
	return nil
}

// QueryFiles queries the files edge of the Group.
func (gr *Group) QueryFiles() *FileQuery {
	return (&GroupClient{gr.config}).QueryFiles(gr)
}

// QueryBlocked queries the blocked edge of the Group.
func (gr *Group) QueryBlocked() *UserQuery {
	return (&GroupClient{gr.config}).QueryBlocked(gr)
}

// QueryUsers queries the users edge of the Group.
func (gr *Group) QueryUsers() *UserQuery {
	return (&GroupClient{gr.config}).QueryUsers(gr)
}

// QueryInfo queries the info edge of the Group.
func (gr *Group) QueryInfo() *GroupInfoQuery {
	return (&GroupClient{gr.config}).QueryInfo(gr)
}

// Update returns a builder for updating this Group.
// Note that, you need to call Group.Unwrap() before calling this method, if this Group
// was returned from a transaction, and the transaction was committed or rolled back.
func (gr *Group) Update() *GroupUpdateOne {
	return (&GroupClient{gr.config}).UpdateOne(gr)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (gr *Group) Unwrap() *Group {
	tx, ok := gr.config.driver.(*txDriver)
	if !ok {
		panic("ent: Group is not a transactional entity")
	}
	gr.config.driver = tx.drv
	return gr
}

// String implements the fmt.Stringer.
func (gr *Group) String() string {
	var builder strings.Builder
	builder.WriteString("Group(")
	builder.WriteString(fmt.Sprintf("id=%v", gr.ID))
	builder.WriteString(", active=")
	builder.WriteString(fmt.Sprintf("%v", gr.Active))
	builder.WriteString(", expire=")
	builder.WriteString(gr.Expire.Format(time.ANSIC))
	if v := gr.Type; v != nil {
		builder.WriteString(", type=")
		builder.WriteString(*v)
	}
	builder.WriteString(", max_users=")
	builder.WriteString(fmt.Sprintf("%v", gr.MaxUsers))
	builder.WriteString(", name=")
	builder.WriteString(gr.Name)
	builder.WriteByte(')')
	return builder.String()
}

// id returns the int representation of the ID field.
func (gr *Group) id() int {
	id, _ := strconv.Atoi(gr.ID)
	return id
}

// Groups is a parsable slice of Group.
type Groups []*Group

// FromResponse scans the gremlin response data into Groups.
func (gr *Groups) FromResponse(res *gremlin.Response) error {
	vmap, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	var scangr []struct {
		ID       string  `json:"id,omitempty"`
		Active   bool    `json:"active,omitempty"`
		Expire   int64   `json:"expire,omitempty"`
		Type     *string `json:"type,omitempty"`
		MaxUsers int     `json:"max_users,omitempty"`
		Name     string  `json:"name,omitempty"`
	}
	if err := vmap.Decode(&scangr); err != nil {
		return err
	}
	for _, v := range scangr {
		*gr = append(*gr, &Group{
			ID:       v.ID,
			Active:   v.Active,
			Expire:   time.Unix(0, v.Expire),
			Type:     v.Type,
			MaxUsers: v.MaxUsers,
			Name:     v.Name,
		})
	}
	return nil
}

func (gr Groups) config(cfg config) {
	for _i := range gr {
		gr[_i].config = cfg
	}
}
