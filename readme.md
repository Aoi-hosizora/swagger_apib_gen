# swagger_apib_gen

+ Auto generate swagger and apib restful api document

### Usage

+ `gen_swagger`

```bash
python3 gen_swagger.py -m ./demo/main.go -o ./demo/swagger.yaml -e go
```

```
usage: gen_swagger.py [-h] -m MAIN -o OUTPUT [-e [EXT [EXT ...]]]

optional arguments:
  -h, --help                                   show this help message and exit
  -m MAIN, --main MAIN                         path of main file containing swagger config
  -o OUTPUT, --output OUTPUT                   path of output yaml
  -e [EXT [EXT ...]], --ext [EXT [EXT ...]]    extensions of files wanted to parse
```

+ `gen_swagger_html`

```bash
python3 gen_swagger_html.py -i ./demo/swagger.yaml -o ./demo/swagger.html
```

```
usage: gen_swagger_html.py [-h] -i INPUT -o OUTPUT

optional arguments:
  -h, --help                   show this help message and exit
  -i INPUT, --input INPUT      path of input yaml file
  -o OUTPUT, --output OUTPUT   path of output html file
```

+ `gen_apib`

```bash
python3 gen_apib.py -m ./demo/main.apib.go -o ./demo/apiary.apib -e go
```

```
usage: gen_apib.py [-h] -m MAIN -o OUTPUT [-e [EXT [EXT ...]]]

optional arguments:
  -h, --help                                   show this help message and exit
  -m MAIN, --main MAIN                         path of main file containing swagger config
  -o OUTPUT, --output OUTPUT                   path of output yaml
  -e [EXT [EXT ...]], --ext [EXT [EXT ...]]    extensions of files wanted to parse
```

### Annotaton

+ See [main.go](https://github.com/Aoi-hosizora/swagger_apib_gen/blob/master/demo/main.go) and [ctrl.go](https://github.com/Aoi-hosizora/swagger_apib_gen/blob/master/demo/ctrl.go)
+ Template only support `@Param` `@ErrorCode` `@Response`

### Format

+ Param
    + `in`: `query` `path` `header` `body` `formData`
    + `type`: `string` `integer` `number(float32)` `boolean`
    + See [Param Type](https://github.com/swaggo/swag#param-type) and [Data Type](https://github.com/swaggo/swag#data-type) 

```go
// @Param uid formData integer true "user id"
// @Param $name $in $type $required "$comment"
```

+ ErrorCode

```go
// @ErrorCode 401 unauthorized user
// @ErrorCode $code $content
```

+ Template (main)

```go
// @Template Auth
// @Template $name1 $name2 $name3
```

+ Template (controller)

```go
// @Template Auth.ErrorCode 401 unauthorized user
// @Template $name.$annotation $code $content
```

+ Tag (main)

```go
// @Tag "User" "User-Controller"
// @Tag "$name" "$description"
```

+ GlobalSecurity (Only support `apiKey`)

```go
// @GlobalSecurity Jwt Authorization header
// @GlobalSecurity $name $field $in
```

+ Accept & Produce: See [Mime Types](https://github.com/swaggo/swag#mime-types)
    + `application/json` `multipart/form-data` 
    + `text/xml` `text/plain` `text/html`
    + `image/png` `image/jpeg` `image/gif`

### References

+ [How can I control what scalar form PyYAML uses for my data?](https://stackoverflow.com/questions/8640959/how-can-i-control-what-scalar-form-pyyaml-uses-for-my-data)
