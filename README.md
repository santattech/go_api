# REST API with GO
Perform CRUD operations, but CUD is in progress.

### Prerequisite
You will need to install GO. Check your GO version.
```
 ~/dev/go_api$ go version
go version go1.16.5 linux/amd64

```
### API Details

1. #### R
```
curl --request GET \
  --url http://localhost:8080/api/articles \
  
```

```
curl --request GET \
  --url http://localhost:10000/articles/1 \
  
```

2. #### C
```
curl --request POST \
  --url http://localhost:10000/article \
  --header 'Content-Type: application/json' \
  --data '{
    "Id": "3", 
    "Title": "Newly Created Post", 
    "desc": "The description for my new post", 
    "content": "my articles content" 
}'

```

3. #### D

```
curl --request DELETE \
  --url http://localhost:10000/articles/1 \
  
```