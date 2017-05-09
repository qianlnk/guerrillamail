GetOlderList
============

-	Function

	get_older_list

-	Description

	The idea of this function is to get some emails that are older (lower ID) than ‘seq’. How can this be useful? Say we have more than 1 page of emails. Then, we delete some emails from the first page, so we are left with some empty slots at the bottom. Instead of calling fetch_email_list, we call this to be more efficient and get exactly the number of email we need to fill.

-	Arguments

	‘sid_token’ - Session ID token returned from get_email_address

	‘seq’ - Integer. get emails that have a lower id than the ‘seq’

	‘limit’ - Integer how many emails to fetch max.

-	Returns

	Same structure as returned by fetch_email_list
