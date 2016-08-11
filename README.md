# gomldb
Golang interface to the Datacratic Machine Learning Database (MLDB). http://mldb.ai/
The library base on "net/http" package
# Installation
To install this library just run
```
go get -u -v github.com/ZzEeKkAa/gomldb
```
# Connections
The gomldb library includes a class called Connection. The recommended usage pattern is shown here:
```
import "github.com/ZzEeKkAa/gomldb"

mldb := new(gomldb.Connection)
mldb.Connect("http://localhost")
```
#Accessing the REST API
Once you have a connection object, you can easily make calls to the REST API:
```
resp, err := mldb.Get("/v1/types",nil)
body, _ = ioutil.ReadAll(resp.Body)
fmt.Println(string(body))
// ["datasets","functions","plugin.setups","plugin.startups","plugins","procedures"]
```
In more complicated way you can get all info you need
```
resp, err := mldb.Get("/v1/types",nil)
if err!= nil {
  fmt.Println(err)
}
else {
  fmt.Println("response Status:", resp.Status)
  fmt.Println("response Headers:", resp.Header)
  body, _ = ioutil.ReadAll(resp.Body)
  fmt.Println("response Body:", string(body))
}
```
You can send parameters in json format
```
mldb.Put("/v1/datasets/sample", []byte(`{"type": "sparse.mutable"}`)) 
```
Here we create a dataset and insert two rows of two columns into it:
```
mldb.Put( "/v1/datasets/demo", []byte(`{"type":"sparse.mutable"}`))
mldb.Post("/v1/datasets/demo/rows", []byte(`{"rowName": "first", "columns":[["a",1,0],["b",2,0]]}`))
mldb.Post("/v1/datasets/demo/rows", []byte(`{"rowName": "second", "columns":[["a",3,0],["b",4,0]]}`))
mldb.Post("/v1/datasets/demo/commit", nil)
```
# SQL Queries
Now that we have a dataset, we can use the Query() method on the connection to run an SQL query and get the results back as a string:
```
resp, _ := mldb.Query("select * from demo")
body, _ = ioutil.ReadAll(resp.Body)
fmt.Println(string(body))
// [["_rowName","a","b"],["second",3,4],["first",1,2]]
```
