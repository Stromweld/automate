// Code generated by protoc-gen-policy. DO NOT EDIT.
// source: api/external/infra_proxy/infra_proxy.proto

package infra_proxy

import (
	request "github.com/chef/automate/api/external/infra_proxy/request"
	policy "github.com/chef/automate/components/automate-gateway/api/iam/v2/policy"
)

func init() {
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/GetVersion", "system:service:version", "system:serviceVersion:get", "GET", "/api/v0/infra/version", func(unexpandedResource string, input interface{}) string {
		return unexpandedResource
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/GetServers", "infra:infraServers", "infra:infraServers:list", "GET", "/api/v0/infra/servers", func(unexpandedResource string, input interface{}) string {
		return unexpandedResource
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/GetServer", "infra:infraServers:{id}", "infra:infraServers:get", "GET", "/api/v0/infra/servers/{id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.GetServer); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/GetServerByName", "infra:infraServers:{name}", "infra:infraServers:get", "GET", "/api/v0/infra/servers/{name}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.GetServerByName); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "name":
					return m.Name
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/CreateServer", "infra:infraServers", "infra:infraServers:create", "POST", "/api/v0/infra/servers", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.CreateServer); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "name":
					return m.Name
				case "description":
					return m.Description
				case "fqdn":
					return m.Fqdn
				case "ip_address":
					return m.IpAddress
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/UpdateServer", "infra:infraServers:{id}", "infra:infraServers:update", "PUT", "/api/v0/infra/servers/{id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.UpdateServer); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				case "name":
					return m.Name
				case "description":
					return m.Description
				case "fqdn":
					return m.Fqdn
				case "ip_address":
					return m.IpAddress
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/DeleteServer", "infra:infraServers:{id}", "infra:infraServers:delete", "DELETE", "/api/v0/infra/servers/{id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.DeleteServer); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/GetOrgs", "infra:infraServers:{server_id}:orgs", "infra:infraServers:get", "GET", "/api/v0/infra/servers/{server_id}/orgs", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.GetOrgs); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "server_id":
					return m.ServerId
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/GetOrg", "infra:infraServers:{server_id}:orgs:{id}", "infra:infraServers:get", "GET", "/api/v0/infra/servers/{server_id}/orgs/{id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.GetOrg); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				case "server_id":
					return m.ServerId
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/GetOrgByName", "infra:infraServers:{server_id}:orgs:{name}", "infra:infraServers:get", "GET", "/api/v0/infra/servers/{server_id}/orgs/{name}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.GetOrgByName); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "name":
					return m.Name
				case "server_id":
					return m.ServerId
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/CreateOrg", "infra:infraServers:{server_id}:orgs", "infra:infraServers:update", "POST", "/api/v0/infra/servers/{server_id}/orgs", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.CreateOrg); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "name":
					return m.Name
				case "admin_user":
					return m.AdminUser
				case "admin_key":
					return m.AdminKey
				case "server_id":
					return m.ServerId
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/UpdateOrg", "infra:infraServers:{server_id}:orgs:{id}", "infra:infraServers:update", "PUT", "/api/v0/infra/servers/{server_id}/orgs/{id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.UpdateOrg); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				case "name":
					return m.Name
				case "admin_user":
					return m.AdminUser
				case "admin_key":
					return m.AdminKey
				case "server_id":
					return m.ServerId
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/DeleteOrg", "infra:infraServers:{server_id}:orgs:{id}", "infra:infraServers:update", "DELETE", "/api/v0/infra/servers/{server_id}/orgs/{id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.DeleteOrg); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				case "server_id":
					return m.ServerId
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/GetCookbooks", "infra:infraServers:{server_id}:orgs:{org_id}:cookbooks", "infra:infraServers:get", "GET", "/api/v0/infra/servers/{server_id}/orgs/{org_id}/cookbooks", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.Cookbooks); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "org_id":
					return m.OrgId
				case "server_id":
					return m.ServerId
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/GetCookbookVersions", "infra:infraServers:{server_id}:orgs:{org_id}:cookbooks", "infra:infraServers:get", "GET", "/api/v0/infra/servers/{server_id}/orgs/{org_id}/cookbooks/{name}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.CookbookVersions); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "org_id":
					return m.OrgId
				case "server_id":
					return m.ServerId
				case "name":
					return m.Name
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/GetCookbook", "infra:infraServers:{server_id}:orgs:{org_id}:cookbooks", "infra:infraServers:get", "GET", "/api/v0/infra/servers/{server_id}/orgs/{org_id}/cookbooks/{name}/{version}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.Cookbook); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "org_id":
					return m.OrgId
				case "server_id":
					return m.ServerId
				case "name":
					return m.Name
				case "version":
					return m.Version
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/GetCookbookFileContent", "infra:infraServers:{server_id}:orgs:{org_id}:cookbooks", "infra:infraServers:get", "GET", "/api/v0/infra/servers/{server_id}/orgs/{org_id}/cookbooks/{name}/{version}/file-content", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.CookbookFileContent); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "org_id":
					return m.OrgId
				case "server_id":
					return m.ServerId
				case "name":
					return m.Name
				case "version":
					return m.Version
				case "url":
					return m.Url
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/GetRoles", "infra:infraServers:{server_id}:orgs:{org_id}:roles", "infra:infraServers:get", "GET", "/api/v0/infra/servers/{server_id}/orgs/{org_id}/roles", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.Roles); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "org_id":
					return m.OrgId
				case "server_id":
					return m.ServerId
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/GetRole", "infra:infraServers:{server_id}:orgs:{org_id}:roles", "infra:infraServers:get", "GET", "/api/v0/infra/servers/{server_id}/orgs/{org_id}/roles/{name}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.Role); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "org_id":
					return m.OrgId
				case "server_id":
					return m.ServerId
				case "name":
					return m.Name
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/GetClients", "infra:infraServers:{server_id}:orgs:{org_id}:clients", "infra:infraServers:get", "GET", "/api/v0/infra/servers/{server_id}/orgs/{org_id}/clients", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.Clients); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "org_id":
					return m.OrgId
				case "server_id":
					return m.ServerId
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/GetClient", "infra:servers:{server_id}:orgs:{org_id}:clients", "infra:infraServers:get", "GET", "/api/v0/infra/servers/{server_id}/orgs/{org_id}/clients/{name}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.Client); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "org_id":
					return m.OrgId
				case "server_id":
					return m.ServerId
				case "name":
					return m.Name
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/GetDataBags", "infra:infraServers:{server_id}:orgs:{org_id}:data_bags", "infra:infraServers:get", "GET", "/api/v0/infra/servers/{server_id}/orgs/{org_id}/data_bags", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.DataBags); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "org_id":
					return m.OrgId
				case "server_id":
					return m.ServerId
				case "name":
					return m.Name
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/GetDataBagItem", "infra:infraServers:{server_id}:orgs:{org_id}:data_bags", "infra:infraServers:get", "GET", "/api/v0/infra/servers/{server_id}/orgs/{org_id}/data_bags/{name}/{item}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.DataBag); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "org_id":
					return m.OrgId
				case "server_id":
					return m.ServerId
				case "name":
					return m.Name
				case "item":
					return m.Item
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/GetEnvironments", "infra:infraServers:{server_id}:orgs:{org_id}:environments", "infra:infraServers:get", "GET", "/api/v0/infra/servers/{server_id}/orgs/{org_id}/environments", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.Environments); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "org_id":
					return m.OrgId
				case "server_id":
					return m.ServerId
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/GetEnvironment", "infra:infraServers:{server_id}:orgs:{org_id}:environments", "infra:infraServers:get", "GET", "/api/v0/infra/servers/{server_id}/orgs/{org_id}/environments/{name}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.Environment); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "org_id":
					return m.OrgId
				case "server_id":
					return m.ServerId
				case "name":
					return m.Name
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/GetAffectedNodes", "infra:infraServers:{server_id}:orgs:{org_id}:nodes", "infra:infraServers:get", "GET", "/api/v0/infra/servers/{server_id}/orgs/{org_id}/affected-nodes/{chef_type}/{name}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.AffectedNodes); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "org_id":
					return m.OrgId
				case "server_id":
					return m.ServerId
				case "chef_type":
					return m.ChefType
				case "name":
					return m.Name
				case "version":
					return m.Version
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/GetPolicyfiles", "infra:infraServers:{server_id}:orgs:{org_id}:policyfiles", "infra:infraServers:get", "GET", "/api/v0/infra/servers/{server_id}/orgs/{org_id}/policyfiles", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.Policyfiles); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "org_id":
					return m.OrgId
				case "server_id":
					return m.ServerId
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/GetPolicyfile", "infra:infraServers:{server_id}:orgs:{org_id}:policyfiles", "infra:infraServers:get", "GET", "/api/v0/infra/servers/{server_id}/orgs/{org_id}/policyfiles/{name}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.Policyfile); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "org_id":
					return m.OrgId
				case "server_id":
					return m.ServerId
				case "name":
					return m.Name
				case "revision_id":
					return m.RevisionId
				default:
					return ""
				}
			})
		}
		return ""
	})
}
