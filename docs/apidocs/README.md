# API Docs

## Criar novo foo
```http
###
POST https://api.qbem.net.br/foos
Content-Type: application/json; charset=utf-8

{
    "foo":"barz"
}


> HTTP/1.1 201 Created
> Content-Type: application/json; charset=utf-8
```

## Consulta pelo slug
```http
###
GET https://api.qbem.net.br/foos/{{slug}}
Accept: application/json; charset=utf-8

> HTTP/1.1 200 Created
> Content-Type: application/json; charset=utf-8
> {
    
}
```