package goerr

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckErrors(t *testing.T) {
	var err error
	err = NewError("test", NotFound)
	assert.True(t, IsType(err, NotFound))
	assert.False(t, IsType(err, BadRequest))
	assert.False(t, IsType(errors.New("whatever error"), NotFound))
}

func TestCheckContext(t *testing.T) {
	var err error
	err = NewError("test", NotFound)
	assert.NotNil(t, GetCtx(err))
	err = errors.New("arbitrary")
	assert.Nil(t, GetCtx(err))
	err = NewWithCtx("test", BadRequest, map[string]string{"ctx1": "ok"})
	c := GetCtx(err)
	assert.NotNil(t, c)
	assert.Equal(t, "ok", (*c)["ctx1"])
}

func TestGetType(t *testing.T) {
	var err error
	err = NewError("test", NotFound)
	assert.Equal(t, NotFound, GetType(err))
	err = NewError("test", BadRequest)
	assert.Equal(t, BadRequest, GetType(err))
	err = errors.New("test")
	assert.Equal(t, Unrecognized, GetType(err))
}
