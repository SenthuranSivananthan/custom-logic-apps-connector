# Use custom Logic Apps connector to integrate with on-premises APIs

### High level steps

1. An application with a REST.  This application is either deployed in Azure VNET or on-premises.  A simple echo API written in Golang is available in [go/src/hello](go/src/hello) folder.

2. On-premise data gateway.  Install the gateway on the same server or another server that can access the target API.  If the gateway is behind firewall, then ensure the domain names are whitelisted for outbound connectivity.  Follow documentation to download and install:  [https://docs.microsoft.com/en-us/azure/logic-apps/logic-apps-gateway-install](https://docs.microsoft.com/en-us/azure/logic-apps/logic-apps-gateway-install).

    > Connectivity information is documented at [https://docs.microsoft.com/en-us/data-integration/gateway/service-gateway-communication](https://docs.microsoft.com/en-us/data-integration/gateway/service-gateway-communication)

3. Connect on-premises gateway to Logic Apps.  Follow documentation: [https://docs.microsoft.com/en-us/azure/logic-apps/logic-apps-gateway-connection](https://docs.microsoft.com/en-us/azure/logic-apps/logic-apps-gateway-connection).

4. Create a custom Logic Apps connector.  Follow documentation: [https://docs.microsoft.com/en-us/connectors/custom-connectors/create-logic-apps-connector](https://docs.microsoft.com/en-us/connectors/custom-connectors/create-logic-apps-connector) and [https://docs.microsoft.com/en-us/connectors/custom-connectors/define-openapi-definition](https://docs.microsoft.com/en-us/connectors/custom-connectors/define-openapi-definition).  Make sure to enable the flag "Connect via on-premises data gateway".

    > Note:  If you don't have an Swagger definition, then create the API defintion manually using [https://docs.microsoft.com/en-us/connectors/custom-connectors/define-blank](https://docs.microsoft.com/en-us/connectors/custom-connectors/define-blank)

5. Create a Logic App in the same region as the custom Logic Apps connector.  The custom connector will now be available under the **custom** tab for use.  On the first use, you will be prompted to select the on-premises data gateway.