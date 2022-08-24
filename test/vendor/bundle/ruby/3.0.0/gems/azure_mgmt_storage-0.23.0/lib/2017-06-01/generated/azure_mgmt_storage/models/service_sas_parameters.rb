# encoding: utf-8
# Code generated by Microsoft (R) AutoRest Code Generator.
# Changes may cause incorrect behavior and will be lost if the code is
# regenerated.

module Azure::Storage::Mgmt::V2017_06_01
  module Models
    #
    # The parameters to list service SAS credentials of a specific resource.
    #
    class ServiceSasParameters

      include MsRestAzure

      # @return [String] The canonical path to the signed resource.
      attr_accessor :canonicalized_resource

      # @return [SignedResource] The signed services accessible with the
      # service SAS. Possible values include: Blob (b), Container (c), File
      # (f), Share (s). Possible values include: 'b', 'c', 'f', 's'
      attr_accessor :resource

      # @return [Permissions] The signed permissions for the service SAS.
      # Possible values include: Read (r), Write (w), Delete (d), List (l), Add
      # (a), Create (c), Update (u) and Process (p). Possible values include:
      # 'r', 'd', 'w', 'l', 'a', 'c', 'u', 'p'
      attr_accessor :permissions

      # @return [String] An IP address or a range of IP addresses from which to
      # accept requests.
      attr_accessor :ipaddress_or_range

      # @return [HttpProtocol] The protocol permitted for a request made with
      # the account SAS. Possible values include: 'https,http', 'https'
      attr_accessor :protocols

      # @return [DateTime] The time at which the SAS becomes valid.
      attr_accessor :shared_access_start_time

      # @return [DateTime] The time at which the shared access signature
      # becomes invalid.
      attr_accessor :shared_access_expiry_time

      # @return [String] A unique value up to 64 characters in length that
      # correlates to an access policy specified for the container, queue, or
      # table.
      attr_accessor :identifier

      # @return [String] The start of partition key.
      attr_accessor :partition_key_start

      # @return [String] The end of partition key.
      attr_accessor :partition_key_end

      # @return [String] The start of row key.
      attr_accessor :row_key_start

      # @return [String] The end of row key.
      attr_accessor :row_key_end

      # @return [String] The key to sign the account SAS token with.
      attr_accessor :key_to_sign

      # @return [String] The response header override for cache control.
      attr_accessor :cache_control

      # @return [String] The response header override for content disposition.
      attr_accessor :content_disposition

      # @return [String] The response header override for content encoding.
      attr_accessor :content_encoding

      # @return [String] The response header override for content language.
      attr_accessor :content_language

      # @return [String] The response header override for content type.
      attr_accessor :content_type


      #
      # Mapper for ServiceSasParameters class as Ruby Hash.
      # This will be used for serialization/deserialization.
      #
      def self.mapper()
        {
          client_side_validation: true,
          required: false,
          serialized_name: 'ServiceSasParameters',
          type: {
            name: 'Composite',
            class_name: 'ServiceSasParameters',
            model_properties: {
              canonicalized_resource: {
                client_side_validation: true,
                required: true,
                serialized_name: 'canonicalizedResource',
                type: {
                  name: 'String'
                }
              },
              resource: {
                client_side_validation: true,
                required: true,
                serialized_name: 'signedResource',
                type: {
                  name: 'String'
                }
              },
              permissions: {
                client_side_validation: true,
                required: false,
                serialized_name: 'signedPermission',
                type: {
                  name: 'String'
                }
              },
              ipaddress_or_range: {
                client_side_validation: true,
                required: false,
                serialized_name: 'signedIp',
                type: {
                  name: 'String'
                }
              },
              protocols: {
                client_side_validation: true,
                required: false,
                serialized_name: 'signedProtocol',
                type: {
                  name: 'Enum',
                  module: 'HttpProtocol'
                }
              },
              shared_access_start_time: {
                client_side_validation: true,
                required: false,
                serialized_name: 'signedStart',
                type: {
                  name: 'DateTime'
                }
              },
              shared_access_expiry_time: {
                client_side_validation: true,
                required: false,
                serialized_name: 'signedExpiry',
                type: {
                  name: 'DateTime'
                }
              },
              identifier: {
                client_side_validation: true,
                required: false,
                serialized_name: 'signedIdentifier',
                constraints: {
                  MaxLength: 64
                },
                type: {
                  name: 'String'
                }
              },
              partition_key_start: {
                client_side_validation: true,
                required: false,
                serialized_name: 'startPk',
                type: {
                  name: 'String'
                }
              },
              partition_key_end: {
                client_side_validation: true,
                required: false,
                serialized_name: 'endPk',
                type: {
                  name: 'String'
                }
              },
              row_key_start: {
                client_side_validation: true,
                required: false,
                serialized_name: 'startRk',
                type: {
                  name: 'String'
                }
              },
              row_key_end: {
                client_side_validation: true,
                required: false,
                serialized_name: 'endRk',
                type: {
                  name: 'String'
                }
              },
              key_to_sign: {
                client_side_validation: true,
                required: false,
                serialized_name: 'keyToSign',
                type: {
                  name: 'String'
                }
              },
              cache_control: {
                client_side_validation: true,
                required: false,
                serialized_name: 'rscc',
                type: {
                  name: 'String'
                }
              },
              content_disposition: {
                client_side_validation: true,
                required: false,
                serialized_name: 'rscd',
                type: {
                  name: 'String'
                }
              },
              content_encoding: {
                client_side_validation: true,
                required: false,
                serialized_name: 'rsce',
                type: {
                  name: 'String'
                }
              },
              content_language: {
                client_side_validation: true,
                required: false,
                serialized_name: 'rscl',
                type: {
                  name: 'String'
                }
              },
              content_type: {
                client_side_validation: true,
                required: false,
                serialized_name: 'rsct',
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
