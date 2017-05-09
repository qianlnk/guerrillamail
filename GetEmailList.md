GetEmailList
============

-	Description

	Gets a maximum of 20 messages from the specified offset. Offset of 0 will fetch a list of the first 10 emails, offset of 10 will fetch a list of the next 10, and so on.

	This function is useful for populating the initial email list, if you want to check for new mail use check_mail. Note: When returned, subject and email excerpt are escaped using HTML Entities.

-	Function

	get_email_list

-	Arguments

	‘sid_token’ - [required] Session ID token returned from get_email_address

	‘offset’ - [required] How many emails to start from (skip). Starts from 0

	‘seq’ - [optional] The sequence number(id) of the first email

-	Returns

	Identical to check_email described above
