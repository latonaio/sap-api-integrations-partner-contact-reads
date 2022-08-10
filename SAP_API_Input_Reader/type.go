package sap_api_input_reader

type EC_MC struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Document      struct {
		DocumentNo     string `json:"document_no"`
		DeliverTo      string `json:"deliver_to"`
		Quantity       string `json:"quantity"`
		PickedQuantity string `json:"picked_quantity"`
		Price          string `json:"price"`
		Batch          string `json:"batch"`
	} `json:"document"`
	ProductionOrder struct {
		DocumentNo           string `json:"document_no"`
		Status               string `json:"status"`
		DeliverTo            string `json:"deliver_to"`
		Quantity             string `json:"quantity"`
		CompletedQuantity    string `json:"completed_quantity"`
		PlannedStartDate     string `json:"planned_start_date"`
		PlannedValidatedDate string `json:"planned_validated_date"`
		ActualStartDate      string `json:"actual_start_date"`
		ActualValidatedDate  string `json:"actual_validated_date"`
		Batch                string `json:"batch"`
		Work                 struct {
			WorkNo                   string `json:"work_no"`
			Quantity                 string `json:"quantity"`
			CompletedQuantity        string `json:"completed_quantity"`
			ErroredQuantity          string `json:"errored_quantity"`
			Component                string `json:"component"`
			PlannedComponentQuantity string `json:"planned_component_quantity"`
			PlannedStartDate         string `json:"planned_start_date"`
			PlannedStartTime         string `json:"planned_start_time"`
			PlannedValidatedDate     string `json:"planned_validated_date"`
			PlannedValidatedTime     string `json:"planned_validated_time"`
			ActualStartDate          string `json:"actual_start_date"`
			ActualStartTime          string `json:"actual_start_time"`
			ActualValidatedDate      string `json:"actual_validated_date"`
			ActualValidatedTime      string `json:"actual_validated_time"`
		} `json:"work"`
	} `json:"production_order"`
	APISchema      string `json:"api_schema"`
	MaterialCode   string `json:"material_code"`
	Plant_Supplier string `json:"plant/supplier"`
	Stock          string `json:"stock"`
	DocumentType   string `json:"document_type"`
	DocumentNo     string `json:"document_no"`
	PlannedDate    string `json:"planned_date"`
	ValidatedDate  string `json:"validated_date"`
	Deleted        bool   `json:"deleted"`
}

type SDC struct {
	ConnectionKey      string `json:"connection_key"`
	Result             bool   `json:"result"`
	RedisKey           string `json:"redis_key"`
	Filepath           string `json:"filepath"`
	PartnerContactCollection struct {
	    ObjectID                          string `json:"ObjectID"`
	    PartnerContactID                  string `json:"PartnerContactID"`
	    PartnerContactUUID                string `json:"PartnerContactUUID"`
	    UserID                            string `json:"UserID"`
	    IdentityUUID                      string `json:"IdentityUUID"`
	    CreateUser                        bool   `json:"CreateUser"`
	    BusinessPartnerID                 string `json:"BusinessPartnerID"`
	    PartnerID                         string `json:"PartnerID"`
	    StatusCode                        string `json:"StatusCode"`
	    StatusCodeText                    string `json:"StatusCodeText"`
	    TitleCode                         string `json:"TitleCode"`
	    TitleCodeText                     string `json:"TitleCodeText"`
	    AcademicTitleCode                 string `json:"AcademicTitleCode"`
	    AcademicTitleCodeText             string `json:"AcademicTitleCodeText"`
	    BusinessPartnerFormattedName      string `json:"BusinessPartnerFormattedName"`
	    FirstName                         string `json:"FirstName"`
	    LastName                          string `json:"LastName"`
	    MiddleName                        string `json:"MiddleName"`
	    AdditionalLastName                string `json:"AdditionalLastName"`
	    GenderCode                        string `json:"GenderCode"`
	    GenderCodeText                    string `json:"GenderCodeText"`
	    MaritalStatusCode                 string `json:"MaritalStatusCode"`
	    MaritalStatusCodeText             string `json:"MaritalStatusCodeText"`
	    LanguageCode                      string `json:"LanguageCode"`
	    LanguageCodeText                  string `json:"LanguageCodeText"`
	    BirthDate                         string `json:"BirthDate"`
	    VIPContactCode                    string `json:"VIPContactCode"`
	    VIPContactCodeText                string `json:"VIPContactCodeText"`
	    Department                        string `json:"Department"`
	    JobTitle                          string `json:"JobTitle"`
	    FormattedPostalAddressDescription string `json:"FormattedPostalAddressDescription"`
	    CountryCode                       string `json:"CountryCode"`
	    CountryCodeText                   string `json:"CountryCodeText"`
	    StateCode                         string `json:"StateCode"`
	    StateCodeText                     string `json:"StateCodeText"`
	    AddressLine1                      string `json:"AddressLine1"`
	    AddressLine2                      string `json:"AddressLine2"`
	    HouseNumber                       string `json:"HouseNumber"`
	    Street                            string `json:"Street"`
	    AddressLine4                      string `json:"AddressLine4"`
	    AddressLine5                      string `json:"AddressLine5"`
	    City                              string `json:"City"`
	    StreetPostalCode                  string `json:"StreetPostalCode"`
	    Phone                             string `json:"Phone"`
	    Mobile                            string `json:"Mobile"`
	    Fax                               string `json:"Fax"`
	    Email                             string `json:"Email"`
	    BestReachedByCode                 string `json:"BestReachedByCode"`
	    BestReachedByCodeText             string `json:"BestReachedByCodeText"`
	    NormalisedPhone                   string `json:"NormalisedPhone"`
	    NormalisedMobile                  string `json:"NormalisedMobile"`
	    CreatedOn                         string `json:"CreatedOn"`
	    CreatedBy                         string `json:"CreatedBy"`
	    CreatedByIdentityUUID             string `json:"CreatedByIdentityUUID"`
	    ChangedOn                         string `json:"ChangedOn"`
	    ChangedBy                         string `json:"ChangedBy"`
	    ChangedByIdentityUUID             string `json:"ChangedByIdentityUUID"`
	    EntityLastChangedOn               string `json:"EntityLastChangedOn"`
	    ETag                              string `json:"ETag"`
	} `json:"PartnerContactCollection"`
	APISchema    string   `json:"api_schema"`
	Accepter     []string `json:"accepter"`
	PartnerContactCode string   `json:"partnercontact_code"`
	Deleted      bool     `json:"deleted"`
}
