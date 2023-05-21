package campaign

type CampaignFormatter struct {
	ID               int    `json:"id"`
	UserID           int    `json:"user_id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	ImageURL         string `json:"image_url"`
	GoalAmount       int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_amount"`
}

func FormatCampaign(campaign Campaign) CampaignFormatter {
	campaignFomatter := CampaignFormatter{}
	campaignFomatter.ID = campaign.ID
	campaignFomatter.UserID = campaign.UserID
	campaignFomatter.Name = campaign.Name
	campaignFomatter.ShortDescription = campaign.ShortDescription
	campaignFomatter.GoalAmount = campaign.GoalAmount
	campaignFomatter.CurrentAmount = campaign.CurrentAmount
	campaignFomatter.ImageURL = ""

	if len(campaign.CampaignImages) > 0 {
		campaignFomatter.ImageURL = campaign.CampaignImages[0].FileName
	}

	return campaignFomatter
}

func FormatCampaigns(campaigns []Campaign) []CampaignFormatter {
	campaignsFormatter := []CampaignFormatter{}

	for _, campaign := range campaigns {
		campaignFormatter := FormatCampaign(campaign)
		campaignsFormatter = append(campaignsFormatter, campaignFormatter)
	}

	return campaignsFormatter
}
