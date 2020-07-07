package csv

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"sort"
	"strconv"

	pbV1 "github.com/golang/protobuf/proto"
	pb "google.golang.org/protobuf/proto"
)

type csvInfo struct {
	Data map[string]map[string][]string
	Rows map[string]int
}

func V1ProtosToCSV(v1Protos []pbV1.Message, fields map[string][]string) (map[string]string, error) {
	v2Protos, err := V1ToV2(v1Protos)
	if err != nil {
		return nil, err
	}

	return V2ProtosToCSV(v2Protos, fields)
}

func V2ProtosToCSV(protos []pb.Message, fields map[string][]string) (map[string]string, error) {
	if len(protos) == 0 {
		return nil, fmt.Errorf("cannot convert empty array of protos")
	}

	info := csvInfo{
		make(map[string]map[string][]string),
		make(map[string]int),
	}
	populateFieldNames(protos[0], &info, fields, "", "", "")
	for _, proto := range protos {
		populateBody(proto, &info, "", "", "", 0)
	}

	// Convert the map into rows
	csvStrings := make(map[string][][]string)
	for fileName, fileData := range info.Data {
		// Length is number of rows + 1 for field names
		csvStrings[fileName] = make([][]string, info.Rows[fileName]+1)

		// Alpha sort fields, go has intentionally non-deterministic key orderings using #range
		fields := []string{}
		for field := range fileData {
			fields = append(fields, field)
		}
		sort.Strings(fields)

		// Populate sorted fields
		for _, field := range fields {
			// Populate field name
			csvStrings[fileName][0] = append(csvStrings[fileName][0], field)

			// Populate field values
			fieldValues := fileData[field]
			for i, value := range fieldValues {
				csvStrings[fileName][i+1] = append(csvStrings[fileName][i+1], value)
			}
		}
	}

	//Convert to CSV strings
	csvs := make(map[string]string)
	for file, data := range csvStrings {
		buff := bytes.Buffer{}
		writer := csv.NewWriter(&buff)
		writer.WriteAll(data)

		csvs[file] = buff.String()
	}

	return csvs, nil
}

func populateFieldNames(proto pb.Message, csv *csvInfo, includedFields map[string][]string, parent, parentFile, file string) {
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
		parentField := parentFile + ".id"
		if stringArrayContains(includedFields[file], parentField) {
			csv.Data[file][parentField] = make([]string, 0)
		}
	}

	for i := 0; i < pFields.Len(); i++ {
		field := pFields.Get(i)
		fieldName := string(field.Name())

		// Append parent to name eg account.amount.currency
		if parent != "" {
			fieldName = parent + "." + fieldName
		}

		// Only include specified fields
		if !stringArrayContains(includedFields[file], fieldName) {
			continue
		}

		// This will only convert message fields else it will panic.
		if field.Cardinality().GoString() == "Repeated" {
			repeated := pr.Get(field).List().NewElement().Message().Interface()
			populateFieldNames(repeated, csv, includedFields, "", file, "")
			continue
		}
		if field.Kind().GoString() == "MessageKind" {
			populateFieldNames(pr.Get(field).Message().Interface(), csv, includedFields, fieldName, "", file)
		} else {
			csv.Data[file][fieldName] = make([]string, 0)
		}
	}
}

func populateBody(proto pb.Message, csv *csvInfo, parent, parentType, file string, parentID int) {
	pr := proto.ProtoReflect()
	d := pr.Descriptor()
	pFields := d.Fields()
	pName := string(d.Name())
	if file == "" {
		file = pName
	}

	//Determine if the row is empty
	nonEmpty := false

	// Populate index of parent if neccessary
	if parentType != "" {
		parentField := parentType + ".id"
		if csv.Data[file][parentField] != nil {
			csv.Data[file][parentField] = append(csv.Data[file][parentField], strconv.Itoa(parentID))
			nonEmpty = true
		}
	}

	// Populate remaining fields
	for i := 0; i < pFields.Len(); i++ {
		field := pFields.Get(i)
		fieldName := string(field.Name())
		if parent != "" {
			fieldName = parent + "." + fieldName
		}

		// Only populate required fields
		if csv.Data[file][fieldName] == nil {
			// If the field is a message or repeated then we check recursively instead
			if field.Kind().GoString() != "MessageKind" && field.Cardinality().GoString() != "Repeated" {
				continue
			}
		}
		nonEmpty = true

		// Populate repeated subfields into a seperate CSV
		// This currently only supports message subfields
		if field.Cardinality().GoString() == "Repeated" {
			repeated := pr.Get(field).List()
			for j := 0; j < repeated.Len(); j++ {
				m := repeated.Get(j).Message().Interface()
				populateBody(m, csv, "", file, "", csv.Rows[file])
			}
			continue
		}

		// Populate singular subfield into same CSV
		if field.Kind().GoString() == "MessageKind" {
			populateBody(pr.Get(field).Message().Interface(), csv, fieldName, "", file, 0)
			continue
		} else {
			csv.Data[file][fieldName] = append(csv.Data[file][fieldName], pr.Get(field).String())
		}
	}

	// Add new row if row non-empty and not a subfield
	if nonEmpty && parent == "" {
		csv.Rows[file]++

		// Populate any untouched fields, eg unrelated ids
		for fieldName, fieldValues := range csv.Data[file] {
			if len(fieldValues) < csv.Rows[file] {
				csv.Data[file][fieldName] = append(csv.Data[file][fieldName], "")
			}
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
