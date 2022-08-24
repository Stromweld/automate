# encoding: utf-8
# Code generated by Microsoft (R) AutoRest Code Generator.
# Changes may cause incorrect behavior and will be lost if the code is
# regenerated.

module Azure::Security::Mgmt::V2020_08_06_preview
  module Models
    #
    # Contains all OVF (virtual machine) full versions for the sensor
    #
    class PackageDownloadsSensorFullOvf

      include MsRestAzure

      # @return [Array<PackageDownloadInfo>] Enterprise package type
      attr_accessor :enterprise

      # @return [Array<PackageDownloadInfo>] Medium package type
      attr_accessor :medium

      # @return [Array<PackageDownloadInfo>] Line package type
      attr_accessor :line


      #
      # Mapper for PackageDownloadsSensorFullOvf class as Ruby Hash.
      # This will be used for serialization/deserialization.
      #
      def self.mapper()
        {
          client_side_validation: true,
          required: false,
          serialized_name: 'PackageDownloads_sensor_full_ovf',
          type: {
            name: 'Composite',
            class_name: 'PackageDownloadsSensorFullOvf',
            model_properties: {
              enterprise: {
                client_side_validation: true,
                required: false,
                read_only: true,
                serialized_name: 'enterprise',
                type: {
                  name: 'Sequence',
                  element: {
                      client_side_validation: true,
                      required: false,
                      serialized_name: 'PackageDownloadInfoElementType',
                      type: {
                        name: 'Composite',
                        class_name: 'PackageDownloadInfo'
                      }
                  }
                }
              },
              medium: {
                client_side_validation: true,
                required: false,
                read_only: true,
                serialized_name: 'medium',
                type: {
                  name: 'Sequence',
                  element: {
                      client_side_validation: true,
                      required: false,
                      serialized_name: 'PackageDownloadInfoElementType',
                      type: {
                        name: 'Composite',
                        class_name: 'PackageDownloadInfo'
                      }
                  }
                }
              },
              line: {
                client_side_validation: true,
                required: false,
                read_only: true,
                serialized_name: 'line',
                type: {
                  name: 'Sequence',
                  element: {
                      client_side_validation: true,
                      required: false,
                      serialized_name: 'PackageDownloadInfoElementType',
                      type: {
                        name: 'Composite',
                        class_name: 'PackageDownloadInfo'
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
