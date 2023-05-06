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

----- not implemented yet -----
- User Profile 
- Contribution calendar 


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
### -----------------`your id= is userID string that can be found in MongoDB `-----------------
When i will(hopefully) work on frontend it will be automated. But for now:
user_id
"64554ba17ea3445a702f170b" It is gonna be unique always, so youn need to copy and paste this field.
- List of Vaults: GET [localhost:4040/user/vaults?id=64554ba17ea3445a702f170b](http://localhost:4040/user/vaults?id=64554ba17ea3445a702f170b)
```json
[
    {
        "VaultID": "64554c6db4355335a107a9e9",
        "vault_name": "ace gpa ",
        "description": "To be productive in every field",
        "created_at": "2023-05-05T18:35:25Z",
        "period_days": 20,
        "status_overall": false,
        "focus_mode": false
    },
    {
        "VaultID": "64554c8bb4355335a107a9ea",
        "vault_name": "monk mode",
        "description": "To be productive in every field",
        "created_at": "2023-05-05T18:35:55Z",
        "period_days": 30,
        "status_overall": false,
        "focus_mode": false
    }
]
```
- Add Friends POST: [localhost:4040/user/addfriend](http://localhost:4040/user/addfriend?id=)
```json
{
    "friend_id" : "64554ba17ea3445a702f170b"
}
```
It will be stored in bridge type of table. Still an issue with unique tables.
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
