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
				Ma: "minion-string 2",
				Mb: 7,
				Mc: nil,
				Md: &test_protos.Child{
					Ca: "child-string 2",
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
			make(map[string]int),
		}

		populateFieldNames(v2TestMaster[0], &info, nil, "", "", "")
		data := info.Data

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
		So(child["Master.id"], ShouldBeEmpty)
		So(child["ca"], ShouldBeEmpty)
	})

	Convey("Specified fields created", t, func() {
		info := csvInfo{
			make(map[string]map[string][]string),
			make(map[string]int),
		}

		headerFields := make(map[string][]string)
		headerFields["Master"] = []string{"a", "b", "c", "j", "i", "j.ma"}
		headerFields["Minion"] = []string{"ma", "mb", "md"}
		headerFields["Child"] = []string{"ca"}

		populateFieldNames(v2TestMaster[0], &info, headerFields, "", "", "")
		data := info.Data
		master := data["Master"]
		minion := data["Minion"]
		child := data["Child"]

		So(master["a"], ShouldBeEmpty)
		So(master["b"], ShouldBeEmpty)
		So(master["c"], ShouldBeEmpty)
		So(master["d"], ShouldBeEmpty)

		// Check subfield
		So(master["j"], ShouldBeNil)
		So(master["j.ma"], ShouldBeEmpty)

		// Check repeated fields
		So(minion["ma"], ShouldBeEmpty)
		So(minion["mb"], ShouldBeEmpty)
		So(minion["mc"], ShouldBeNil)
		So(child["ca"], ShouldBeEmpty)
	})
}

func TestProtoBody(t *testing.T) {
	Convey("All fields created", t, func() {
		info := csvInfo{
			make(map[string]map[string][]string),
			make(map[string]int),
		}

		populateFieldNames(v2TestMaster[0], &info, nil, "", "", "")
		populateBody(v2TestMaster[0], &info, "", "", "", 0)
		populateBody(v2TestMaster[1], &info, "", "", "", 0)

		data := info.Data
		master := data["Master"]
		minion := data["Minion"]
		child := data["Child"]

		So(master["a"][0], ShouldEqual, "string")
		So(master["a"][1], ShouldEqual, "")
		So(master["j"], ShouldBeNil)
		So(master["j.ma"][0], ShouldEqual, "minion-string")
		So(master["j.ma"][1], ShouldEqual, "minion-string 2")
		So(master["j.md"], ShouldBeNil)
		So(master["j.md.ca"][0], ShouldEqual, "child-string")
		So(master["j.md.ca"][1], ShouldEqual, "child-string 2")

		So(minion["Master.id"][0], ShouldEqual, "1")
		So(minion["Master.id"][1], ShouldEqual, "1")
		So(minion["ma"][0], ShouldEqual, "minion-string &2")
		So(minion["ma"][1], ShouldEqual, "minion-string =3")

		So(child["Master.id"][0], ShouldEqual, "")
	})

	Convey("Specified fields created", t, func() {
		info := csvInfo{
			make(map[string]map[string][]string),
			make(map[string]int),
		}

		headerFields := make(map[string][]string)
		headerFields["Master"] = []string{"a", "b", "c", "j", "i", "j.ma", "j.md", "j.md.ca"}
		headerFields["Minion"] = []string{"ma", "mb", "md"}
		headerFields["Child"] = []string{"ca"}

		populateFieldNames(v2TestMaster[0], &info, headerFields, "", "", "")
		populateBody(v2TestMaster[0], &info, "", "", "", 0)
		populateBody(v2TestMaster[1], &info, "", "", "", 0)

		data := info.Data
		master := data["Master"]
		// minion := data["Minion"]
		// child := data["Child"]

		So(master["a"][0], ShouldEqual, "string")
		So(master["a"][1], ShouldEqual, "")
		So(master["j"], ShouldBeNil)
		So(master["j.ma"][0], ShouldEqual, "minion-string")
		So(master["j.ma"][1], ShouldEqual, "minion-string 2")
		So(master["j.md"], ShouldBeNil)
		So(master["j.md.ca"][0], ShouldEqual, "child-string")
		So(master["j.md.ca"][1], ShouldEqual, "child-string 2")
	})
}
