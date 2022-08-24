# encoding: utf-8
# Code generated by Microsoft (R) AutoRest Code Generator.
# Changes may cause incorrect behavior and will be lost if the code is
# regenerated.

module Azure::Security::Mgmt::V2020_01_01
  module Models
    #
    # Model object.
    #
    #
    class SecuritySolutionsReferenceData

      include MsRestAzure

      # @return [String] Resource Id
      attr_accessor :id

      # @return [String] Resource name
      attr_accessor :name

      # @return [String] Resource type
      attr_accessor :type

      # @return [String] Location where the resource is stored
      attr_accessor :location

      # @return [SecurityFamily] The security family of the security solution.
      # Possible values include: 'Waf', 'Ngfw', 'SaasWaf', 'Va'
      attr_accessor :security_family

      # @return [String] The security solutions' vendor name
      attr_accessor :alert_vendor_name

      # @return [String] The security solutions' package info url
      attr_accessor :package_info_url

      # @return [String] The security solutions' product name
      attr_accessor :product_name

      # @return [String] The security solutions' publisher
      attr_accessor :publisher

      # @return [String] The security solutions' publisher display name
      attr_accessor :publisher_display_name

      # @return [String] The security solutions' template
      attr_accessor :template


      #
      # Mapper for SecuritySolutionsReferenceData class as Ruby Hash.
      # This will be used for serialization/deserialization.
      #
      def self.mapper()
        {
          client_side_validation: true,
          required: false,
          serialized_name: 'securitySolutionsReferenceData',
          type: {
            name: 'Composite',
            class_name: 'SecuritySolutionsReferenceData',
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
              security_family: {
                client_side_validation: true,
                required: true,
                serialized_name: 'properties.securityFamily',
                type: {
                  name: 'String'
                }
              },
              alert_vendor_name: {
                client_side_validation: true,
                required: true,
                serialized_name: 'properties.alertVendorName',
                type: {
                  name: 'String'
                }
              },
              package_info_url: {
                client_side_validation: true,
                required: true,
                serialized_name: 'properties.packageInfoUrl',
                type: {
                  name: 'String'
                }
              },
              product_name: {
                client_side_validation: true,
                required: true,
                serialized_name: 'properties.productName',
                type: {
                  name: 'String'
                }
              },
              publisher: {
                client_side_validation: true,
                required: true,
                serialized_name: 'properties.publisher',
                type: {
                  name: 'String'
                }
              },
              publisher_display_name: {
                client_side_validation: true,
                required: true,
                serialized_name: 'properties.publisherDisplayName',
                type: {
                  name: 'String'
                }
              },
              template: {
                client_side_validation: true,
                required: true,
                serialized_name: 'properties.template',
                type: {
                  name: 'String'
                }
              }
            }
          }
        }
      end
    end
  end
end
