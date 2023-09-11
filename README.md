# gqlgen-graphql-tutorial

## Mutations

### register
```bash
xdg-open http://localhost:8080/
```

```graphql
mutation register {
  register(
    input: {
      username: "user1", 
      email: "user1@gmail.com", 
      password: "user1234567890!", 
      confirmPassword: "user1234567890!", 
      firstName: "User", 
      lastName: "One"}
  ) {
    authToken {
      accessToken
    }
  }
}
```

expected result:

```json
{
  "data": {
    "register": {
      "authToken": {
        "accessToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjU0NDQ3NzYsImp0aSI6IjIiLCJpYXQiOjE2NjQ4Mzk5NzYsImlzcyI6Im1lZXRtZXVwIn0.erk1xgejBMJd2k0DNQZEZ-SzN6RxO6U44Jz_75zHmzk"
      }
    }
  }
}
```

## References

* [https://github.com/EQuimper/youtube-golang-graphql-tutorial](https://github.com/EQuimper/youtube-golang-graphql-tutorial)
* [https://www.youtube.com/playlist?list=PLzQWIQOqeUSNwXcneWYJHUREAIucJ5UZn](https://www.youtube.com/playlist?list=PLzQWIQOqeUSNwXcneWYJHUREAIucJ5UZn)
