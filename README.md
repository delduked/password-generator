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

## Example curl requests

### GO
```go
  url := "http://localhost:8080/generateParams?length=32&lower=true&upper=true&number=true&special=true"
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
```
### Javascript
```javascript
var http = require('follow-redirects').http;
var fs = require('fs');

var options = {
  'method': 'GET',
  'hostname': 'localhost',
  'port': 8080,
  'path': '/generateParams?length=32&lower=true&upper=true&number=true&special=true',
  'headers': {
  },
  'maxRedirects': 20
};

var req = http.request(options, function (res) {
  var chunks = [];

  res.on("data", function (chunk) {
    chunks.push(chunk);
  });

  res.on("end", function (chunk) {
    var body = Buffer.concat(chunks);
    console.log(body.toString());
  });

  res.on("error", function (error) {
    console.error(error);
  });
});

req.end();
```
### Curl
```curl
curl --location --request GET 'http://localhost:8080/generateParams?length=32&lower=true&upper=true&number=true&special=true'
```