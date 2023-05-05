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

----- not implemented yet -----
- User Profile 
- Contribution calendar 
- Friends field

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
## `your id= is userID string that can be found in MongoDB `
When i will(hopefully) work on frontend it will be automated 
- List of Vaults: GET [localhost:4040/user/vaults](http://localhost:4040/user/vaults)

# What is "All-in"
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
