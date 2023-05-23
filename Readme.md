### В PowerShell:

Invoke-WebRequest -Uri "http://localhost:8080/token" -Method POST -Body "client_id=my_client_id&client_secret=my_client_secret"

### Ответ

StatusCode        : 200

StatusDescription : OK

Content           : {"access_token":"FWdD07QTfcA2pIikIA2X","token_type":"Bearer"}

RawContent        : HTTP/1.1 200 OK

Content-Length: 61

Content-Type: application/json; charset=utf-8

Date: Tue, 23 May 2023 15:44:43 GMT

                    {"access_token":"FWdD07QTfcA2pIikIA2X","token_type":"Bearer"}
Forms             : {}

Headers           : {[Content-Length, 61], [Content-Type, application/json; charset=utf-8], [Date, Tue, 23 May 2023 15:
44:43 GMT]}

Images            : {}

InputFields       : {}

Links             : {}

ParsedHtml        : mshtml.HTMLDocumentClass

RawContentLength  : 61