# encoding: utf-8
# Code generated by Microsoft (R) AutoRest Code Generator.
# Changes may cause incorrect behavior and will be lost if the code is
# regenerated.

module Azure::Security::Mgmt::V2020_01_01_preview
  module Models
    #
    # Describes an Azure resource with location
    #
    class AzureTrackedResourceLocation

      include MsRestAzure

      # @return [String] Location where the resource is stored
      attr_accessor :location


      #
      # Mapper for AzureTrackedResourceLocation class as Ruby Hash.
      # This will be used for serialization/deserialization.
      #
      def self.mapper()
        {
          client_side_validation: true,
          required: false,
          serialized_name: 'AzureTrackedResourceLocation',
          type: {
            name: 'Composite',
            class_name: 'AzureTrackedResourceLocation',
            model_properties: {
              location: {
                client_side_validation: true,
                required: false,
                serialized_name: 'location',
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
