package csv

import (
	"io/ioutil"
	"strconv"
	"strings"
	"testing"

	"github.com/creditkudos/go-convert/test_protos"
	pb "github.com/golang/protobuf/proto"
	. "github.com/smartystreets/goconvey/convey"
)

var (
	v1TestData = []pb.Message{
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
	v2TestData, _ = V1ToV2(v1TestData)
)

func TestStringArrayContains(t *testing.T) {
	array := []string{"a", "b", "c", "&"}

	Convey("Populated array", t, func() {
		result := shouldIncludeField(array, "a")
		So(result, ShouldBeTrue)
		result = shouldIncludeField(array, "f")
		So(result, ShouldBeFalse)
	})

	Convey("Nil array", t, func() {
		result := shouldIncludeField(nil, "a")
		So(result, ShouldBeTrue)
	})
}

func TestRegression(t *testing.T) {
	protos := []pb.Message{}

	for i := 0; i < 3; i++ {
		bytes, err := ioutil.ReadFile("./test_data/testprotodata" + strconv.Itoa(i))
		if err != nil {
			t.Fatal(err)
		}
		p := &test_protos.Master{}
		pb.Unmarshal(bytes, p)
		protos = append(protos, p)
	}

	masterBytes, err := ioutil.ReadFile("./test_data/Master.txt")
	if err != nil {
		t.Fatal(err)
	}
	minionBytes, err := ioutil.ReadFile("./test_data/Minion.txt")
	if err != nil {
		t.Fatal(err)
	}
	childBytes, err := ioutil.ReadFile("./test_data/Child.txt")
	if err != nil {
		t.Fatal(err)
	}

	Convey("Correctly produces csv files", t, func() {
		csvData, err := V1ProtosToCSV(protos, nil)

		So(err, ShouldBeNil)
		So(strings.Split(csvData["Master"], "\n"), ShouldResemble, strings.Split(string(masterBytes), "\n"))
		So(strings.Split(csvData["Minion"], "\n"), ShouldResemble, strings.Split(string(minionBytes), "\n"))
		So(strings.Split(csvData["Child"], "\n"), ShouldResemble, strings.Split(string(childBytes), "\n"))
	})
}

func TestV2ProtosToCSV(t *testing.T) {
	Convey("All fields created", t, func() {
		csvData, err := V2ProtosToCSV(v2TestData, nil)

		expectedMaster := []string{
			"a,b,c,d,e,f,g,h,j.ma,j.mb,j.md.ca,k,l",
			"string,1,-2,3,4,5.5,6.5,true,minion-string,7,child-string,,",
			",0,0,0,0,0,0,false,minion-string 2,7,child-string 2,,",
			"",
		}
		expectedMinion := []string{
			"Master.id,Master.m.key,ma,mb,md.ca",
			"1,,minion-string &2,0,",
			"1,,minion-string =3,0,child-string ",
			"",
		}
		expectedChild := []string{
			"Master.id,Minion.id,ca",
			",0,child-string ",
			"",
		}

		So(err, ShouldBeNil)
		So(csvData["Master"], ShouldEqual, strings.Join(expectedMaster, "\n"))
		So(csvData["Minion"], ShouldEqual, strings.Join(expectedMinion, "\n"))
		So(csvData["Child"], ShouldEqual, strings.Join(expectedChild, "\n"))
	})

	Convey("Specific fields created", t, func() {
		includedFields := make(map[string][]string)
		includedFields["Master"] = []string{"a", "b", "c", "j", "i", "j.ma"}
		includedFields["Minion"] = []string{"ma", "mb", "md"}
		includedFields["Child"] = []string{"ca"}
		csvData, err := V2ProtosToCSV(v2TestData, includedFields)

		So(err, ShouldBeNil)
		So(csvData["Master"], ShouldEqual,
			"a,b,c,j.ma\n"+
				"string,1,-2,minion-string\n"+
				",0,0,minion-string 2\n")
		So(csvData["Minion"], ShouldEqual,
			"ma,mb\n"+
				"minion-string &2,0\n"+
				"minion-string =3,0\n")
		So(csvData["Child"], ShouldEqual, "")
	})
}

func TestPopulateFieldNames(t *testing.T) {
	Convey("All fields created", t, func() {
		info := csvInfo{
			make(map[string]map[string][]string),
			make(map[string]int),
		}

		populateFieldNames(v2TestData[0], &info, nil, "", "", "")
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
		So(master["l"], ShouldBeEmpty)
		So(master["m"], ShouldBeNil)

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

		includedFields := make(map[string][]string)
		includedFields["Master"] = []string{"a", "b", "c", "j", "i", "j.ma"}
		includedFields["Minion"] = []string{"ma", "mb", "md"}
		includedFields["Child"] = []string{"ca"}

		err := populateFieldNames(v2TestData[0], &info, includedFields, "", "", "")
		So(err, ShouldBeNil)

		data := info.Data
		master := data["Master"]
		minion := data["Minion"]
		child := data["Child"]

		So(master["a"], ShouldBeEmpty)
		So(master["b"], ShouldBeEmpty)
		So(master["c"], ShouldBeEmpty)
		So(master["d"], ShouldBeEmpty)
		So(master["l"], ShouldBeNil)
		So(master["m"], ShouldBeNil)

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

func TestPopulateBody(t *testing.T) {
	Convey("All fields created", t, func() {
		info := csvInfo{
			make(map[string]map[string][]string),
			make(map[string]int),
		}

		populateFieldNames(v2TestData[0], &info, nil, "", "", "")
		populateBody(v2TestData[0], &info, "", "", "", 0)
		populateBody(v2TestData[1], &info, "", "", "", 0)

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

		includedFields := make(map[string][]string)
		includedFields["Master"] = []string{"a", "b", "c", "j", "i", "j.md", "j.md.ca"}
		includedFields["Minion"] = []string{"Master.id", "ma", "mb", "md"}
		includedFields["Child"] = []string{"ca"}

		populateFieldNames(v2TestData[0], &info, includedFields, "", "", "")
		populateBody(v2TestData[0], &info, "", "", "", 0)
		populateBody(v2TestData[1], &info, "", "", "", 0)

		data := info.Data
		master := data["Master"]
		minion := data["Minion"]

		So(master["a"][0], ShouldEqual, "string")
		So(master["a"][1], ShouldEqual, "")
		So(master["j"], ShouldBeNil)
		So(master["j.md"], ShouldBeNil)
		So(master["j.md.ca"][0], ShouldEqual, "child-string")
		So(master["j.md.ca"][1], ShouldEqual, "child-string 2")

		So(minion["Master.id"][0], ShouldEqual, "1")
		So(minion["Master.id"][1], ShouldEqual, "1")
		So(minion["ma"][0], ShouldEqual, "minion-string &2")
		So(minion["mb"][1], ShouldEqual, "0")
	})
}
