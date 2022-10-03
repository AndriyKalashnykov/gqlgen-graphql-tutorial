# gqlgen-graphql-tutorial

## Mutations

### register

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

## References

* [https://github.com/EQuimper/youtube-golang-graphql-tutorial](https://github.com/EQuimper/youtube-golang-graphql-tutorial)
* [https://www.youtube.com/playlist?list=PLzQWIQOqeUSNwXcneWYJHUREAIucJ5UZn](https://www.youtube.com/playlist?list=PLzQWIQOqeUSNwXcneWYJHUREAIucJ5UZn)
