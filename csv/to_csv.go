package csv

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"strconv"

	pbV1 "github.com/golang/protobuf/proto"
	pb "google.golang.org/protobuf/proto"
)

type csvInfo struct {
	Data          map[string]map[string][]string
	FieldRequired map[string]bool
}

func V1ProtosToCSV(v1Protos []pbV1.Message, fields []string) (map[string]string, error) {
	v2Protos, err := V1ToV2(v1Protos)
	if err != nil {
		return nil, err
	}

	return V2ProtosToCSV(v2Protos, fields)
}

func V2ProtosToCSV(protos []pb.Message, fields []string) (map[string]string, error) {
	if len(protos) == 0 {
		return nil, fmt.Errorf("cannot convert empty array of protos")
	}

	/*
		Generate a map of each CSV file, each containing a map of the fieldname with an array of row values
	*/
	info := csvInfo{
		make(map[string]map[string][]string),
		make(map[string]bool),
	}
	protoHeader(protos[0], &info, fields, "", "", "")
	for _, proto := range protos {
		protoBody(proto, &info, "", "", "", 0)
	}

	// Convert the map into rows
	csvSlice := make(map[string][][]string)
	for file, data := range info.Data {
		csvSlice[file] = make([][]string, len(protos)+1)
		for field, fieldValues := range data {
			csvSlice[file][0] = append(csvSlice[file][0], field)

			for i, value := range fieldValues {
				csvSlice[file][i+1] = append(csvSlice[file][i+1], value)
			}
		}
	}

	//Convert to CSV strings
	csvs := make(map[string]string)
	for file, data := range csvSlice {
		buff := bytes.Buffer{}
		writer := csv.NewWriter(&buff)
		writer.WriteAll(data)

		csvs[file] = buff.String()
	}

	return csvs, nil
}

func protoHeader(proto pb.Message, csv *csvInfo, fields []string, parent, parentFile, file string) {
	pr := proto.ProtoReflect()
	d := pr.Descriptor()
	pFields := d.Fields()

	// Initialise new CSV file if necessary
	pName := string(d.Name())
	if file == "" {
		file = pName
	}
	if len(csv.Data[file]) == 0 {
		csv.Data[file] = make(map[string][]string)
	}

	// Add index of parent file if neccessary
	if parentFile != "" {
		csv.Data[file][parentFile+".id"] = make([]string, 0)
	}

	for i := 0; i < pFields.Len(); i++ {
		field := pFields.Get(i)
		fieldName := string(field.Name())

		if parent != "" {
			fieldName = parent + "." + fieldName
		}

		if !stringArrayContains(fields, fieldName) {
			csv.FieldRequired[fieldName] = false
			continue
		}

		csv.FieldRequired[fieldName] = true

		if field.Cardinality().GoString() == "Repeated" {
			repeated := pr.Get(field).List()
			for j := 0; j < repeated.Len(); j++ {
				m := repeated.Get(j).Message().Interface()

				protoHeader(m, csv, fields, "", file, "")
			}
			continue
		}
		if field.Kind().GoString() == "MessageKind" {
			protoHeader(pr.Get(field).Message().Interface(), csv, fields, fieldName, "", file)
		} else if field.Kind().GoString() == "GroupKind" {
		} else {
			csv.Data[file][fieldName] = make([]string, 0)
		}
	}
}

func protoBody(proto pb.Message, csv *csvInfo, parent, parentType, file string, parentID int) {
	pr := proto.ProtoReflect()
	d := pr.Descriptor()
	pFields := d.Fields()
	pName := string(d.Name())
	if file == "" {
		file = pName
	}

	// Populate index of parent if neccessary
	if parentType != "" && csv.FieldRequired[parentType+".id"] {
		csv.Data[file][parentType+".id"] = append(csv.Data[file][parentType], strconv.Itoa(parentID))
	}

	// Populate remaining fields
	for i := 0; i < pFields.Len(); i++ {
		field := pFields.Get(i)
		fieldName := string(field.Name())

		if parent != "" {
			fieldName = parent + "." + fieldName
		}

		if field.Cardinality().GoString() == "Repeated" {
			repeated := pr.Get(field).List()
			for j := 0; j < repeated.Len(); j++ {
				m := repeated.Get(j).Message().Interface()

				protoBody(m, csv, "", file, "", len(csv.Data[file][fieldName]))
			}
			continue
		}
		// // Only populate desired fields
		if !csv.FieldRequired[fieldName] {
			fmt.Println(fieldName)
			continue
		}
		if field.Kind().GoString() == "MessageKind" {
			protoBody(pr.Get(field).Message().Interface(), csv, fieldName, "", file, 0)
		} else if field.Kind().GoString() == "GroupKind" {
		} else {
			value := pr.Get(field).String()

			csv.Data[file][fieldName] = append(csv.Data[file][fieldName], value)
		}
	}

}

/* Convert between V1 and V2 go implementations of protos, while both protos are V3 the go representation
of them is different between versions.
*/
func V1ToV2(v1Protos []pbV1.Message) ([]pb.Message, error) {
	v2Protos := make([]pb.Message, len(v1Protos))

	for i, v1Proto := range v1Protos {
		v2Protos[i] = pbV1.MessageV2(v1Proto)
	}

	return v2Protos, nil
}

func stringArrayContains(array []string, value string) bool {
	//if nil then always true, if empty then always false
	if array == nil {
		return true
	}

	for _, val := range array {
		if val == value {
			return true
		}
	}

	return false
}
