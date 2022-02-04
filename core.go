package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"reflect"
	"strings"

	"./types"
)

func parser(xml string, yml string, out string) {
	fmt.Println("\nParser")

	f_xml := types.XML_parser(xml)
	f_yml := types.YML_parser(yml)

	file, err := os.Create(string(out))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	file.WriteString("version: '2'\nservices:\n")

	for i := 0; i < len(f_xml.Service); i++ { // XML
		for k_general, v_general := range f_yml { // YML
			if k_general == "services" {
				for k_svc, v_svc := range v_general.(map[string]interface{}) {
					if f_xml.Service[i].Name == k_svc { // Node
						fmt.Println("\n", k_svc+":")
						c := fmt.Sprintf("  %s:\n", k_svc) // Compose - Service
						file.WriteString(c)
						for k_property, v_property := range v_svc.(map[string]interface{}) { // Properties
							// XML
							for s := 0; s < len(f_xml.Service); s++ {
								gType := reflect.TypeOf(f_xml.Service[s])
								strNumFields := gType.NumField()
								for j := 0; j < strNumFields; j++ {
									field := gType.Field(j).Name
									if strings.ToLower(field) != "xmlname" && strings.ToLower(field) != "type" {
										if strings.ToLower(field) == k_property || v_property == nil {
											if v_property != nil { // YAML
												if reflect.ValueOf(v_property).Kind() == reflect.Map {
													//if len(f_xml.Service[i].Volumes.Volume) != 0 {
													fmt.Println("  ", k_property+":")
													c := fmt.Sprintf("    %s:\n", k_property) // COMPOSE PROPERTY
													file.WriteString(c)
													for p := 0; p < len(f_xml.Service[i].Environment.Vars); p++ {
														fmt.Println("    ", f_xml.Service[i].Environment.Vars[p].Name+":", f_xml.Service[i].Environment.Vars[p].Value)
														c := fmt.Sprintf("      %s: %s\n", f_xml.Service[i].Environment.Vars[p].Name, f_xml.Service[i].Environment.Vars[p].Value)
														file.WriteString(c)
														//}
													}
												} else if reflect.ValueOf(v_property).Kind() == reflect.Slice {
													//fmt.Println("  ", k_property+":")
													//c := fmt.Sprintf("    %s:\n", k_property)
													//file.WriteString(c)
													if k_property == "volumes" {
														if len(f_xml.Service[i].Volumes.Volume) != 0 {
															fmt.Println("  ", k_property+":")
															c := fmt.Sprintf("    %s:\n", k_property)
															file.WriteString(c)
															for p := 0; p < len(f_xml.Service[i].Volumes.Volume); p++ {
																if f_xml.Service[i].Volumes.Volume[p].Name != " " {
																	fmt.Println("    -", f_xml.Service[i].Volumes.Volume[p].Name)
																	c := fmt.Sprintf("       - %s\n", f_xml.Service[i].Volumes.Volume[p].Name)
																	file.WriteString(c)
																}
															}
														}
													} else if k_property == "ports" {
														if len(f_xml.Service[i].Ports.Port) != 0 {
															fmt.Println("  ", k_property+":")
															c := fmt.Sprintf("    %s:\n", k_property)
															file.WriteString(c)
															for p := 0; p < len(f_xml.Service[i].Ports.Port); p++ {
																fmt.Println("    -", f_xml.Service[i].Ports.Port[p].Num)
																c := fmt.Sprintf("       - %s\n", f_xml.Service[i].Ports.Port[p].Num)
																file.WriteString(c)
															}
														}
													} else if k_property == "labels" {
														if len(f_xml.Service[i].Labels.Label) != 0 {
															fmt.Println("  ", k_property+":")
															c := fmt.Sprintf("    %s:\n", k_property)
															file.WriteString(c)
															for p := 0; p < len(f_xml.Service[i].Labels.Label); p++ {
																fmt.Println("    -", f_xml.Service[i].Labels.Label[p].Value)
																c := fmt.Sprintf("       - %s\n", f_xml.Service[i].Labels.Label[p].Value)
																file.WriteString(c)
															}
														}
													} else if reflect.ValueOf(v_property).Kind() == reflect.Invalid {
														fmt.Println(reflect.ValueOf(v_property).Kind())
														fmt.Println(" ", k_property, v_property)
														//continue
													} //else { fmt.Println("Pointer") continue }
												} else { // Other Property
													//if len(f_xml.Service[i].Image) != 0 {
													if k_property == "image" {
														fmt.Println("  ", k_property+":", f_xml.Service[i].Image)
														c := fmt.Sprintf("    %s: %s\n", k_property, f_xml.Service[i].Image)
														file.WriteString(c)
													}
													//}
												}
											} else { // XML || nils
												//continue
												if strings.ToLower(field) == "volumes" {
													//if len(f_xml.Service[i].Volumes.Volume) != 0 {
													fmt.Println("  ", strings.ToLower(field)+":")
													c := fmt.Sprintf("    %s:\n", strings.ToLower(field))
													file.WriteString(c)
													for p := 0; p < len(f_xml.Service[i].Volumes.Volume); p++ {
														if f_xml.Service[i].Volumes.Volume[p].Name != " " {
															fmt.Println("    -", f_xml.Service[i].Volumes.Volume[p].Name)
															c := fmt.Sprintf("       - %s\n", f_xml.Service[i].Volumes.Volume[p].Name)
															file.WriteString(c)
														}
													}
													//}
												} else if strings.ToLower(field) == "ports" {
													//if len(f_xml.Service[i].Ports.Port) != 0 {
													fmt.Println("  ", strings.ToLower(field)+":")
													c := fmt.Sprintf("    %s:\n", strings.ToLower(field))
													file.WriteString(c)
													for p := 0; p < len(f_xml.Service[i].Ports.Port); p++ {
														fmt.Println("    -", f_xml.Service[i].Ports.Port[p].Num)
														c := fmt.Sprintf("       - %s\n", f_xml.Service[i].Ports.Port[p].Num)
														file.WriteString(c)
													}
													//}
												} else if strings.ToLower(field) == "environment" {
													//if len(f_xml.Service[i].Environment.Vars) != 0 {
													fmt.Println("  ", strings.ToLower(field)+":")
													c := fmt.Sprintf("    %s:\n", strings.ToLower(field))
													file.WriteString(c)
													for p := 0; p < len(f_xml.Service[i].Environment.Vars); p++ {
														fmt.Println("    ", f_xml.Service[i].Environment.Vars[p].Name+":", f_xml.Service[i].Environment.Vars[p].Value)
														c := fmt.Sprintf("      %s: %s\n", f_xml.Service[i].Environment.Vars[p].Name, f_xml.Service[i].Environment.Vars[p].Value)
														file.WriteString(c)
													}
													//}
												} else if strings.ToLower(field) == "labels" {
													if len(f_xml.Service[i].Labels.Label) != 0 {
														fmt.Println("  ", strings.ToLower(field)+":")
														c := fmt.Sprintf("    %s:\n", strings.ToLower(field))
														file.WriteString(c)
														for p := 0; p < len(f_xml.Service[i].Labels.Label); p++ {
															fmt.Println("    -", f_xml.Service[i].Labels.Label[p].Value)
															c := fmt.Sprintf("       - %s\n", f_xml.Service[i].Labels.Label[p].Value)
															file.WriteString(c)
														}
													}
												} // else { fmt.Println("Pointer") }
											}
											continue
										}
										//break
									}
								}
								break
							}
						}
					} else { // XML
						fmt.Println("\n", f_xml.Service[i].Name+":")
						c := fmt.Sprintf("  %s:\n", f_xml.Service[i].Name)
						file.WriteString(c)
						for s := 0; s < len(f_xml.Service); s++ {
							gType := reflect.TypeOf(f_xml.Service[s])
							strNumFields := gType.NumField()
							for j := 0; j < strNumFields; j++ {
								field := gType.Field(j).Name
								if strings.ToLower(field) != "xmlname" && strings.ToLower(field) != "type" {
									if strings.ToLower(field) == "volumes" {
										if len(f_xml.Service[i].Volumes.Volume) != 0 {
											fmt.Println("  ", strings.ToLower(field)+":")
											c := fmt.Sprintf("    %s:\n", strings.ToLower(field))
											file.WriteString(c)
											for p := 0; p < len(f_xml.Service[i].Volumes.Volume); p++ {
												if f_xml.Service[i].Volumes.Volume[p].Name != " " {
													fmt.Println("    -", f_xml.Service[i].Volumes.Volume[p].Name)
													c := fmt.Sprintf("       - %s\n", f_xml.Service[i].Volumes.Volume[p].Name)
													file.WriteString(c)
												}
											}
										}
									} else if strings.ToLower(field) == "image" {
										//if len(f_xml.Service[i].Image) != 0 {
										fmt.Println("  ", strings.ToLower(field)+":", f_xml.Service[i].Image)
										c := fmt.Sprintf("    %s: %s\n", strings.ToLower(field), f_xml.Service[i].Image)
										file.WriteString(c)
										//}
									} else if strings.ToLower(field) == "ports" {
										if len(f_xml.Service[i].Ports.Port) != 0 {
											fmt.Println("  ", strings.ToLower(field)+":")
											c := fmt.Sprintf("    %s:\n", strings.ToLower(field))
											file.WriteString(c)
											for p := 0; p < len(f_xml.Service[i].Ports.Port); p++ {
												fmt.Println("    -", f_xml.Service[i].Ports.Port[p].Num)
												c := fmt.Sprintf("       - %s\n", f_xml.Service[i].Ports.Port[p].Num)
												file.WriteString(c)
											}
										}
									} else if strings.ToLower(field) == "environment" {
										//if len(f_xml.Service[i].Environment.Vars) != 0 {
										fmt.Println("  ", strings.ToLower(field)+":")
										c := fmt.Sprintf("    %s:\n", strings.ToLower(field))
										file.WriteString(c)
										for p := 0; p < len(f_xml.Service[i].Environment.Vars); p++ {
											fmt.Println("    ", f_xml.Service[i].Environment.Vars[p].Name+":"+f_xml.Service[i].Environment.Vars[p].Value)
											c := fmt.Sprintf("      %s: %s\n", f_xml.Service[i].Environment.Vars[p].Name, f_xml.Service[i].Environment.Vars[p].Value)
											file.WriteString(c)
										}
										//}
									} else if strings.ToLower(field) == "labels" {
										//fmt.Println(len(f_xml.Service[i].Labels.Label))
										if len(f_xml.Service[i].Labels.Label) != 0 {
											fmt.Println("  ", strings.ToLower(field)+":")
											c := fmt.Sprintf("    %s:\n", strings.ToLower(field))
											file.WriteString(c)
											for p := 0; p < len(f_xml.Service[i].Labels.Label); p++ {
												fmt.Println("    -", f_xml.Service[i].Labels.Label[p].Value)
												c := fmt.Sprintf("       - %s\n", f_xml.Service[i].Labels.Label[p].Value)
												file.WriteString(c)
											}
										}
									}
								}
							}
							break
						}
					}
					break
				} // } else { //	fmt.Println("YML ERROR")
			}
		}
		if len(f_yml) == 0 {
			fmt.Println("\n\tYML ERROR FILE")
			break
		}
	}
	if len(f_xml.Service) == 0 {
		fmt.Println("\n\tXML ERROR FILE")
	}
	fmt.Println(" ")
}

func help() {
	help := `Rancher Template Compose - Continues Delivery
	
	[*] AYUDA

	Uso: core -xml [XML] -yml [YML] -out [OUT]

	`
	fmt.Println("\n\t" + help)
}

func Path(file string) string {
	f := "echo $(pwd)/" + file
	cmd := exec.Command("bash", "-c", f)
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
		return "FAILED"
	}
	// Print the output
	return string(stdout)
}

func main() {

	if len(os.Args) == 7 {
		if os.Args[1] == "-xml" && os.Args[3] == "-yml" {
			if Path(os.Args[2]) != "FAILED" && Path(os.Args[4]) != "FAILED" && Path(os.Args[2]) != "FAILED" {
				//fmt.Println(Path(os.Args[2]), Path(os.Args[4]), Path(os.Args[6]))
				parser(os.Args[2], os.Args[4], os.Args[6])
			}
		}
	} else {
		help()
	}
}
