# Use custom Logic Apps connector to integrate with on-premises APIs

### High level steps

1. An application with a REST endpoint.  This application is either deployed in Azure VNET or on-premises.  A simple echo API written in Golang is available in [go/src/hello](go/src/hello) folder.

2. On-premise data gateway.  Install the gateway on the same server or another server that can access the target API.  If the gateway is behind firewall, then ensure the domain names are whitelisted for outbound connectivity.  Follow documentation to download and install:  [https://docs.microsoft.com/en-us/azure/logic-apps/logic-apps-gateway-install](https://docs.microsoft.com/en-us/azure/logic-apps/logic-apps-gateway-install).

    > Connectivity information is documented at [https://docs.microsoft.com/en-us/data-integration/gateway/service-gateway-communication](https://docs.microsoft.com/en-us/data-integration/gateway/service-gateway-communication)

3. Connect on-premises gateway to Logic Apps.  Follow documentation: [https://docs.microsoft.com/en-us/azure/logic-apps/logic-apps-gateway-connection](https://docs.microsoft.com/en-us/azure/logic-apps/logic-apps-gateway-connection).

4. Create a custom Logic Apps connector.  Follow documentation: [https://docs.microsoft.com/en-us/connectors/custom-connectors/create-logic-apps-connector](https://docs.microsoft.com/en-us/connectors/custom-connectors/create-logic-apps-connector) and [https://docs.microsoft.com/en-us/connectors/custom-connectors/define-openapi-definition](https://docs.microsoft.com/en-us/connectors/custom-connectors/define-openapi-definition).  Make sure to enable the flag "Connect via on-premises data gateway".

    > If you don't have an Swagger definition, then create the API defintion manually using [https://docs.microsoft.com/en-us/connectors/custom-connectors/define-blank](https://docs.microsoft.com/en-us/connectors/custom-connectors/define-blank)

5. Create a Logic App in the same region as the custom Logic Apps connector.  The custom connector will now be available under the **custom** tab for use.  On the first use, you will be prompted to select the on-premises data gateway.

# Integrate Logic Apps with API Management for advanced security (i.e. built-in JWT validation), routing and message handling requirements

1. Import Logic App to API Management using [https://docs.microsoft.com/en-us/azure/api-management/import-logic-app-as-api](https://docs.microsoft.com/en-us/azure/api-management/import-logic-app-as-api).

2. Add **validate-jwt** policy to API under **Inbound Processing**.  The example policy will check whether the JWT token is valid (i.e. not expired, a jwt token, not missing), but additional restrictions such as claims and audiences can be added per [https://docs.microsoft.com/en-us/azure/api-management/api-management-access-restriction-policies#ValidateJWT](https://docs.microsoft.com/en-us/azure/api-management/api-management-access-restriction-policies#ValidateJWT)

```xml
        <validate-jwt header-name="Authorization" failed-validation-httpcode="401" failed-validation-error-message="Unauthorized. Access token is missing or invalid.">
            <openid-config url="https://login.microsoftonline.com/microsoft.onmicrosoft.com/.well-known/openid-configuration" />
        </validate-jwt>
```

3. Add policy to remove the **Authorization** header from being passed downstream to Logic App.

```xml
        <set-header name="Authorization" exists-action="delete" />
```

All of the policies for the API:

```xml
<policies>
    <inbound>
        <base />
        <validate-jwt header-name="Authorization" failed-validation-httpcode="401" failed-validation-error-message="Unauthorized. Access token is missing or invalid.">
            <openid-config url="https://login.microsoftonline.com/microsoft.onmicrosoft.com/.well-known/openid-configuration" />
        </validate-jwt>
        <set-backend-service id="apim-generated-policy" backend-id="LogicApp_demo_test-logic-apps" />
        <set-header name="Authorization" exists-action="delete" />
    </inbound>
    <backend>
        <base />
    </backend>
    <outbound>
        <base />
    </outbound>
    <on-error>
        <base />
    </on-error>
</policies>
```

# Ensure API Management is the only path to Logic App

1. Identify the Public IP of API Management (available in Developer, Standard and Premium SKUs).  It is located on the **Overview** screen on Azure Portal.

2. Navigate to Logic App -> Workflow settings and add the IP address; follow [https://docs.microsoft.com/en-us/azure/logic-apps/logic-apps-securing-a-logic-app#restrict-inbound-ip-addresses](https://docs.microsoft.com/en-us/azure/logic-apps/logic-apps-securing-a-logic-app#restrict-inbound-ip-addresses).  Once enabled, access from IPs that are not whitelisted will produce the error:

```json
{
    "error": {
        "code": "AuthorizationFailed",
        "message": "The client IP address 'a.b.c.d' is not in the allowed caller IP address ranges specified in the workflow access control configuration."
    }
}
```