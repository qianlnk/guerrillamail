guerrillamail
=============

Avoid spam and stay safe - use a disposable email address!

Free domain list
----------------

-	[X] guerrillamail.com
-	[ ] guerrillamailblock.com
-	[ ] sharklasers.com
-	[ ] guerrillamail.net
-	[ ] guerrillamail.org
-	[ ] guerrillamail.biz
-	[ ] spam4.me
-	[ ] grr.la
-	[ ] guerrillamail.de
-	[ ] pokemail.net

Custom domain
-------------

To add your domain name to Guerrilla Mai, go to [https://grr.la/ryo/guerrillamail.com](https://grr.la/ryo/guerrillamail.com/) and create yourself an account.

GuerrillamailClient
-------------------

```go
import (
    "github.com/qianlnk/guerrillamail"
)

# TODO

    mailClient := guerrillamail.NewGuerrillamailClient(http.DefaultClient)

# TODO
```

API list
--------

-	[x] SetEmailAddress

-	Description

	The function is used to initialize a session and set the client with an email address. If the session already exists, then it will return the email address details of the existing session. If a new session needs to be created, then it will create new email address randomly. The function will return a number of arguments for the client to remember, including the ‘sid_token’. The ‘sid_token’ is passed to each subsequent API call to maintain state.

- Arguments

    ‘lang’ - a string representing the language code. Currently supported: en, fr, nl, ru, tr, uk, ar, ko, jp, zh, zh-hant, de, es, it, pt


    ‘sid_token’ [optional] - The session id token used to maintain state. The API will also check a cookie named PHPSESSID for this token if it was not passed as an argument.


    ‘site’ [optional] - If you have your own domain and would like to access your custom domains, use this site identifier for your site. Defaults to guerrillamail.com (See RYO -> Setup -> API page for details) 

- Return

    An object with the following properties:

    ‘email_addr’ - The email address that that was determined. If a previous session was found, then it will be the email address of that session. A new random email address will be created.

    ‘email_timestamp’ - a UNIX timestamp when the email address was created. Used by the client to keep track of expiry.

    ‘sid_token’ - Session ID token. You would need to supply this parameter to every subsequent API call

-	[ ] GetEmailAddress

-	[ ] CheckEmail

-	[ ] GetEmailList

-	[ ] FetchEmail

-	[ ] ForgetMe

-	[ ] DelEmail

-	[x] GetOlderList
