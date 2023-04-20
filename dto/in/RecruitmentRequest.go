package in

type RecruitmentRequest struct {
	Page        int    `json:"page"`
	Description string `json:"description"`
	Location    string `json:"location"`
	FullTime    bool   `json:"full_time"`
	FullTimeSet bool   `json:"full_time_set"`
}
