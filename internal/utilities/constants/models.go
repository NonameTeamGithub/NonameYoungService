package constants

type ResponseMessagesStruct struct {
	AccessDenied        string
	EmailAlreadyInUse   string
	InternalServerError string
	InvalidData         string
	InvalidToken        string
	MissingData         string
	MissingToken        string
	NotFound            string
	Ok                  string
	TooManyRequests     string
}

type RolesStruct struct {
	Admin     string
	Intern    string
	Candidate string
	Curator   string
	Mentor    string
	Hr        string
}
