# encoding: utf-8
# Code generated by Microsoft (R) AutoRest Code Generator.
# Changes may cause incorrect behavior and will be lost if the code is
# regenerated.

module Azure::Security::Mgmt::V2020_08_06_preview
  #
  # API spec for Microsoft.Security (Azure Security Center) resource provider
  #
  class IotAlerts
    include MsRestAzure

    #
    # Creates and initializes a new instance of the IotAlerts class.
    # @param client service class for accessing basic functionality.
    #
    def initialize(client)
      @client = client
    end

    # @return [SecurityCenter] reference to the SecurityCenter
    attr_reader :client

    #
    # List IoT alerts
    #
    # @param scope [String] Scope of the query: Subscription (i.e.
    # /subscriptions/{subscriptionId}) or IoT Hub (i.e.
    # /subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Devices/iotHubs/{iotHubName})
    # @param min_start_time_utc [String] Filter by minimum startTimeUtc (ISO 8601
    # format)
    # @param max_start_time_utc [String] Filter by maximum startTimeUtc (ISO 8601
    # format)
    # @param alert_type [String] Filter by alert type
    # @param device_management_type [ManagementState] Get devices only from
    # specific type, Managed or Unmanaged. Possible values include: 'Managed',
    # 'Unmanaged'
    # @param compromised_entity [String] Filter by compromised device
    # @param limit [Integer] Limit the number of items returned in a single page
    # @param skip_token [String] Skip token used for pagination
    # @param custom_headers [Hash{String => String}] A hash of custom headers that
    # will be added to the HTTP request.
    #
    # @return [Array<IotAlertModel>] operation results.
    #
    def list(scope, min_start_time_utc:nil, max_start_time_utc:nil, alert_type:nil, device_management_type:nil, compromised_entity:nil, limit:nil, skip_token:nil, custom_headers:nil)
      first_page = list_as_lazy(scope, min_start_time_utc:min_start_time_utc, max_start_time_utc:max_start_time_utc, alert_type:alert_type, device_management_type:device_management_type, compromised_entity:compromised_entity, limit:limit, skip_token:skip_token, custom_headers:custom_headers)
      first_page.get_all_items
    end

    #
    # List IoT alerts
    #
    # @param scope [String] Scope of the query: Subscription (i.e.
    # /subscriptions/{subscriptionId}) or IoT Hub (i.e.
    # /subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Devices/iotHubs/{iotHubName})
    # @param min_start_time_utc [String] Filter by minimum startTimeUtc (ISO 8601
    # format)
    # @param max_start_time_utc [String] Filter by maximum startTimeUtc (ISO 8601
    # format)
    # @param alert_type [String] Filter by alert type
    # @param device_management_type [ManagementState] Get devices only from
    # specific type, Managed or Unmanaged. Possible values include: 'Managed',
    # 'Unmanaged'
    # @param compromised_entity [String] Filter by compromised device
    # @param limit [Integer] Limit the number of items returned in a single page
    # @param skip_token [String] Skip token used for pagination
    # @param custom_headers [Hash{String => String}] A hash of custom headers that
    # will be added to the HTTP request.
    #
    # @return [MsRestAzure::AzureOperationResponse] HTTP response information.
    #
    def list_with_http_info(scope, min_start_time_utc:nil, max_start_time_utc:nil, alert_type:nil, device_management_type:nil, compromised_entity:nil, limit:nil, skip_token:nil, custom_headers:nil)
      list_async(scope, min_start_time_utc:min_start_time_utc, max_start_time_utc:max_start_time_utc, alert_type:alert_type, device_management_type:device_management_type, compromised_entity:compromised_entity, limit:limit, skip_token:skip_token, custom_headers:custom_headers).value!
    end

    #
    # List IoT alerts
    #
    # @param scope [String] Scope of the query: Subscription (i.e.
    # /subscriptions/{subscriptionId}) or IoT Hub (i.e.
    # /subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Devices/iotHubs/{iotHubName})
    # @param min_start_time_utc [String] Filter by minimum startTimeUtc (ISO 8601
    # format)
    # @param max_start_time_utc [String] Filter by maximum startTimeUtc (ISO 8601
    # format)
    # @param alert_type [String] Filter by alert type
    # @param device_management_type [ManagementState] Get devices only from
    # specific type, Managed or Unmanaged. Possible values include: 'Managed',
    # 'Unmanaged'
    # @param compromised_entity [String] Filter by compromised device
    # @param limit [Integer] Limit the number of items returned in a single page
    # @param skip_token [String] Skip token used for pagination
    # @param [Hash{String => String}] A hash of custom headers that will be added
    # to the HTTP request.
    #
    # @return [Concurrent::Promise] Promise object which holds the HTTP response.
    #
    def list_async(scope, min_start_time_utc:nil, max_start_time_utc:nil, alert_type:nil, device_management_type:nil, compromised_entity:nil, limit:nil, skip_token:nil, custom_headers:nil)
      fail ArgumentError, '@client.api_version is nil' if @client.api_version.nil?
      fail ArgumentError, 'scope is nil' if scope.nil?


      request_headers = {}
      request_headers['Content-Type'] = 'application/json; charset=utf-8'

      # Set Headers
      request_headers['x-ms-client-request-id'] = SecureRandom.uuid
      request_headers['accept-language'] = @client.accept_language unless @client.accept_language.nil?
      path_template = '{scope}/providers/Microsoft.Security/iotAlerts'

      request_url = @base_url || @client.base_url

      options = {
          middlewares: [[MsRest::RetryPolicyMiddleware, times: 3, retry: 0.02], [:cookie_jar]],
          skip_encoding_path_params: {'scope' => scope},
          query_params: {'api-version' => @client.api_version,'startTimeUtc>' => min_start_time_utc,'startTimeUtc<' => max_start_time_utc,'alertType' => alert_type,'deviceManagementType' => device_management_type,'compromisedEntity' => compromised_entity,'$limit' => limit,'$skipToken' => skip_token},
          headers: request_headers.merge(custom_headers || {}),
          base_url: request_url
      }
      promise = @client.make_request_async(:get, path_template, options)

      promise = promise.then do |result|
        http_response = result.response
        status_code = http_response.status
        response_content = http_response.body
        unless status_code == 200
          error_model = JSON.load(response_content)
          fail MsRestAzure::AzureOperationError.new(result.request, http_response, error_model)
        end

        result.request_id = http_response['x-ms-request-id'] unless http_response['x-ms-request-id'].nil?
        result.correlation_request_id = http_response['x-ms-correlation-request-id'] unless http_response['x-ms-correlation-request-id'].nil?
        result.client_request_id = http_response['x-ms-client-request-id'] unless http_response['x-ms-client-request-id'].nil?
        # Deserialize Response
        if status_code == 200
          begin
            parsed_response = response_content.to_s.empty? ? nil : JSON.load(response_content)
            result_mapper = Azure::Security::Mgmt::V2020_08_06_preview::Models::IotAlertListModel.mapper()
            result.body = @client.deserialize(result_mapper, parsed_response)
          rescue Exception => e
            fail MsRest::DeserializationError.new('Error occurred in deserializing the response', e.message, e.backtrace, result)
          end
        end

        result
      end

      promise.execute
    end

    #
    # Get IoT alert
    #
    # @param scope [String] Scope of the query: Subscription (i.e.
    # /subscriptions/{subscriptionId}) or IoT Hub (i.e.
    # /subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Devices/iotHubs/{iotHubName})
    # @param iot_alert_id [String] Id of the alert
    # @param custom_headers [Hash{String => String}] A hash of custom headers that
    # will be added to the HTTP request.
    #
    # @return [IotAlertModel] operation results.
    #
    def get(scope, iot_alert_id, custom_headers:nil)
      response = get_async(scope, iot_alert_id, custom_headers:custom_headers).value!
      response.body unless response.nil?
    end

    #
    # Get IoT alert
    #
    # @param scope [String] Scope of the query: Subscription (i.e.
    # /subscriptions/{subscriptionId}) or IoT Hub (i.e.
    # /subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Devices/iotHubs/{iotHubName})
    # @param iot_alert_id [String] Id of the alert
    # @param custom_headers [Hash{String => String}] A hash of custom headers that
    # will be added to the HTTP request.
    #
    # @return [MsRestAzure::AzureOperationResponse] HTTP response information.
    #
    def get_with_http_info(scope, iot_alert_id, custom_headers:nil)
      get_async(scope, iot_alert_id, custom_headers:custom_headers).value!
    end

    #
    # Get IoT alert
    #
    # @param scope [String] Scope of the query: Subscription (i.e.
    # /subscriptions/{subscriptionId}) or IoT Hub (i.e.
    # /subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Devices/iotHubs/{iotHubName})
    # @param iot_alert_id [String] Id of the alert
    # @param [Hash{String => String}] A hash of custom headers that will be added
    # to the HTTP request.
    #
    # @return [Concurrent::Promise] Promise object which holds the HTTP response.
    #
    def get_async(scope, iot_alert_id, custom_headers:nil)
      fail ArgumentError, '@client.api_version is nil' if @client.api_version.nil?
      fail ArgumentError, 'scope is nil' if scope.nil?
      fail ArgumentError, 'iot_alert_id is nil' if iot_alert_id.nil?


      request_headers = {}
      request_headers['Content-Type'] = 'application/json; charset=utf-8'

      # Set Headers
      request_headers['x-ms-client-request-id'] = SecureRandom.uuid
      request_headers['accept-language'] = @client.accept_language unless @client.accept_language.nil?
      path_template = '{scope}/providers/Microsoft.Security/iotAlerts/{iotAlertId}'

      request_url = @base_url || @client.base_url

      options = {
          middlewares: [[MsRest::RetryPolicyMiddleware, times: 3, retry: 0.02], [:cookie_jar]],
          path_params: {'iotAlertId' => iot_alert_id},
          skip_encoding_path_params: {'scope' => scope},
          query_params: {'api-version' => @client.api_version},
          headers: request_headers.merge(custom_headers || {}),
          base_url: request_url
      }
      promise = @client.make_request_async(:get, path_template, options)

      promise = promise.then do |result|
        http_response = result.response
        status_code = http_response.status
        response_content = http_response.body
        unless status_code == 200
          error_model = JSON.load(response_content)
          fail MsRestAzure::AzureOperationError.new(result.request, http_response, error_model)
        end

        result.request_id = http_response['x-ms-request-id'] unless http_response['x-ms-request-id'].nil?
        result.correlation_request_id = http_response['x-ms-correlation-request-id'] unless http_response['x-ms-correlation-request-id'].nil?
        result.client_request_id = http_response['x-ms-client-request-id'] unless http_response['x-ms-client-request-id'].nil?
        # Deserialize Response
        if status_code == 200
          begin
            parsed_response = response_content.to_s.empty? ? nil : JSON.load(response_content)
            result_mapper = Azure::Security::Mgmt::V2020_08_06_preview::Models::IotAlertModel.mapper()
            result.body = @client.deserialize(result_mapper, parsed_response)
          rescue Exception => e
            fail MsRest::DeserializationError.new('Error occurred in deserializing the response', e.message, e.backtrace, result)
          end
        end

        result
      end

      promise.execute
    end

    #
    # List IoT alerts
    #
    # @param next_page_link [String] The NextLink from the previous successful call
    # to List operation.
    # @param custom_headers [Hash{String => String}] A hash of custom headers that
    # will be added to the HTTP request.
    #
    # @return [IotAlertListModel] operation results.
    #
    def list_next(next_page_link, custom_headers:nil)
      response = list_next_async(next_page_link, custom_headers:custom_headers).value!
      response.body unless response.nil?
    end

    #
    # List IoT alerts
    #
    # @param next_page_link [String] The NextLink from the previous successful call
    # to List operation.
    # @param custom_headers [Hash{String => String}] A hash of custom headers that
    # will be added to the HTTP request.
    #
    # @return [MsRestAzure::AzureOperationResponse] HTTP response information.
    #
    def list_next_with_http_info(next_page_link, custom_headers:nil)
      list_next_async(next_page_link, custom_headers:custom_headers).value!
    end

    #
    # List IoT alerts
    #
    # @param next_page_link [String] The NextLink from the previous successful call
    # to List operation.
    # @param [Hash{String => String}] A hash of custom headers that will be added
    # to the HTTP request.
    #
    # @return [Concurrent::Promise] Promise object which holds the HTTP response.
    #
    def list_next_async(next_page_link, custom_headers:nil)
      fail ArgumentError, 'next_page_link is nil' if next_page_link.nil?


      request_headers = {}
      request_headers['Content-Type'] = 'application/json; charset=utf-8'

      # Set Headers
      request_headers['x-ms-client-request-id'] = SecureRandom.uuid
      request_headers['accept-language'] = @client.accept_language unless @client.accept_language.nil?
      path_template = '{nextLink}'

      request_url = @base_url || @client.base_url

      options = {
          middlewares: [[MsRest::RetryPolicyMiddleware, times: 3, retry: 0.02], [:cookie_jar]],
          skip_encoding_path_params: {'nextLink' => next_page_link},
          headers: request_headers.merge(custom_headers || {}),
          base_url: request_url
      }
      promise = @client.make_request_async(:get, path_template, options)

      promise = promise.then do |result|
        http_response = result.response
        status_code = http_response.status
        response_content = http_response.body
        unless status_code == 200
          error_model = JSON.load(response_content)
          fail MsRestAzure::AzureOperationError.new(result.request, http_response, error_model)
        end

        result.request_id = http_response['x-ms-request-id'] unless http_response['x-ms-request-id'].nil?
        result.correlation_request_id = http_response['x-ms-correlation-request-id'] unless http_response['x-ms-correlation-request-id'].nil?
        result.client_request_id = http_response['x-ms-client-request-id'] unless http_response['x-ms-client-request-id'].nil?
        # Deserialize Response
        if status_code == 200
          begin
            parsed_response = response_content.to_s.empty? ? nil : JSON.load(response_content)
            result_mapper = Azure::Security::Mgmt::V2020_08_06_preview::Models::IotAlertListModel.mapper()
            result.body = @client.deserialize(result_mapper, parsed_response)
          rescue Exception => e
            fail MsRest::DeserializationError.new('Error occurred in deserializing the response', e.message, e.backtrace, result)
          end
        end

        result
      end

      promise.execute
    end

    #
    # List IoT alerts
    #
    # @param scope [String] Scope of the query: Subscription (i.e.
    # /subscriptions/{subscriptionId}) or IoT Hub (i.e.
    # /subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Devices/iotHubs/{iotHubName})
    # @param min_start_time_utc [String] Filter by minimum startTimeUtc (ISO 8601
    # format)
    # @param max_start_time_utc [String] Filter by maximum startTimeUtc (ISO 8601
    # format)
    # @param alert_type [String] Filter by alert type
    # @param device_management_type [ManagementState] Get devices only from
    # specific type, Managed or Unmanaged. Possible values include: 'Managed',
    # 'Unmanaged'
    # @param compromised_entity [String] Filter by compromised device
    # @param limit [Integer] Limit the number of items returned in a single page
    # @param skip_token [String] Skip token used for pagination
    # @param custom_headers [Hash{String => String}] A hash of custom headers that
    # will be added to the HTTP request.
    #
    # @return [IotAlertListModel] which provide lazy access to pages of the
    # response.
    #
    def list_as_lazy(scope, min_start_time_utc:nil, max_start_time_utc:nil, alert_type:nil, device_management_type:nil, compromised_entity:nil, limit:nil, skip_token:nil, custom_headers:nil)
      response = list_async(scope, min_start_time_utc:min_start_time_utc, max_start_time_utc:max_start_time_utc, alert_type:alert_type, device_management_type:device_management_type, compromised_entity:compromised_entity, limit:limit, skip_token:skip_token, custom_headers:custom_headers).value!
      unless response.nil?
        page = response.body
        page.next_method = Proc.new do |next_page_link|
          list_next_async(next_page_link, custom_headers:custom_headers)
        end
        page
      end
    end

  end
end
