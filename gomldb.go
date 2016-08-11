package gomldb

import (
    "net/http"
    "bytes"
    "strings"
    "errors"
    "net/url"
)


type Connection struct {
    uri string
}

func (conn *Connection) Connect(host string) error{
    if strings.HasPrefix(host,"http") {
        conn.uri = strings.TrimSuffix(host,"/")
        return nil
    } else {
        return errors.New("URIs must start with 'http'")
    }
}

func (conn *Connection) Get(url string, json []byte) (*http.Response, error) {
    req, err := http.NewRequest("GET", conn.uri+url, bytes.NewBuffer(json))
    if err != nil {
        return nil,err
    }
    req.Header.Set("Content-Type", "application/json")
    client := &http.Client{}
    return client.Do(req)
}

func (conn *Connection) Post(url string, json []byte) (*http.Response, error) {
    req, err := http.NewRequest("POST", conn.uri+url, bytes.NewBuffer(json))
    if err != nil {
        return nil,err
    }
    req.Header.Set("Content-Type", "application/json")
    client := &http.Client{}
    return client.Do(req)
}

func (conn *Connection) Put(url string, json []byte) (*http.Response, error) {
    req, err := http.NewRequest("PUT", conn.uri+url, bytes.NewBuffer(json))
    if err != nil {
        return nil,err
    }
    req.Header.Set("Content-Type", "application/json")
    client := &http.Client{}
    return client.Do(req)
}

func (conn *Connection) Delete(url string) (*http.Response, error) {
    req, err := http.NewRequest("DELETE", conn.uri+url, nil)
    if err != nil {
        return nil,err
    }
    req.Header.Set("Content-Type", "application/json")
    client := &http.Client{}
    return client.Do(req)
}

func (conn *Connection) Query(sql string) (*http.Response, error) {
    return http.Get(conn.uri+"/v1/query?"+url.Values{"q":{sql},"format":{"table"}}.Encode())
}
