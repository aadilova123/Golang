package main
import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
)

func main() {
    // Open jsonFile
    jsonFile, err := os.Open("students.json")
    // if we os.Open returns an error then handle it
    if err != nil {
        fmt.Println(err)
    }

    fmt.Println("Successfully Opened students.json")
    // defer the closing of our jsonFile so that we can parse it later on
    defer jsonFile.Close()

    // read our opened xmlFile as a byte array.
    byteValue, _ := ioutil.ReadAll(jsonFile)

    // we initialize our Students array
    var students Students

    // we unmarshal our byteArray which contains our
    // jsonFile's content into 'students' which we defined above
    json.Unmarshal(byteValue, &students)

    
    for i := 0; i < len(students.Students); i++ {
        fmt.Println("Student id: " + students.Students[i].Id)      
        fmt.Println("Student Faculty: " + students.Students[i].Faculty)
		fmt.Println("Student Full name: " + students.Students[i].FullName)
		fmt.Println("Student Gpa: " + students.Students[i].Gpa)
		fmt.Println("Student gender: " + students.Students[i].Gender)
		fmt.Println(students.Students[i].Age)
        fmt.Println("Facebook Url: " + students.Students[i].Social.Facebook)
    }

}