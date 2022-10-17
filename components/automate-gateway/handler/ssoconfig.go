package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"bytes"
	"io/ioutil"
	"strings"
    "net/http"
	"github.com/chef/automate/api/external/sso"
	deployment "github.com/chef/automate/api/interservice/deployment"
	license_control "github.com/chef/automate/api/interservice/license_control"
	"github.com/golang/protobuf/ptypes/empty"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// SsoConfig - the ssoconfig service data structure
type SsoConfig struct {
	license_client license_control.LicenseControlServiceClient
	client         deployment.DeploymentClient
}

type PostConfig struct {
	Ca_contents           string   `json:"ca_contents"`
	Sso_url               string   `json:"sso_url"`
	Email_attr            string   `json:"email_attr"`
	Username_attr         string   `json:"username_attr"`
	Groups_attr           string   `json:"groups_attr"`
	Allowed_groups        []string `json:"allowed_groups"`
	Entity_issuer         string   `json:"entity_issuer"`
	Name_id_policy_format string   `json:"name_id_policy_format"`
}
const ssoFilesPath = "/var/automate-ha/"

// NewSsoConfigHandler - create a new ssoconfig service handler
func NewSsoConfigHandler(license_client license_control.LicenseControlServiceClient, client deployment.DeploymentClient) *SsoConfig {
	return &SsoConfig{
		license_client: license_client,
		client:         client,
	}
}

func (a *SsoConfig) GetSsoConfig(ctx context.Context, in *empty.Empty) (*sso.GetSsoConfigResponse, error) {
	err := a.validateDeploymentType(ctx)
	if err != nil {
		return nil, err
	}

	res, err := a.getConfigData(ctx)
	if err !=nil {
		return nil, err
	}

	if res.Config.Dex != nil {
		ssoConfig := res.Config.Dex.V1.Sys.Connectors.Saml

		return &sso.GetSsoConfigResponse{
			CaContents:         ssoConfig.CaContents.GetValue(),
			SsoUrl:             ssoConfig.SsoUrl.GetValue(),
			EmailAttr:          ssoConfig.EmailAttr.GetValue(),
			UsernameAttr:       ssoConfig.UsernameAttr.GetValue(),
			GroupsAttr:         ssoConfig.GroupsAttr.GetValue(),
			AllowedGroups:      ssoConfig.AllowedGroups,
			EntityIssuer:       ssoConfig.EntityIssuer.GetValue(),
			NameIdPolicyFormat: ssoConfig.NameIdPolicyFormat.GetValue(),
		}, nil
	}

	return &sso.GetSsoConfigResponse{
		CaContents:         "",
		SsoUrl:             "",
		EmailAttr:          "",
		UsernameAttr:       "",
		GroupsAttr:         "",
		AllowedGroups:      []string{},
		EntityIssuer:       "",
		NameIdPolicyFormat: "",
	}, nil
}

func (a *SsoConfig) getConfigData(ctx context.Context) (*deployment.GetAutomateConfigResponse, error) {
	req := &deployment.GetAutomateConfigRequest{}

	return a.client.GetAutomateConfig(ctx, req)
}

func (a *SsoConfig) DeleteSsoConfig(ctx context.Context, in *empty.Empty) (*sso.DeleteSsoConfigResponse, error) {
	err := a.validateDeploymentType(ctx)
	if err != nil {
		return nil, err
	}

	log.Printf("Received request to delete sso config")
	res, err := a.getConfigData(ctx)

	if err !=nil {
		return nil, err
	}
	
	if res.Config.Dex != nil {
		url, err := getBastionUrl()
		if err != nil {
			return nil, err
		}
		fileName := "revert-status.txt"
		err = ioutil.WriteFile(ssoFilesPath+fileName, []byte("Pending"), 0777)
		if err != nil {
			log.Printf("Unable to write file:", err)
			return nil, err
		}
		go makeRequest("DELETE", *url, nil, fileName)
		return &sso.DeleteSsoConfigResponse{
            Message: "Started Disabling SSO Configuration",
        }, nil
    }

	return &sso.DeleteSsoConfigResponse{
		Message: "SSO Configuration not disabled successfully",
	}, nil
}

func(a *SsoConfig) validateDeploymentType(ctx context.Context) error {
	deploymentType, err := a.getDeploymentDetails(ctx)
	if err != nil {
		return err
	}

	if deploymentType != "SAAS" {
		msg := "Unauthorized: Deployment type is not SAAS"
		return status.Error(codes.PermissionDenied, msg)
	}
	return nil
}

func (a *SsoConfig) getDeploymentDetails(ctx context.Context) (string, error) {
	deployIDResponse, err := a.license_client.GetDeploymentID(ctx, &license_control.GetDeploymentIDRequest{})
	if err != nil {
		if isServiceDownError(err) {
			log.WithFields(log.Fields{
				"err":  err,
				"func": "getDeploymentDetails",
			}).Error("connecting to the license client")
			return "", nil
		}
		return "", err
	}

	log.Debugf("deployIDResponse.DeploymentType: %s ", deployIDResponse.DeploymentType)

	return deployIDResponse.DeploymentType, nil
}


func makeRequest(requestType string, url string, jsonData []byte, fileName string) {
	req, err := http.NewRequest(requestType, url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal("Error occurred", err)
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Fatal("Error occurred", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		ioutil.WriteFile(ssoFilesPath+fileName, []byte("Success"), 0777)
		return
	}
	ioutil.WriteFile(ssoFilesPath+fileName, []byte("Failure"), 0777)
}

func (a *SsoConfig) SetSsoConfig(ctx context.Context, in *sso.SetSsoConfigRequest) (*sso.SetSsoConfigResponse, error) {
	err := a.validateDeploymentType(ctx)
	if err != nil {
		return nil, err
	}

	req := &sso.SetSsoConfigRequest{
		CaContents:         in.CaContents,
		SsoUrl:             in.SsoUrl,
		EmailAttr:          in.EmailAttr,
		UsernameAttr:       in.UsernameAttr,
		GroupsAttr:         in.GroupsAttr,
		AllowedGroups:      in.AllowedGroups,
		EntityIssuer:       in.EntityIssuer,
		NameIdPolicyFormat: in.NameIdPolicyFormat,
	}
	bodyParams := &PostConfig{
		Ca_contents:           req.CaContents,
		Sso_url:               req.SsoUrl,
		Email_attr:            req.EmailAttr,
		Username_attr:         req.UsernameAttr,
		Groups_attr:           req.GroupsAttr,
		Allowed_groups:        req.AllowedGroups,
		Entity_issuer:         req.EntityIssuer,
		Name_id_policy_format: req.NameIdPolicyFormat,
	}
	jsonValue, _ := json.Marshal(bodyParams)
	url, err := getBastionUrl()
	if err != nil {
		log.Fatalln("Errror while creating bastion url = ", err)
		return nil, err
	}
	fileName := "post-status.txt"
	err = ioutil.WriteFile(ssoFilesPath+fileName, []byte("Pending"), 0777)
	if err != nil {
		fmt.Printf("Unable to write the file: %v", err)
	}
	go makeRequest("POST", *url, jsonValue, fileName)
	return &sso.SetSsoConfigResponse{
		Response: "Config patch was successfull",
	}, nil
}

func getBastionUrl() (*string, error) {
	content, err := ioutil.ReadFile(ssoFilesPath+"bastion_info.txt")
	if err != nil {
		log.Fatal("Error occurred while reading file: ", err)
		return nil, err
	}
	url := "http://" + strings.TrimSpace(string(content))
	return &url, nil
}
