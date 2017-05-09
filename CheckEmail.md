CheckEmail
==========

-	Function

	check_email

-	Description

	Check for new email on the server. Returns a list of the newest messages. The maximum size of the list is 20 items.

	The client should not check email too many times as to not overload the server. Do not check if the email expired, the email checking routing should pause if the email expired.

-	Arguments

	‘sid_token’ - Session ID token returned from get_email_address

	seq - The sequence number (id) of the oldest email

-	Returns

	An an object with the following properties:

	‘list’ - A list of messages, represented as an array of objects. Each object has the following properties:

	```
	‘mail_id’, 
	```

	‘mail_from’ (email address of sender), ‘mail_subject’, ‘mail_excerpt’ (snippet from the email), ‘mail_timestamp’ (a UNIX timestamp), ‘mail_read’ (1 if read, 0 if not),

	```
	‘mail_date’
	```

	Note: ‘mail_subject’ and ‘mail_excerpt’ are escaped using HTML Entities.

	‘count’ - The number of emails matched. This can be much more than the maximum that can be returned (20)! This number indicates the total number of new emails in the database. (If you want to get the emails after the first 20, then make another call with the get_older_list function

	‘email’ - The email address which the list is for. Eg. test@guerrillamailblock.com The client would need to watch this variable to detect any changes to sync the address. ‘ts’ - The timestamp of the email address, time of when the email address was created. The client can use this variable to sync the time. Emails expire after 60 minutes ‘sid_token’ - Session ID token ‘users’ - the number of users viewing this inbox (not exact)
