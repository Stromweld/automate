# encoding: utf-8
# Code generated by Microsoft (R) AutoRest Code Generator.
# Changes may cause incorrect behavior and will be lost if the code is
# regenerated.

module Azure::KeyVault::Mgmt::V2019_09_01
  module Models
    #
    # Parameters for updating the access policy in a vault
    #
    class VaultAccessPolicyParameters

      include MsRestAzure

      # @return [String] The resource id of the access policy.
      attr_accessor :id

      # @return [String] The resource name of the access policy.
      attr_accessor :name

      # @return [String] The resource name of the access policy.
      attr_accessor :type

      # @return [String] The resource type of the access policy.
      attr_accessor :location

      # @return [VaultAccessPolicyProperties] Properties of the access policy
      attr_accessor :properties


      #
      # Mapper for VaultAccessPolicyParameters class as Ruby Hash.
      # This will be used for serialization/deserialization.
      #
      def self.mapper()
        {
          client_side_validation: true,
          required: false,
          serialized_name: 'VaultAccessPolicyParameters',
          type: {
            name: 'Composite',
            class_name: 'VaultAccessPolicyParameters',
            model_properties: {
              id: {
                client_side_validation: true,
                required: false,
                read_only: true,
                serialized_name: 'id',
                type: {
                  name: 'String'
                }
              },
              name: {
                client_side_validation: true,
                required: false,
                read_only: true,
                serialized_name: 'name',
                type: {
                  name: 'String'
                }
              },
              type: {
                client_side_validation: true,
                required: false,
                read_only: true,
                serialized_name: 'type',
                type: {
                  name: 'String'
                }
              },
              location: {
                client_side_validation: true,
                required: false,
                read_only: true,
                serialized_name: 'location',
                type: {
                  name: 'String'
                }
              },
              properties: {
                client_side_validation: true,
                required: true,
                serialized_name: 'properties',
                type: {
                  name: 'Composite',
                  class_name: 'VaultAccessPolicyProperties'
                }
              }
            }
          }
        }
      end
    end
  end
end
