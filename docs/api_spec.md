# Customer API Spec

## Get Customer
  
- Endpoint : GET /api/cutomers/{customer_id}

Response Body (Success) 200 :

```json
{
    "keluarga": [
        {
            "hubungan": "istri",
            "nama": "istri saitama",
            "tanggal_lahir": "1999-10-21"
        },
        {
            "hubungan": "anak",
            "nama": "anak saitama",
            "tanggal_lahir": "2020-12-24"
        }
    ],
    "nama": "nama saya",
    "tanggal_lahir": "1998",
    "telepon": "62891726182",
    "kewarganegaraan": "Indonesia (ID)",
    "email": "nama@gmail.com"
}
```

Response Body (Failed) 404 :
```json
{
    "message": "customer not found"
}
```


## Create Customer
  
- Endpoint : POST /api/cutomers/

Request Body :

```json
{
    "keluarga": [
        {
            "hubungan": "istri",
            "nama": "istri saitama",
            "tanggal_lahir": "1999-10-21"
        },
        {
            "hubungan": "anak",
            "nama": "anak saitama",
            "tanggal_lahir": "2020-12-24"
        }
    ],
    "nama": "nama saya",
    "tanggal_lahir": "1998",
    "telepon": "62891726182",
    "kewarganegaraan": "Indonesia (ID)",
    "email": "nama@gmail.com"
}
```

Response Body (Success) 200 :

```json
{
    "keluarga": [
        {
            "hubungan": "istri",
            "nama": "istri saitama",
            "tanggal_lahir": "1999-10-21"
        },
        {
            "hubungan": "anak",
            "nama": "anak saitama",
            "tanggal_lahir": "2020-12-24"
        }
    ],
    "nama": "nama saya",
    "tanggal_lahir": "1998",
    "telepon": "62891726182",
    "kewarganegaraan": "Indonesia (ID)",
    "email": "nama@gmail.com"
}
```

Response Body (Failed) 400 :
```json
{
    "message": "field cannot blank"
}
```


## Update Customer
  
- Endpoint : PUT /api/cutomers/{customer_id}

Request Body :

```json
{
    "keluarga": [
        {
            "hubungan": "istri",
            "nama": "istri saitama",
            "tanggal_lahir": "1999-10-21"
        },
        {
            "hubungan": "anak",
            "nama": "anak saitama",
            "tanggal_lahir": "2020-12-24"
        }
    ],
    "nama": "nama saya",
    "tanggal_lahir": "1998",
    "telepon": "62891726182",
    "kewarganegaraan": "Indonesia (ID)",
    "email": "nama@gmail.com"
}
```

Response Body (Success) 200 :

```json
{
    "keluarga": [
        {
            "hubungan": "istri",
            "nama": "istri saitama",
            "tanggal_lahir": "1999-10-21"
        },
        {
            "hubungan": "anak",
            "nama": "anak saitama",
            "tanggal_lahir": "2020-12-24"
        }
    ],
    "nama": "nama saya",
    "tanggal_lahir": "1998",
    "telepon": "62891726182",
    "kewarganegaraan": "Indonesia (ID)",
    "email": "nama@gmail.com"
}
```

Response Body (Failed) 400 :
```json
{
    "message": "field cannot blank"
}
```


Response Body (Failed) 404 :
```json
{
    "message": "customer not found"
}
```



## Delete Family Member
  
- Endpoint : DELETE /api/cutomers/{customer_id}/{family_id}


Response Body (Success) 200 :

```json
{
    "message": "delete family success"
}
```


Response Body (Failed) 404 :
```json
{
    "message": "customer not found"
}
```

Response Body (Failed) 404 :
```json
{
    "message": "family not found"
}
```


Response Body (Failed) 400 :
```json
{
    "message": "delete family failed"
}
```
