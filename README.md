# All-in

To be productive in every field.

runs app at: [localhost:4040](http://localhost:4040/)
```bash
go run main.go
```
Make sure you have MongoDB installed in your local machine
## Functionality
- MongoDB
- JWT authentication
- Adding friends
- User Profile

----- not implemented yet -----

- Contribution calendar 
- Update todos and refresh daily


### Endpoints:
- Sign up: POST [localhost:4040/user/signup](http://localhost:4040/user/signup)
```json
{
    "user_name": "Jateq",
    "email": "temirlan.eraly@gmail.com",
    "password": "123123"
}
```
- Login: POST [localhost:4040/user/login](http://localhost:4040/user/login)

```json
{   
    "email": "temirlan.eraly@gmail.com",
    "password": "123123"
}
```



- Create Vault: POST [localhost:4040/user/addvault?id=](http://localhost:4040/user/addvault?id=)



### Vaults are LINKED to user by id!!!
```json 
{
    "vault_name" : "ALL-IN",
    "description" : "To be productive in every field",
    "period_days" : 20,
    "focus_mode" : false
}
```
### -------------`your id= is userID string that can be found in MongoDB `-------------
When i will(hopefully) work on frontend it will be automated. But for now:
user_id
"64554ba17ea3445a702f170b" It is gonna be unique always, so you need to copy and paste this field.
- List of Vaults: GET [localhost:4040/user/vaults?id=64554ba17ea3445a702f170b](http://localhost:4040/user/vaults?id=64554ba17ea3445a702f170b)
```json
[
    {
        "VaultID": "64554c6db4355335a107a9e9",
        "vault_name": "ace-gpa ",
        "description": "To be productive in every field",
        "created_at": "2023-05-05T18:35:25Z",
        "period_days": 20,
        "status_overall": false,
        "focus_mode": false
    },
    {
        "VaultID": "64554c8bb4355335a107a9ea",
        "vault_name": "monk-mode",
        "description": "To be productive in every field",
        "created_at": "2023-05-05T18:35:55Z",
        "period_days": 30,
        "status_overall": false,
        "focus_mode": false
    }
]
```
- Add Friends: POST [localhost:4040/user/addfriend?id=](http://localhost:4040/user/addfriend?id=)
```json
{
    "friend_id" : "64554ba17ea3445a702f170b"
}
```
It will be stored in bridge type of table. Still an issue with unique tables.

- List of Friends: GET [localhost:4040/user/friends?id=](http://localhost:4040/user/friends?id=)

note that id query in query `?id=` is id of user, so you can find friends list by this id 

-------------------------- that is how your output will look like --------------------------
```json
[
    {
        "_id": "6456294bf5c329dd498f2bc4",
        "user_name": "damir",
        "email": "damir@gmail.com",
        "password": "$2a$12$eDApAw5GN6mHiWddsQqUh.GkbB/qGku6OIy6JTbxTlAjvRXfFR8lu",
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6ImRhbWlyQGdtYWlsLmNvbSIsIlVzZXIiOiJkYW1pciIsIlVpZCI6IjY0NTYyOTRiZjVjMzI5ZGQ0OThmMmJjNCIsImV4cCI6MTY4MzQ1NDY2N30.os2J3MKs8PpM9_BaZjIcP8KH0iRjxgIKIXUc77vVGlo",
        "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6IiIsIlVzZXIiOiIiLCJVaWQiOiIiLCJleHAiOjE2ODM0NTQ2Njd9.EJajWf1FCWB6nAjoGYWAX5PH-2FBvTKk1tuac737OVw",
        "user_id": "6456294bf5c329dd498f2bc4",
        "vaults": []
    },
    {
        "_id": "645622c34bc0ef0152bf6700",
        "user_name": "Asylniet",
        "email": "asylniet@gmail.com",
        "password": "$2a$12$Pnd1.8/T6EEXqRJgoIwyAOMSj./cCroX6dLvpVCeea4vCJVzp7Xhq",
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6ImFzeWxuaWV0QGdtYWlsLmNvbSIsIlVzZXIiOiJBc3lsbmlldCIsIlVpZCI6IjY0NTYyMmMzNGJjMGVmMDE1MmJmNjcwMCIsImV4cCI6MTY4MzQ1Mjk5NX0.4DN2ydS8DbVTtiSojQ9ZpajpAkM2xtcgKRwNiERcwKU",
        "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6IiIsIlVzZXIiOiIiLCJVaWQiOiIiLCJleHAiOjE2ODM0NTI5OTV9.Sj5b6ldwtSrF1eQvXzCgZBvT8r05KUCM_QqsfyPVbKg",
        "user_id": "645622c34bc0ef0152bf6700",
        "vaults": []
    }
]
```
- Create To Do plan: POST: [localhost:4040/user/vault/all-in?id=64554ba17ea3445a702f170b](http://localhost:4040/user/vault/all-in?id=64554ba17ea3445a702f170b)
```json
[
  {
    "to_do_name": "Wake up at 7 am"
  },
  {
    "to_do_name": "Run"
  },
  {
    "to_do_name": "Prepare to nfac"
  }
]

```

- Profile info: POST [localhost:4040/user/profile?id=64554ba17ea3445a702f170b](http://localhost:4040/user/profile?id=64554ba17ea3445a702f170b)
```json
{
    "_id": "64554ba17ea3445a702f170b",
    "user_name": "Jateq",
    "email": "temirlan.eraly@gmail.com",
    "password": "",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6InRlbWlybGFuLmVyYWx5QGdtYWlsLmNvbSIsIlVzZXIiOiJKYXRlcSIsIlVpZCI6IjY0NTU0YmExN2VhMzQ0NWE3MDJmMTcwYiIsImV4cCI6MTY4MzM5NzkyMX0.Kqsq_8A8p7DTcn_uiUd2ZD7dz_b3Phvc8RfIyGGB0fI",
    "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6IiIsIlVzZXIiOiIiLCJVaWQiOiIiLCJleHAiOjE2ODMzOTc5MjF9.qKmVW30mI37asdBOZyk3udKyPUzh9pZRcAoiJnObdfU",
    "user_id": "64554ba17ea3445a702f170b",
    "vaults": [
        {
            "VaultID": "64554c6db4355335a107a9e9",
            "vault_name": "ace gpa ",
            "description": "To be productive in every field",
            "created_at": "2023-05-05T18:35:25Z",
            "period_days": 20,
            "status_overall": false,
            "focus_mode": false,
            "day": ""
        },
        {
            "VaultID": "64554c8bb4355335a107a9ea",
            "vault_name": "monk mode",
            "description": "To be productive in every field",
            "created_at": "2023-05-05T18:35:55Z",
            "period_days": 30,
            "status_overall": false,
            "focus_mode": false,
            "day": ""
        }
    ]
}
```

# `What is "All-in"`
Lets' say you need focus and try hard on one field to get best result and don't get overwhelmed.
"All-in" helps you to keep track of your goal and share your focus with others.

You have profile section with your Vaults, in essence to get max gpa on these semester when only one month is left, 
so you will create Vault, for period of 30 days
that will have 

- Wake up at 7:00 am
- Study this subject for 2 hours
- Get practice with this one
- 2 hours or less on phone
- journaling
- etc...

Every day for your period of time (30 days) your todos will refresh.
You can set focus mode - punishments for missed days. 
Most important thing that i want to implement is friends and contribution calendar.

### Project is still in process! 
Any help would be appreciated 

So I'm left with:
- refresh func for todo plan
- return vault day plan by commitID, "day" in json
- return user info
- store commit infos for calendar
- cookies?
