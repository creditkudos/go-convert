package csv

import (
	"testing"

	"github.com/creditkudos/go-convert/test_protos"
	pb "github.com/golang/protobuf/proto"
	. "github.com/smartystreets/goconvey/convey"
)

var (
	v1TestMaster = []pb.Message{
		&test_protos.Master{
			A: "string",
			B: 1,
			C: -2,
			D: 3,
			E: 4,
			F: 5.5,
			G: 6.5,
			H: true,
			I: nil,
			J: &test_protos.Minion{
				Ma: "minion-string",
				Mb: 7,
				Mc: nil,
				Md: &test_protos.Child{
					Ca: "child-string",
				},
			},
		},
		&test_protos.Master{
			I: []*test_protos.Minion{
				&test_protos.Minion{
					Ma: "minion-string &2",
					Mc: []*test_protos.Child{
						&test_protos.Child{
							Ca: "child-string ",
						}},
				},
				&test_protos.Minion{
					Ma: "minion-string =3",
					Md: &test_protos.Child{
						Ca: "child-string ",
					}},
			},
			J: &test_protos.Minion{
				Ma: "minion-string",
				Mb: 7,
				Mc: nil,
				Md: &test_protos.Child{
					Ca: "child-string",
				},
			},
		},
	}
	v2TestMaster, _ = V1ToV2(v1TestMaster)
)

func TestStringArrayContains(t *testing.T) {

}

func TestV1ToV2(t *testing.T) {

}

func TestV2ProtosToCSV(t *testing.T) {

}

func TestProtoHeaders(t *testing.T) {
	Convey("All fields created", func() {

	})

	Convey("Certain fields created", func() {

	})
}

func TestProtoBody(t *testing.T) {

}
