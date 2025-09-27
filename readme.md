# C.A.C.A
vim config
switch between searching files grepping or all the other functions
is kinda annoying to close and reopen telescope each time

Create
A
Cool
App

![caca](https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQYF-iq3jqMvvE4CIDoOdxVYRIW1xmmjkZiRg&s) 
cool shit

make a mini framework to create directories easier like

mkproject(
  {
    "steins_gate": {
    "cmd" :{
      "steins_gate" : {
        "main.go" : `func main(){fmt.Println("hello")}`
        "logger.go" : `func main(){fmt.Println("hello")}`
        }
      }
    }
  }
)
and extract this from a toml file cuz json cannot be used for evilness

func (r *root) UnmarshalJSON(data []byte) error {
	var tmp map[string]any
	err := json.Unmarshal(data, &tmp)
	if err != nil {
		return err
	}

	parent := tmp["parent"].(map[string]any)
	nodes := parent["nodes"].([]interface{})
	fmt.Println(nodes)

	for _, n := range nodes {
		f := n.(map[string]any)
		fmt.Println(f["name"])
		fmt.Println(f["content"])
	}

	return nil
}

func interfaceDebugMarshal() {

	d := directory{
		Name: "utils",
		Nodes: []node{
			&file{
				Name:    "helpers.go",
				Content: "func help(){}",
			},
			&file{
				Name:    "main.go",
				Content: "func main() {fmt.Println(hellope);}",
			},
		},
	}

	r := root{
		ProjectName: "steins_gate",
		Parent:      &d,
	}

	b := make([]byte, 1024)
	b, err := yaml.Marshal(r)
	logErr(err)
	err = os.WriteFile("dolls.yaml", b, os.ModePerm)
	logErr(err)

}

func AssertFile(data map[string]interface{}) file {
	var f file
	val := reflect.ValueOf(f)
	for i := 0; i < val.NumField(); i++ {
		tag := val.Type().Field(i).Tag.Get("json")
		fmt.Println(data[tag])
	}
	return file{}
}

func interfaceSilly() {
	nodes := []interface{}{
		map[string]interface{}{
			"content": "func help(){}",
			"name":    "helpers.go",
		},
		map[string]interface{}{
			"content": "func main() {fmt.Println(hellope);}",
			"name":    "main.go"},
	}

	fmt.Println(nodes)

	// for _, inf := range nodes {
	// 	json.Marshal()
	// 	b := inf.(file)
	// 	fmt.Println(b)
	// }

}
