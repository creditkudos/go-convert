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
	array := []string{"a", "b", "c", "&"}

	Convey("Populated array", t, func() {
		result := stringArrayContains(array, "a")
		So(result, ShouldBeTrue)
		result = stringArrayContains(array, "f")
		So(result, ShouldBeFalse)
	})

	Convey("Nil array", t, func() {
		result := stringArrayContains(nil, "a")
		So(result, ShouldBeTrue)
	})
}

func TestV2ProtosToCSV(t *testing.T) {

}

func TestProtoHeaders(t *testing.T) {
	Convey("All fields created", t, func() {
		info := csvInfo{
			make(map[string]map[string][]string),
			make(map[string]bool),
		}

		protoHeader(v2TestMaster[0], &info, nil, "", "", "")
		data := info.Data
		fields := info.FieldRequired

		for _, b := range fields {
			So(b, ShouldBeTrue)
		}

		master := data["Master"]
		minion := data["Minion"]
		child := data["Child"]

		So(master["a"], ShouldBeEmpty)
		So(master["i"], ShouldBeNil)
		So(master["j"], ShouldBeNil)
		So(master["j.ma"], ShouldBeEmpty)
		So(master["j.md"], ShouldBeNil)
		So(master["j.md.ca"], ShouldBeEmpty)

		So(minion["Master.id"], ShouldBeEmpty)
		So(minion["md"], ShouldBeNil)
		So(minion["md.ca"], ShouldBeEmpty)

		So(child["Minion.id"], ShouldBeEmpty)
		So(child["Master.id"], ShouldBeNil)
		So(child["ca"], ShouldBeEmpty)
	})

	Convey("Certain fields created", t, func() {

	})
}

func TestProtoBody(t *testing.T) {

}
