package main

import (
    "bytes"
    "fmt"
    "os"
    "strings"

    model "go-protobuf/model"
    "github.com/golang/protobuf/jsonpb"
)

func main() {
    var user1 = &model.User{
        Id:       "u001",
        Name:     "Sylvana Windrunner",
        Password: "f0r Th3 H0rD3",
        Gender:   model.UserGender_FEMALE,
    }

    // var userList = &model.UserList{
    //     List: []*model.User{
    //         user1,
    //     },
    // }

    var garage1 = &model.Garage{
        Id:   "g001",
        Name: "Kalimdor",
        Coordinate: &model.GarageCoordinate{
            Latitude:  23.2212847,
            Longitude: 53.22033123,
        },
    }

    var garageList = &model.GarageList{
        List: []*model.Garage{
            garage1,
        },
    }

    // var garageListByUser = &model.GarageListByUser{
    //     List: map[string]*model.GarageList{
    //         user1.Id: garageList,
    //     },
    // }

    // =========== original
    fmt.Printf("# ==== Original\n       %#v \n", user1)

    // =========== as string
    fmt.Printf("# ==== As String\n       %v \n", user1.String())


    // JSON string to Object proto
    // Cara 1
    var buf bytes.Buffer
    err1 := (&jsonpb.Marshaler{}).Marshal(&buf, garageList)
    if err1 != nil {
        fmt.Println(err1.Error())
        os.Exit(0)
    }
    jsonString := buf.String()
    fmt.Printf("# ==== As JSON String\n       %v \n", jsonString)


    buf2 := strings.NewReader(jsonString)
    protoObject := new(model.GarageList)

    err2 := (&jsonpb.Unmarshaler{}).Unmarshal(buf2, protoObject)
    if err2 != nil {
        fmt.Println(err2.Error())
        os.Exit(0)
    }

    fmt.Printf("# ==== As String\n       %v \n", protoObject.String())

    // Cara 2
    protoObject2 := new(model.GarageList)

    err3 := jsonpb.UnmarshalString(jsonString, protoObject2)
    if err3 != nil {
        fmt.Println(err3.Error())
        os.Exit(0)
    }

    fmt.Printf("# ==== As String\n       %v \n", protoObject2.String())
}
