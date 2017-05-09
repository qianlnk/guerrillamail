package guerrillamail

type Authoraze struct {
	Success    bool          `json:"success"`
	ErrorCodes []interface{} `json:"error_codes"`
}

type Stats struct {
	SequenceMail     string `json:"sequence_mail"`
	CreatedAddresses int    `json:"created_addresses"`
	ReceivedEmails   string `json:"received_emails"`
	Total            string `json:"total"`
	TotalPerHour     string `json:"total_per_hour"`
}

type CheckEmailList struct {
	MailID        string `json:"mail_id"`
	MailFrom      string `json:"mail_from"`
	MailSubject   string `json:"mail_subject"`
	MailExcerpt   string `json:"mail_excerpt"`
	MailTimestamp string `json:"mail_timestamp"`
	MailRead      string `json:"mail_read"`
	MailDate      string `json:"mail_date"`
	Att           string `json:"att"`
	MailSize      string `json:"mail_size"`
}

type GetEmailList struct {
	MailFrom      string `json:"mail_from"`
	MailTimestamp int    `json:"mail_timestamp"`
	MailRead      int    `json:"mail_read"`
	MailDate      string `json:"mail_date"`
	ReplyTo       string `json:"reply_to"`
	MailSubject   string `json:"mail_subject"`
	MailExcerpt   string `json:"mail_excerpt"`
	MailID        int    `json:"mail_id"`
	Att           int    `json:"att"`
	ContentType   string `json:"content_type"`
	MailRecipient string `json:"mail_recipient"`
	SourceID      int    `json:"source_id"`
	SourceMailID  int    `json:"source_mail_id"`
	MailBody      string `json:"mail_body"`
	Size          int    `json:"size"`
}

type SetEmailUserResponse struct {
	AliasError     string    `json:"alias_error"`
	Alias          string    `json:"alias"`
	EmailAddr      string    `json:"email_addr"`
	EmailTimestamp int       `json:"email_timestamp"`
	SiteID         int       `json:"site_id"`
	SidToken       string    `json:"sid_token"`
	Site           string    `json:"site"`
	Auth           Authoraze `json:"auth"`
}

type GetEmailAddressResponse struct {
	EmailAddr      string `json:"email_addr"`
	EmailTimestamp int    `json:"email_timestamp"`
	Alias          string `json:"alias"`
	SidToken       string `json:"sid_token"`
}

type CheckEmailResponse struct {
	List     []CheckEmailList `json:"list"`
	Count    interface{}      `json:"count"`
	Email    string           `json:"email"`
	Alias    string           `json:"alias"`
	Ts       int              `json:"ts"`
	SidToken string           `json:"sid_token"`
	Stats    Stats            `json:"stats"`
	Auth     Authoraze        `json:"auth"`
}

type GetEmailListResponse struct {
	List     []GetEmailList `json:"list"`
	Count    string         `json:"count"`
	Email    string         `json:"email"`
	Alias    string         `json:"alias"`
	Ts       int            `json:"ts"`
	SidToken string         `json:"sid_token"`
	Stats    Stats          `json:"stats"`
	Auth     Authoraze      `json:"auth"`
}

type FetchEmailResponse struct {
	MailID        string    `json:"mail_id"`
	MailFrom      string    `json:"mail_from"`
	MailRecipient string    `json:"mail_recipient"`
	MailSubject   string    `json:"mail_subject"`
	MailExcerpt   string    `json:"mail_excerpt"`
	MailBody      string    `json:"mail_body"`
	MailTimestamp string    `json:"mail_timestamp"`
	MailDate      string    `json:"mail_date"`
	MailRead      string    `json:"mail_read"`
	ContentType   string    `json:"content_type"`
	SourceID      string    `json:"source_id"`
	SourceMailID  string    `json:"source_mail_id"`
	ReplyTo       string    `json:"reply_to"`
	MailSize      string    `json:"mail_size"`
	Ver           string    `json:"ver"`
	RefMid        string    `json:"ref_mid"`
	SidToken      string    `json:"sid_token"`
	Auth          Authoraze `json:"auth"`
}
