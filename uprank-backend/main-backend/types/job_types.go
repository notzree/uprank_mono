package types

type CreateJobRequest struct {
	UpworkJobRequest *AttachUpworkJobRequest `json:"upwork_job_request,omitempty"`
}

func (req *CreateJobRequest) Validate() map[string]interface{} {
	errors := make(map[string]interface{})

	// Validate Upwork job request if it exists
	if req.UpworkJobRequest != nil {
		upworkErrors := req.UpworkJobRequest.Validate()
		if len(upworkErrors) > 0 {
			errors["UpworkJobRequest"] = upworkErrors
		}
	}
	//todo: When we add on more platforms, check and ensure that at least one platform is specified
	if req.UpworkJobRequest == nil {
		errors["UpworkJobRequest"] = "UpworkJobRequest is required"
	}
	return errors
}
