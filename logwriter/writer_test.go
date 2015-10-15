package logwriter

import (
//	"time"

//	"github.com/megamsys/gulp/provision"
//	"gopkg.in/check.v1"
)
/*
type WriterSuite struct {
	//	conn *db.Storage
}

var _ = check.Suite(&WriterSuite{})

func (s *WriterSuite) TestLogWriter(c *check.C) {
	a := provision.Box{Name: "testing"}
	writer := LogWriter{Box: &a}
	data := []byte("ble")
	_, err := writer.Write(data)
	c.Assert(err, check.IsNil)
}

func (s *WriterSuite) TestLogWriterShouldReturnTheDataSize(c *check.C) {
	a := provision.Box{}
	writer := LogWriter{Box: &a}
	data := []byte("ble")
	n, err := writer.Write(data)
	c.Assert(err, check.IsNil)
	c.Assert(n, check.Equals, len(data))
}

func (s *WriterSuite) TestLogWriterAsync(c *check.C) {
	a := provision.Box{}
	writer := LogWriter{Box: &a}
	writer.Async()
	data := []byte("ble")
	_, err := writer.Write(data)
	c.Assert(err, check.IsNil)
	writer.Close()
	err = writer.Wait(5 * time.Second)
	c.Assert(err, check.IsNil)
}

func (s *WriterSuite) TestLogWriterAsyncCopySlice(c *check.C) {
	a := provision.Box{}
	writer := LogWriter{Box: &a}
	writer.Async()
	for i := 0; i < 100; i++ {
		data := []byte("ble")
		_, err := writer.Write(data)
		data[0] = 'X'
		c.Assert(err, check.IsNil)
	}
	writer.Close()
	err := writer.Wait(5 * time.Second)
	c.Assert(err, check.IsNil)
	//	instance := provision.Box{}
	//	err = s.conn.Apps().Find(bson.M{"name": a.Name}).One(&instance)
	//	logs, err := instance.LastLogs(100, provision.Boxlog{})
	//	c.Assert(err, check.IsNil)
	//	c.Assert(logs, check.HasLen, 100)
	//	for i := 0; i < 100; i++ {
	//		c.Assert(logs[i].Message, check.Equals, "ble")
	//		c.Assert(logs[i].Source, check.Equals, "tsuru")
	//	}
}*/
