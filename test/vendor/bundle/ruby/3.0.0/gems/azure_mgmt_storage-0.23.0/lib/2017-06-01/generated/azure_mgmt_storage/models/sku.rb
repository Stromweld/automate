# encoding: utf-8
# Code generated by Microsoft (R) AutoRest Code Generator.
# Changes may cause incorrect behavior and will be lost if the code is
# regenerated.

module Azure::Storage::Mgmt::V2017_06_01
  module Models
    #
    # The SKU of the storage account.
    #
    class Sku

      include MsRestAzure

      # @return [SkuName] Gets or sets the sku name. Required for account
      # creation; optional for update. Note that in older versions, sku name
      # was called accountType. Possible values include: 'Standard_LRS',
      # 'Standard_GRS', 'Standard_RAGRS', 'Standard_ZRS', 'Premium_LRS'
      attr_accessor :name

      # @return [SkuTier] Gets the sku tier. This is based on the SKU name.
      # Possible values include: 'Standard', 'Premium'
      attr_accessor :tier

      # @return [String] The type of the resource, usually it is
      # 'storageAccounts'.
      attr_accessor :resource_type

      # @return [Kind] Indicates the type of storage account. Possible values
      # include: 'Storage', 'BlobStorage'
      attr_accessor :kind

      # @return [Array<String>] The set of locations that the SKU is available.
      # This will be supported and registered Azure Geo Regions (e.g. West US,
      # East US, Southeast Asia, etc.).
      attr_accessor :locations

      # @return [Array<SKUCapability>] The capability information in the
      # specified sku, including file encryption, network acls, change
      # notification, etc.
      attr_accessor :capabilities

      # @return [Array<Restriction>] The restrictions because of which SKU
      # cannot be used. This is empty if there are no restrictions.
      attr_accessor :restrictions


      #
      # Mapper for Sku class as Ruby Hash.
      # This will be used for serialization/deserialization.
      #
      def self.mapper()
        {
          client_side_validation: true,
          required: false,
          serialized_name: 'Sku',
          type: {
            name: 'Composite',
            class_name: 'Sku',
            model_properties: {
              name: {
                client_side_validation: true,
                required: true,
                serialized_name: 'name',
                type: {
                  name: 'Enum',
                  module: 'SkuName'
                }
              },
              tier: {
                client_side_validation: true,
                required: false,
                read_only: true,
                serialized_name: 'tier',
                type: {
                  name: 'Enum',
                  module: 'SkuTier'
                }
              },
              resource_type: {
                client_side_validation: true,
                required: false,
                read_only: true,
                serialized_name: 'resourceType',
                type: {
                  name: 'String'
                }
              },
              kind: {
                client_side_validation: true,
                required: false,
                read_only: true,
                serialized_name: 'kind',
                type: {
                  name: 'Enum',
                  module: 'Kind'
                }
              },
              locations: {
                client_side_validation: true,
                required: false,
                read_only: true,
                serialized_name: 'locations',
                type: {
                  name: 'Sequence',
                  element: {
                      client_side_validation: true,
                      required: false,
                      serialized_name: 'StringElementType',
                      type: {
                        name: 'String'
                      }
                  }
                }
              },
              capabilities: {
                client_side_validation: true,
                required: false,
                read_only: true,
                serialized_name: 'capabilities',
                type: {
                  name: 'Sequence',
                  element: {
                      client_side_validation: true,
                      required: false,
                      serialized_name: 'SKUCapabilityElementType',
                      type: {
                        name: 'Composite',
                        class_name: 'SKUCapability'
                      }
                  }
                }
              },
              restrictions: {
                client_side_validation: true,
                required: false,
                serialized_name: 'restrictions',
                type: {
                  name: 'Sequence',
                  element: {
                      client_side_validation: true,
                      required: false,
                      serialized_name: 'RestrictionElementType',
                      type: {
                        name: 'Composite',
                        class_name: 'Restriction'
                      }
                  }
                }
              }
            }
          }
        }
      end
    end
  end
end
