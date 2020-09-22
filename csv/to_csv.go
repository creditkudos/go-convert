package csv

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"sort"
	"strconv"
	"strings"

	pbV1 "github.com/golang/protobuf/proto"
	pb "google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
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
	if err := populateFieldNames(protos[0], &info, fields, "", "", ""); err != nil {
		return nil, err
	}

	for _, proto := range protos {
		if err := populateBody(proto, &info, "", "", "", 0); err != nil {
			return nil, err
		}
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

func populateFieldNames(proto pb.Message, csv *csvInfo, includedFields map[string][]string, parent, parentFile, file string) (err error) {
	pr := proto.ProtoReflect()
	d := pr.Descriptor()

	// Initialise new CSV file if necessary
	pName := string(d.Name())
	if file == "" {
		file = pName
	}
	if csv.Data[file] == nil {
		csv.Data[file] = make(map[string][]string)
	}

	// Add index of parent file if necessary
	if parentFile != "" {
		parentField := parentFile + ".id"
		if shouldIncludeField(includedFields[file], parentField) {
			csv.Data[file][parentField] = make([]string, 0)
		}
	}

	pFields := d.Fields()
	for i := 0; i < pFields.Len(); i++ {
		field := pFields.Get(i)
		fieldName := string(field.Name())

		// Append parent to name eg account.amount.currency
		if parent != "" {
			fieldName = parent + "." + fieldName
		}

		// Only include specified fields
		if !shouldIncludeField(includedFields[file], fieldName) {
			continue
		}

		// This will only convert message fields else it will error.
		if field.Cardinality() == protoreflect.Repeated {
			fieldValue := pr.Get(field)
			if field.IsList() {
				if field.Kind() == protoreflect.MessageKind {
					repeated := fieldValue.List().NewElement().Message().Interface()
					if err := populateFieldNames(repeated, csv, includedFields, "", file, ""); err != nil {
						return err
					}
					continue
				}

				// For flat field values we format as "array[0],array[1],..."
			} else if field.IsMap() {
				if field.MapValue().Kind() == protoreflect.MessageKind {
					if err := populateMapFieldNames(field, fieldValue.Map(), csv, includedFields, "", file); err != nil {
						return err
					}
					continue
				}
			}
		} else if field.Kind() == protoreflect.MessageKind {
			if err := populateFieldNames(pr.Get(field).Message().Interface(), csv, includedFields, fieldName, "", file); err != nil {
				return err
			}
			continue
		}

		csv.Data[file][fieldName] = []string{}
	}

	return nil
}

func populateMapFieldNames(field protoreflect.FieldDescriptor, m protoreflect.Map, csv *csvInfo, includedFields map[string][]string, parent, parentType string) (err error) {
	if parentType == "" {
		return fmt.Errorf("unable to populate map with empty parentType")
	}
	if field.MapValue().Kind() != protoreflect.MessageKind {
		return fmt.Errorf("unable to populate map body with non-message field values")
	}
	if field.MapKey().Kind() == protoreflect.MessageKind {
		return fmt.Errorf("unable to populate map body with message key values")
	}

	file := string(m.NewValue().Message().Descriptor().Name())
	keyColumnField := parentType + "." + string(field.Name()) + ".key"
	if csv.Data[file] == nil {
		csv.Data[file] = make(map[string][]string)
	}
	csv.Data[file][keyColumnField] = []string{}

	if err := populateFieldNames(m.NewValue().Message().Interface(), csv, includedFields, "", parentType, file); err != nil {
		return err
	}

	return nil
}

func populateBody(proto pb.Message, csv *csvInfo, parent, parentType, file string, parentID int) (err error) {
	pr := proto.ProtoReflect()
	d := pr.Descriptor()
	pName := string(d.Name())
	if file == "" {
		file = pName
	}

	//Determine if the row is empty
	nonEmpty := false

	// Populate index of parent if necessary
	if parentType != "" {
		parentField := parentType + ".id"
		if csv.Data[file][parentField] != nil {
			csv.Data[file][parentField] = append(csv.Data[file][parentField], strconv.Itoa(parentID))
			nonEmpty = true
		}
	}

	// Populate remaining fields
	pFields := d.Fields()
	for i := 0; i < pFields.Len(); i++ {
		field := pFields.Get(i)
		fieldName := string(field.Name())
		if parent != "" {
			fieldName = parent + "." + fieldName
		}

		// Only populate required fields
		if csv.Data[file][fieldName] == nil {
			// Message fields are checked recursively
			if field.Kind() != protoreflect.MessageKind {
				continue
			}
		}
		nonEmpty = true

		// Populate repeated subfields into a seperate CSV
		// This currently only supports message subfields
		if field.Cardinality() == protoreflect.Repeated {
			fieldValue := pr.Get(field)
			if field.IsList() {
				if field.Kind() != protoreflect.MessageKind {
					csv.Data[file][fieldName] = append(csv.Data[file][fieldName], protoListToString(fieldValue.List()))
				} else {
					repeated := fieldValue.List()
					for j := 0; j < repeated.Len(); j++ {
						m := repeated.Get(j).Message().Interface()
						if err := populateBody(m, csv, "", file, "", csv.Rows[file]); err != nil {
							return err
						}
					}
				}
			} else if field.IsMap() {
				if field.MapKey().Kind() != protoreflect.MessageKind && field.MapValue().Kind() != protoreflect.MessageKind {
					// Map fields are stored as message fields so require an additional check
					if csv.Data[file][fieldName] == nil {
						continue
					}
					csv.Data[file][fieldName] = append(csv.Data[file][fieldName], protoMapToString(fieldValue.Map()))
				} else {
					if err := populateMapBody(field, fieldValue.Map(), csv, file, csv.Rows[file]); err != nil {
						return err
					}
				}
			}

			continue
		}

		// Populate singular subfield into same CSV
		if field.Kind() == protoreflect.MessageKind {
			if err := populateBody(pr.Get(field).Message().Interface(), csv, fieldName, "", file, 0); err != nil {
				return err
			}

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

	return nil
}

func populateMapBody(field protoreflect.FieldDescriptor, m protoreflect.Map, csv *csvInfo, parentType string, parentID int) (err error) {
	if parentType == "" {
		return fmt.Errorf("unable to populate map with empty parentType")
	}
	if field.MapValue().Kind() != protoreflect.MessageKind {
		return fmt.Errorf("unable to populate map body with non-message field values")
	}
	if field.MapKey().Kind() == protoreflect.MessageKind {
		return fmt.Errorf("unable to populate map body with message key values")
	}

	file := string(m.NewValue().Message().Descriptor().Name())
	keyColumnField := parentType + "." + string(field.Name()) + ".key"

	m.Range(func(key protoreflect.MapKey, value protoreflect.Value) bool {
		csv.Data[file][keyColumnField] = append(csv.Data[file][keyColumnField], key.String())

		if err := populateBody(value.Message().Interface(), csv, "", parentType, file, parentID); err != nil {
			return false
		}

		return true
	})

	return nil
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

func shouldIncludeField(array []string, value string) bool {
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

func protoListToString(a protoreflect.List) string {
	outputString := []string{}

	for i := 0; i < a.Len(); i++ {
		outputString = append(outputString, a.Get(i).String())
	}

	return strings.Join(outputString, ",")
}

func protoMapToString(a protoreflect.Map) string {
	stringMap := make(map[string]string)
	fields := []string{}

	a.Range(func(k protoreflect.MapKey, v protoreflect.Value) bool {
		ks := k.String()
		stringMap[ks] = v.String()
		fields = append(fields, ks)
		return true
	})

	// Obtain consistent ordering
	sort.Strings(fields)
	mapString := []string{}
	for _, key := range fields {
		mapString = append(mapString, key+":"+stringMap[key])
	}

	return strings.Join(mapString, ",")
}
