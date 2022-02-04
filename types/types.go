package types

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"

	"gopkg.in/yaml.v3"
)

func XML_parser(fx string) XML {

	//fmt.Println("\nXML")

	xmlFile, err := os.Open(fx)

	if err != nil {
		fmt.Println(err, "X")
	}
	defer xmlFile.Close()
	byteValue, _ := ioutil.ReadAll(xmlFile)

	var file_xml XML
	xml.Unmarshal(byteValue, &file_xml)

	//for i := 0; i < len(file_xml.Service); i++ {
	//	fmt.Println("\nType: " + file_xml.Service[i].Type)
	//	fmt.Println("Name: " + file_xml.Service[i].Name)
	//	fmt.Println("Image: " + file_xml.Service[i].Image)
	//	fmt.Println(file_xml.Service[i].Env.Vars[i].Name+":", file_xml.Service[i].Env.Vars[i].Value)
	//	print("\n")
	//}

	//for i := 0; i < len(file_xml.Service); i++ {
	//	gType := reflect.TypeOf(file_xml.Service[1])
	//	strNumFields := gType.NumField()
	//	for i := 0; i < strNumFields; i++ {
	//		field := gType.Field(i)
	//		if strings.ToLower(field.Name) != "xmlname" && strings.ToLower(field.Name) != "type" {
	//			fmt.Printf("%s\n", strings.ToLower(field.Name))
	//		}
	//	}
	//	break
	//}

	//fmt.Printf(string(file_xml))

	// json encoding
	//fmt.Printf("---\njson encoding\n")
	//jsonData, err := json.Marshal(&file_xml)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(string(jsonData))

	//fmt.Println(reflect.TypeOf(file_xml.Service[0]).Name())

	return file_xml

}

func getType(myvar interface{}) string {
	return reflect.TypeOf(myvar).Name()
}

func YML_parser(fy string) map[interface{}]interface{} {

	//fmt.Println("\nYML")

	ymlFile, err := ioutil.ReadFile(fy)
	if err != nil {
		log.Fatal(err, "Y")
	}

	file_yml := make(map[interface{}]interface{})
	err2 := yaml.Unmarshal(ymlFile, &file_yml)
	if err2 != nil {
		log.Fatal(err2)
	}

	//for k_general, v_general := range file_yml {
	//	if k_general == "services" {
	//		for k_svc, v_svc := range v_general.(map[string]interface{}) {
	//			fmt.Println("\n", k_svc+":")
	//			for k_property, v_property := range v_svc.(map[string]interface{}) {
	//				if reflect.ValueOf(v_property).Kind() == reflect.Map {
	//					fmt.Println("  ", k_property+":")
	//					for k, v := range v_property.(map[string]interface{}) {
	//						fmt.Println("    ", k+":", v)
	//					}
	//				} else if reflect.ValueOf(v_property).Kind() == reflect.Slice {
	//					fmt.Println("  ", k_property+":")
	//					for _, v := range v_property.([]interface{}) {
	//						fmt.Println("    -", v)
	//					}
	//				} else {
	//					fmt.Println("  ", k_property+": ", v_property)
	//				}
	//			}
	//		}
	//	}
	//}

	//fmt.Println(reflect.TypeOf(file_yml))

	return file_yml
}

// XML
type XML struct {
	XMLName xml.Name `xml:"stack"`
	Name    string   `xml:"name,attr"`
	Service []struct {
		XMLName     xml.Name `xml:"svc"`
		Type        string   `xml:"type,attr"`
		Name        string   `xml:"name"`
		Image       string   `xml:"image"`
		Environment struct {
			XMLName xml.Name `xml:"environment"`
			Vars    []struct {
				Name  string `xml:"name"`
				Value string `xml:"value"`
			} `xml:"vars"`
		} `xml:"environment"`
		Volumes struct {
			XMLName xml.Name `xml:"volumes"`
			Volume  []struct {
				Name string `xml:"name"`
			} `xml:"volume"`
		} `xml:"volumes"`
		Ports struct {
			XMLName xml.Name `xml:"ports"`
			Port    []struct {
				Num string `xml:"num"`
			} `xml:"port"`
		} `xml:"ports"`
		Labels struct {
			XMLName xml.Name `xml:"labels"`
			Label   []struct {
				Value string `xml:"value"`
			} `xml:"label"`
		} `xml:"labels"`
	} `xml:"svc"`
}
