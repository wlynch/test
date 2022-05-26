package main

import (
	"io/ioutil"
	"log"

	"github.com/grafeas/grafeas/proto/v1/grafeas_go_proto"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	b, _ := ioutil.ReadFile("provenance.json")
	m := new(grafeas_go_proto.SlsaProvenanceZeroTwo)
	if err := protojson.Unmarshal(b, m); err != nil {
		log.Fatal(err)
	}
}
