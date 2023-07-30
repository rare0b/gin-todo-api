## CreateTask taskを作成
curl.exe -X POST -d '"{\"title\": \"test1\"}"' http://localhost:8000/todo
curl.exe -X POST -d '"{\"title\": \"test2\"}"' http://localhost:8000/todo

[GIN] 2023/07/31 - 00:59:17 | 204 |      8.5157ms |       127.0.0.1 | POST     "/todo"
[GIN] 2023/07/31 - 00:59:31 | 204 |     17.6412ms |       127.0.0.1 | POST     "/todo"

## GetAll すべてのtodoを取得
curl http://localhost:8000/todo

[GIN] 2023/07/31 - 00:59:44 | 200 |      6.4054ms |             ::1 | GET      "/todo"

StatusCode        : 200
StatusDescription : OK
Content           : [{"ID":1,"CreatedAt":"2023-07-31T00:59:17.037545+09:00","UpdatedAt":"2023-07-31T00:59:17.037545+09:
                    00","DeletedAt":null,"Title":"test1"},{"ID":2,"CreatedAt":"2023-07-31T00:59:31.402074+09:00","Updat
                    ed...
RawContent        : HTTP/1.1 200 OK
                    Content-Length: 273
                    Content-Type: application/json; charset=utf-8
                    Date: Sun, 30 Jul 2023 15:59:44 GMT

                    [{"ID":1,"CreatedAt":"2023-07-31T00:59:17.037545+09:00","UpdatedAt":"2023-07...
Forms             : {}
Headers           : {[Content-Length, 273], [Content-Type, application/json; charset=utf-8], [Date, Sun, 30 Jul 2023 15
                    :59:44 GMT]}
Images            : {}
InputFields       : {}
Links             : {}
ParsedHtml        : mshtml.HTMLDocumentClass
RawContentLength  : 273

## UpdateTask taskを編集、GETを実行しUpdatedAtが更新されていることも確認
**id:1について、titleがtest1→test3に更新**
curl.exe -X PUT -d '"{\"title\": \"test3\"}"' http://localhost:8000/todo/edit/1
curl http://localhost:8000/todo

[GIN] 2023/07/31 - 01:00:18 | 204 |      13.555ms |       127.0.0.1 | PUT      "/todo/edit/1"
[GIN] 2023/07/31 - 01:00:38 | 200 |      2.8636ms |             ::1 | GET      "/todo"

StatusCode        : 200
StatusDescription : OK
Content           : [{"ID":1,"CreatedAt":"2023-07-31T00:59:17.037545+09:00","UpdatedAt":"2023-07-31T01:00:18.110072+09:
                    00","DeletedAt":null,"Title":"test3"},{"ID":2,"CreatedAt":"2023-07-31T00:59:31.402074+09:00","Updat
                    ed...
RawContent        : HTTP/1.1 200 OK
                    Content-Length: 273
                    Content-Type: application/json; charset=utf-8
                    Date: Sun, 30 Jul 2023 16:00:38 GMT

                    [{"ID":1,"CreatedAt":"2023-07-31T00:59:17.037545+09:00","UpdatedAt":"2023-07...
Forms             : {}
Headers           : {[Content-Length, 273], [Content-Type, application/json; charset=utf-8], [Date, Sun, 30 Jul 2023 16
                    :00:38 GMT]}
Images            : {}
InputFields       : {}
Links             : {}
ParsedHtml        : mshtml.HTMLDocumentClass
RawContentLength  : 273

## UpdateTask idに一致するレコードがない場合
curl.exe -X PUT -d '"{\"title\": \"test3\"}"' http://localhost:8000/todo/edit/5

[GIN] 2023/07/31 - 01:01:38 | 404 |      6.8638ms |       127.0.0.1 | PUT      "/todo/edit/5"

{}

## DeleteTask タスクを削除
curl.exe -X DELETE http://localhost:8000/todo/delete/2
curl http://localhost:8000/todo

[GIN] 2023/07/31 - 01:02:54 | 204 |     18.3499ms |       127.0.0.1 | DELETE   "/todo/delete/2"
[GIN] 2023/07/31 - 01:03:09 | 200 |     12.7057ms |             ::1 | GET      "/todo"

StatusCode        : 200
StatusDescription : OK
Content           : [{"ID":1,"CreatedAt":"2023-07-31T00:59:17.037545+09:00","UpdatedAt":"2023-07-31T01:00:18.110072+09:
                    00","DeletedAt":null,"Title":"test3"}]
RawContent        : HTTP/1.1 200 OK
                    Content-Length: 137
                    Content-Type: application/json; charset=utf-8
                    Date: Sun, 30 Jul 2023 16:03:09 GMT

                    [{"ID":1,"CreatedAt":"2023-07-31T00:59:17.037545+09:00","UpdatedAt":"2023-07...
Forms             : {}
Headers           : {[Content-Length, 137], [Content-Type, application/json; charset=utf-8], [Date, Sun, 30 Jul 2023 16
                    :03:09 GMT]}
Images            : {}
InputFields       : {}
Links             : {}
ParsedHtml        : mshtml.HTMLDocumentClass
RawContentLength  : 137

## DeleteTask idに一致するレコードがない場合
curl.exe -X DELETE http://localhost:8000/todo/delete/5

[GIN] 2023/07/31 - 01:03:37 | 404 |      5.4535ms |       127.0.0.1 | DELETE   "/todo/delete/5"

{}
