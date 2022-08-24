# encoding: utf-8
# Code generated by Microsoft (R) AutoRest Code Generator.
# Changes may cause incorrect behavior and will be lost if the code is
# regenerated.

module Azure::KeyVault::Mgmt::V2016_10_01
  module Models
    #
    # Properties of the vault
    #
    class VaultProperties

      include MsRestAzure

      # @return The Azure Active Directory tenant ID that should be used for
      # authenticating requests to the key vault.
      attr_accessor :tenant_id

      # @return [Sku] SKU details
      attr_accessor :sku

      # @return [Array<AccessPolicyEntry>] An array of 0 to 16 identities that
      # have access to the key vault. All identities in the array must use the
      # same tenant ID as the key vault's tenant ID.
      attr_accessor :access_policies

      # @return [String] The URI of the vault for performing operations on keys
      # and secrets.
      attr_accessor :vault_uri

      # @return [Boolean] Property to specify whether Azure Virtual Machines
      # are permitted to retrieve certificates stored as secrets from the key
      # vault.
      attr_accessor :enabled_for_deployment

      # @return [Boolean] Property to specify whether Azure Disk Encryption is
      # permitted to retrieve secrets from the vault and unwrap keys.
      attr_accessor :enabled_for_disk_encryption

      # @return [Boolean] Property to specify whether Azure Resource Manager is
      # permitted to retrieve secrets from the key vault.
      attr_accessor :enabled_for_template_deployment

      # @return [Boolean] Property specifying whether recoverable deletion is
      # enabled for this key vault. Setting this property to true activates the
      # soft delete feature, whereby vaults or vault entities can be recovered
      # after deletion. Enabling this functionality is irreversible - that is,
      # the property does not accept false as its value.
      attr_accessor :enable_soft_delete

      # @return [CreateMode] The vault's create mode to indicate whether the
      # vault need to be recovered or not. Possible values include: 'recover',
      # 'default'
      attr_accessor :create_mode

      # @return [Boolean] Property specifying whether protection against purge
      # is enabled for this vault. Setting this property to true activates
      # protection against purge for this vault and its content - only the Key
      # Vault service may initiate a hard, irrecoverable deletion. The setting
      # is effective only if soft delete is also enabled. Enabling this
      # functionality is irreversible - that is, the property does not accept
      # false as its value.
      attr_accessor :enable_purge_protection


      #
      # Mapper for VaultProperties class as Ruby Hash.
      # This will be used for serialization/deserialization.
      #
      def self.mapper()
        {
          client_side_validation: true,
          required: false,
          serialized_name: 'VaultProperties',
          type: {
            name: 'Composite',
            class_name: 'VaultProperties',
            model_properties: {
              tenant_id: {
                client_side_validation: true,
                required: true,
                serialized_name: 'tenantId',
                type: {
                  name: 'String'
                }
              },
              sku: {
                client_side_validation: true,
                required: true,
                serialized_name: 'sku',
                default_value: {},
                type: {
                  name: 'Composite',
                  class_name: 'Sku'
                }
              },
              access_policies: {
                client_side_validation: true,
                required: false,
                serialized_name: 'accessPolicies',
                type: {
                  name: 'Sequence',
                  element: {
                      client_side_validation: true,
                      required: false,
                      serialized_name: 'AccessPolicyEntryElementType',
                      type: {
                        name: 'Composite',
                        class_name: 'AccessPolicyEntry'
                      }
                  }
                }
              },
              vault_uri: {
                client_side_validation: true,
                required: false,
                serialized_name: 'vaultUri',
                type: {
                  name: 'String'
                }
              },
              enabled_for_deployment: {
                client_side_validation: true,
                required: false,
                serialized_name: 'enabledForDeployment',
                type: {
                  name: 'Boolean'
                }
              },
              enabled_for_disk_encryption: {
                client_side_validation: true,
                required: false,
                serialized_name: 'enabledForDiskEncryption',
                type: {
                  name: 'Boolean'
                }
              },
              enabled_for_template_deployment: {
                client_side_validation: true,
                required: false,
                serialized_name: 'enabledForTemplateDeployment',
                type: {
                  name: 'Boolean'
                }
              },
              enable_soft_delete: {
                client_side_validation: true,
                required: false,
                serialized_name: 'enableSoftDelete',
                type: {
                  name: 'Boolean'
                }
              },
              create_mode: {
                client_side_validation: true,
                required: false,
                serialized_name: 'createMode',
                type: {
                  name: 'Enum',
                  module: 'CreateMode'
                }
              },
              enable_purge_protection: {
                client_side_validation: true,
                required: false,
                serialized_name: 'enablePurgeProtection',
                type: {
                  name: 'Boolean'
                }
              }
            }
          }
        }
      end
    end
  end
end
