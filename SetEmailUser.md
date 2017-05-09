SetEmailUser
============

-	Function

	set_email_user

-	Description

	Set the email address to a different email address.. If the email has already been set, it will be given additional 60 minutes. Otherwise, a new email address will be generated if the email address is not in the database.

-	Arguments

	‘email_user’ - String, 74 chars max. The ‘local part’ of an email address. example: test@guerrillamailblock.com, the email_user part would be ‘test’

	‘lang’ - String. The language code, eg. ‘en’.

	‘sid_token’ - Pass this to maintain state, the value for this argument is returned after calling get_email_address

	Note: You don’t need to pass the PHPSESSID cookie if you pass the ‘sid_token’ argument

	‘site’ [optional] - If you have your own domain and would like to access your custom domains, use this site identifier for your site. Defaults to guerrillamail.com (See RYO -> Setup -> API page for details)

-	Returns

	Similar to get_email_address, this function returns an object with the following properties:

	‘email_addr’ - The email address that that was set.

	‘email_timestamp’ - a UNIX timestamp when the email address was created. Used by the client to keep track of expiry.

	‘sid_token’ - Session ID token

	‘site’ - (string) The identifier of the site matched

	‘site_id’ - (int) The ID of the site which the domain belongs to

	‘alias’ - (string) This is the scrambled version of email_addr, eg “fcs5d3+vgdtknvsyt4”. The scrambled address is an alias for email_addr, so all email going to alias will arrive at the inbox of email_addr. Scrambled addresses are used to mask the email_addr, so it is difficult for someone else to know which email_addr was used.

	‘alias_error’ (string) Error message if the alias could not be set
