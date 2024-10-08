# E-mail Normalizer

Normalizing e-mail addresses in Go.

## Steps

1. remove leading and trailing whitespaces
2. convert to lowercase
3. remove dots from email's username
4. remove the part after '+' symbol of the email's username

## Why normalize emails ?

Because some people use the same email with aliases to abuse your services' free tier.

In Gmail, the email `abanoub@gmail.com` is the same email as `abanoub+1@gmail.com` and it is the same email as `abanou.b@gmail.com`.

I like to store the user written email in the database alongside the normalized email, and check the free tier validation against the normalized emails column in the database table.

| email                | normalized_email  |
|:---------------------|:------------------|
| `abanoub+4@gmail.com`|`abanoub@gmail.com`|
| `a.shraf@gmail.com`  |`ashraf@gmail.com` |
