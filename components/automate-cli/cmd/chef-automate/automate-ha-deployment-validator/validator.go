package automatehadeploymentvalidator

import (
	"fmt"

	automatenodevalidator "github.com/chef/automate/components/automate-cli/cmd/chef-automate/automate-ha-deployment-validator/automate-node-validator"
	bastionnodevalidator "github.com/chef/automate/components/automate-cli/cmd/chef-automate/automate-ha-deployment-validator/bastion-node-validator"
	csnodevalidator "github.com/chef/automate/components/automate-cli/cmd/chef-automate/automate-ha-deployment-validator/cs-node-validator"
	osnodevalidator "github.com/chef/automate/components/automate-cli/cmd/chef-automate/automate-ha-deployment-validator/os-node-validator"
	pgnodevalidator "github.com/chef/automate/components/automate-cli/cmd/chef-automate/automate-ha-deployment-validator/pg-node-validator"
)

// ValidatorFlags struct definition,
// The Validation will be performed only for the nodes for which the corresponding flag is true
type ValidatorFlags struct {
	All        bool
	Bastion    bool
	Automate   bool
	Chefserver bool
	Postgresql bool
	Opensearch bool
}

// HADeploymentValidator interface definition
type HADeploymentValidator interface {
	RunHAValidator() ([]string, error)
}

// NewHADeploymentValidator returns the new HA deployment validator
func NewHADeploymentValidator(flags ValidatorFlags, configPath string) (HADeploymentValidator, error) {

	bValidator, err := bastionnodevalidator.NewBastionNodeValidator()
	if err != nil {
		return nil, fmt.Errorf("Unable to initialize bastion validator")
	}

	a2Validator, err := automatenodevalidator.NewAutomateNodeValidator()
	if err != nil {
		return nil, fmt.Errorf("Unable to initialize automate validator")
	}

	csValidator, err := csnodevalidator.NewCSNodeValidator()
	if err != nil {
		return nil, fmt.Errorf("Unable to initialize chef server validator")
	}

	pgValidator, err := pgnodevalidator.NewPGNodeValidator()
	if err != nil {
		return nil, fmt.Errorf("Unable to initialize postgres validator")
	}

	osValidator, err := osnodevalidator.NewOSNodeValidator()
	if err != nil {
		return nil, fmt.Errorf("Unable to initialize opensearch validator")
	}

	return &HAValidator{
		Flags:             flags,
		ConfigPath:        configPath,
		BastionValidator:  bValidator,
		AutomateValidator: a2Validator,
		CSValidator:       csValidator,
		PGValidator:       pgValidator,
		OSValidator:       osValidator,
	}, nil
}

// HAValidator struct definition
type HAValidator struct {
	Flags             ValidatorFlags
	ConfigPath        string
	AutomateValidator automatenodevalidator.AutomateNodeValidator
	BastionValidator  bastionnodevalidator.BastionNodeValidator
	CSValidator       csnodevalidator.CSNodeValidator
	PGValidator       pgnodevalidator.PGNodeValidator
	OSValidator       osnodevalidator.OSNodeValidator
}

// RunHAValidator run the infra validation
func (validator *HAValidator) RunHAValidator() ([]string, error) {
	fmt.Println("Flags:", validator.Flags)
	fmt.Println("ConfigPath:", validator.ConfigPath)
	fmt.Println("Running HA Validator")
	fmt.Println("valid config file")

	var errResp error
	validationsResp := []string{}
	if validator.Flags.Bastion || validator.Flags.All {
		resp, err := validator.BastionValidator.Run()
		if err != nil {
			errResp = fmt.Errorf("%s, %s", errResp.Error(), err.Error())
		}
		if len(resp) > 0 {
			validationsResp = append(validationsResp, resp...)
		}
	}

	if validator.Flags.Automate || validator.Flags.All {
		resp, err := validator.AutomateValidator.Run()
		if err != nil {
			errResp = fmt.Errorf("%s, %s", errResp.Error(), err.Error())
		}
		if len(resp) > 0 {
			validationsResp = append(validationsResp, resp...)
		}
	}

	if validator.Flags.Chefserver || validator.Flags.All {
		resp, err := validator.CSValidator.Run()
		if err != nil {
			errResp = fmt.Errorf("%s, %s", errResp.Error(), err.Error())
		}
		if len(resp) > 0 {
			validationsResp = append(validationsResp, resp...)
		}
	}

	if validator.Flags.Postgresql || validator.Flags.All {
		resp, err := validator.PGValidator.Run()
		if err != nil {
			errResp = fmt.Errorf("%s, %s", errResp.Error(), err.Error())
		}
		if len(resp) > 0 {
			validationsResp = append(validationsResp, resp...)
		}
	}

	if validator.Flags.Opensearch || validator.Flags.All {
		resp, err := validator.OSValidator.Run()
		if err != nil {
			errResp = fmt.Errorf("%s, %s", errResp.Error(), err.Error())
		}
		if len(resp) > 0 {
			validationsResp = append(validationsResp, resp...)
		}
	}
	return validationsResp, errResp
}
