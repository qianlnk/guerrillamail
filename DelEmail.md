DelEmail
========

-	Function del_email

-	Description

	Delete the emails from the server

-	Arguments

	‘sid_token’ - Session ID token returned from get_email_address

	email_ids - array or an integer, in the following format: email_ids[]=425&email_ids[]=426&email_ids[]=427

	Where 425, 426 and 427 are the ids of emails to delete. In a HTTP request, this parameter would be encoded like this: email_ids%5B%5D=425&email_ids%5B%5D=426&email_ids%5B%5D=427

-	Returns

	An array of deleted email ids An object with the following property: deleted_ids - an array of deleted email ids
