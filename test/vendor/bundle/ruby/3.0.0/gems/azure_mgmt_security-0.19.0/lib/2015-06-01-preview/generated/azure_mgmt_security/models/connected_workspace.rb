# encoding: utf-8
# Code generated by Microsoft (R) AutoRest Code Generator.
# Changes may cause incorrect behavior and will be lost if the code is
# regenerated.

module Azure::Security::Mgmt::V2015_06_01_preview
  module Models
    #
    # Represents an OMS workspace to which the solution is connected
    #
    #
    class ConnectedWorkspace

      include MsRestAzure

      # @return [String] Azure resource ID of the connected OMS workspace
      attr_accessor :id


      #
      # Mapper for ConnectedWorkspace class as Ruby Hash.
      # This will be used for serialization/deserialization.
      #
      def self.mapper()
        {
          client_side_validation: true,
          required: false,
          serialized_name: 'ConnectedWorkspace',
          type: {
            name: 'Composite',
            class_name: 'ConnectedWorkspace',
            model_properties: {
              id: {
                client_side_validation: true,
                required: false,
                serialized_name: 'id',
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
