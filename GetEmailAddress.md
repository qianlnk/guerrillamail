GetEmailAddress
===============

-	Function

	get_email_address

-	Description

	The function is used to initialize a session and set the client with an email address. If the session already exists, then it will return the email address details of the existing session. If a new session needs to be created, then it will create new email address randomly. The function will return a number of arguments for the client to remember, including the ‘sid_token’. The ‘sid_token’ is passed to each subsequent API call to maintain state.

-	Arguments

	‘lang’ - a string representing the language code. Currently supported: en, fr, nl, ru, tr, uk, ar, ko, jp, zh, zh-hant, de, es, it, pt

	‘sid_token’ [optional] - The session id token used to maintain state. The API will also check a cookie named PHPSESSID for this token if it was not passed as an argument.

	‘site’ [optional] - If you have your own domain and would like to access your custom domains, use this site identifier for your site. Defaults to guerrillamail.com (See RYO -> Setup -> API page for details)

	Alternative method for keeping track of the session:

	The session can be also maintained using HTTP Cookies. The cookie name is ‘PHPSESSID’, and it will be given in the resulting HTTP header. The client can store this cookie and send it whenever making an API call. A new session will be created when the ‘PHPSESSID’ cookie is not given by the client. The function also generates a new welcome email for the user.

	To send ‘PHPSESSID’ in a HTTP request, add the following line to the header:

	Cookie: PHPSESSID=ABC1234\r\n";

	Where ‘ABC1234’ is the data for the ‘PHPSESSID’ cookie.

	The client should always watch the ‘PHPSESSID’ cookie for changes. If it changes, it needs to update this value, and pass it in subsequent calls.

-	Returns

	An object with the following properties:

	‘email_addr’ - The email address that that was determined. If a previous session was found, then it will be the email address of that session. A new random email address will be created.

	‘email_timestamp’ - a UNIX timestamp when the email address was created. Used by the client to keep track of expiry.

	‘sid_token’ - Session ID token. You would need to supply this parameter to every subsequent API call
