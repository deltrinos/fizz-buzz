# fizz-buzz

## Environment variables

It starts a http server on :3000 by default (environment variable ADDR for modify).

PRODUCTION environment variable is available for switch off debug logs.

## API

### GET /stats

show statistics endpoint allowing users to know what the most frequent request

### GET /fizz

Returns a list of strings with numbers from 1 to limit, where: all multiples of int1 are replaced by str1, all multiples of int2 are replaced by str2, all multiples of int1 and int2 are replaced by str1str2.

| Parameter | Mandatory     |  
| ------- |:--------------:| 
| int1    | Yes            |  
| int2    | Yes            |   
| limit   | Yes            | 
| str1    | Yes            |  
| str2    | Yes            |    

Examples :
```
$ curl 'http://localhost:3000/fizz?int1=3&int2=5&limit=15&str1=fizz&str2=buzz'
["1","2","fizz","4","buzz","fizz","7","8","fizz","buzz","11","fizz","13","14","fizzbuzz"]
```
