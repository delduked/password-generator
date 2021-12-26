# GO Password generator

Go API for generating passwords.

## Getting started

To get started clone this repo and install the necessary dependencies with the following commands.

```
git clone https://gitlab.com/alienate/password-generator.git
cd password-generator
go mod download
```
Once you have moved to the cloned directory from the previous and downloaded the dependencies with the command `go mod download` you can run the API with the below command:
```
go run ./main.go
```

## GO
### Health check
```go
package main

import (
  "fmt"
  "net/http"
  "io/ioutil"
)

func main() {

  url := "http://localhost:8080/healthcheck"
  method := "GET"

  client := &http.Client {
  }
  req, err := http.NewRequest(method, url, nil)

  if err != nil {
    fmt.Println(err)
    return
  }
  res, err := client.Do(req)
  if err != nil {
    fmt.Println(err)
    return
  }
  defer res.Body.Close()

  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(string(body))
}
```



## Javascript
### Health Check
```javascript
var requestOptions = {
  method: 'GET',
  redirect: 'follow'
};

fetch("http://localhost:8080/healthcheck", requestOptions)
  .then(response => response.text())
  .then(result => console.log(result))
  .catch(error => console.log('error', error));
```

### Generate new random password with POST
```javascript
var myHeaders = new Headers();
myHeaders.append("Content-Type", "application/json");

var raw = JSON.stringify({
  "length": 16,
  "lower": true,
  "upper": true,
  "number": true,
  "special": true
});

var requestOptions = {
  method: 'POST',
  headers: myHeaders,
  body: raw,
  redirect: 'follow'
};

fetch("http://localhost:8080/pw", requestOptions)
  .then(response => response.text())
  .then(result => console.log(result))
  .catch(error => console.log('error', error));
```

## Curl
### Health Check
```curl
curl --location --request GET 'http://localhost:8080/healthcheck'
```

### Generate new password with POST request
```curl
curl --location --request POST 'http://localhost:8080/pw' \
--header 'Content-Type: application/json' \
--data-raw '{
    "length": 16,
    "lower": true,
    "upper": true,
    "number": true,
    "special": true
}'
```
### Generate new password with paramaters
```curl
curl --location --request GET 'http://localhost:8080/pw?length=16&lower=true&upper=true&number=true&special=true'
```
### Get a specific password field
```curl
curl --location --request GET 'http://localhost:8080/db/e7556ee2-bb8f-4dac-ad48-865a733f2a90'
```
### Get all saved password fields
```curl
curl --location --request GET 'http://localhost:8080/db'
```
### Save a new password field
```curl
curl --location --request POST 'http://localhost:8080/db' \
--header 'Content-Type: application/json' \
--data-raw '{
    "Account":"newpassword",
    "Username":"newpassword",
    "Password":"newpassword"
}'
```
### Update an existing password field
```curl
curl --location --request PATCH 'http://localhost:8080/db' \
--header 'Content-Type: application/json' \
--data-raw '{
    "Key":"e7556ee2-bb8f-4dac-ad48-865a733f2a90",
    "Account":"update",
    "Username":"update",
    "Password":"update"
}'
```
### Delete an existing password field
```curl
curl --location --request DELETE 'http://localhost:8080/db/e7556ee2-bb8f-4dac-ad48-865a733f2a90'
```