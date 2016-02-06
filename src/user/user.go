package user

import (
  "io/ioutil"
  "encoding/json"
)

type User struct {
  Name string
  Email string
  Password string
}

func List() ([]User, error) {
  jsonString, err := ioutil.ReadFile("data/users.json")
  if err != nil {
    return nil, err
  }
  jsonBlob := []byte(jsonString)
  var users []User
  perr := json.Unmarshal(jsonBlob, &users)
  if perr != nil {
    return nil, perr
  }
  return users, nil
}

func (u *User) Save() error {
  users, err := List()
  if err != nil {
    return err
  }

  users = append(users, *u)

  jsonString, perr := json.Marshal(users)
  if perr != nil {
    return perr
  }

  return ioutil.WriteFile("data/users.json", jsonString, 0600)
}
