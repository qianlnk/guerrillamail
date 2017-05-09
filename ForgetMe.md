ForgetMe
========

-	Function

	forget_me

-	Description

	Forget the current email address. This will not stop the session, the existing session will be maintained. A subsequent call to get_email_address will fetch a new email address or the client can call set_email_user to set a new address. Typically, a user would want to set a new address manually after clicking the ‘forget me’ button.

-	Arguments

	‘sid_token’ - Session ID token returned from get_email_address

	‘email_addr’ - the email address to forget. The email_addr should be the same value as returned by set_email_user or get_email_address

-	Returns

	True if successful.
