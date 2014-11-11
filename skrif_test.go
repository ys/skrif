package skrif

import (
	"bytes"
	"log"
	"testing"
)

func TestWriteToStdout(t *testing.T) {
	lol := New()
	var buf bytes.Buffer
	lol.l = log.New(&buf, "", 0)
	lol.SetGlobalContext("app=lol")
	lol.Println("yolo=true") //=> "app=lol yolo=true"
	if buf.String() != "app=lol yolo=true\n" {
		t.Error("expected the lol to be correct" + buf.String())
	}
}

func TestWriteToStdoutWithoutGlobal(t *testing.T) {
	lol := New()
	var buf bytes.Buffer
	lol.l = log.New(&buf, "", 0)
	lol.Println("yolo=true") //=> "yolo=true"
	if buf.String() != "yolo=true\n" {
		t.Error("expected the lol to be correct" + buf.String())
	}
}

func TestWriteToStdoutWithContext(t *testing.T) {
	lol := New()
	var buf bytes.Buffer
	lol.l = log.New(&buf, "", 0)
	lol.SetGlobalContext("app=lol")
	lol.Context("somuch=true", func(lol *Skrif) {
		lol.Println("yolo=true") //=> "app=lol somuch=true yolo=true"
		if buf.String() != "app=lol somuch=true yolo=true\n" {
			t.Error("expected the lol to be correct" + buf.String())
		}
	})
}

func TestWriteToStdoutWithDoubleContext(t *testing.T) {
	lol := New()
	var buf bytes.Buffer
	lol.l = log.New(&buf, "", 0)
	lol.SetGlobalContext("app=lol")
	lol.Context("somuch=true", func(lol *Skrif) {
		lol.Context("RTFM=verbose", func(lol *Skrif) {
			lol.Println("yolo=true") //=> "app=lol somuch=true yolo=true"
			if buf.String() != "app=lol somuch=true RTFM=verbose yolo=true\n" {
				t.Error("expected the lol to be correct" + buf.String())
			}
		})
	})
}
