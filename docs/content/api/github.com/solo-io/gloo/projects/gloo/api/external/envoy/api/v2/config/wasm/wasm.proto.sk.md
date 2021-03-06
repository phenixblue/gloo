
---
title: "wasm.proto"
weight: 5
---

<!-- Code generated by solo-kit. DO NOT EDIT. -->


### Package: `envoy.config.wasm.v2` 
#### Types:


- [VmConfig](#vmconfig)
- [PluginConfig](#pluginconfig)
- [WasmService](#wasmservice)
  



##### Source File: [github.com/solo-io/gloo/projects/gloo/api/external/envoy/api/v2/config/wasm/wasm.proto](https://github.com/solo-io/gloo/blob/master/projects/gloo/api/external/envoy/api/v2/config/wasm/wasm.proto)





---
### VmConfig

 
Configuration for a Wasm VM.
[#next-free-field: 6]

```yaml
"vmId": string
"runtime": string
"code": .envoy.api.v2.core.AsyncDataSource
"configuration": string
"allowPrecompiled": bool

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `vmId` | `string` | An ID which will be used along with a hash of the wasm code (or null_vm_id) to determine which VM will be used for the plugin. All plugins which use the same vm_id and code will use the same VM. May be left blank. |  |
| `runtime` | `string` | The Wasm runtime type (see source/extensions/commmon/wasm/well_known_names.h). |  |
| `code` | [.envoy.api.v2.core.AsyncDataSource](../../../../../../../../../../../../../envoy/api/v2/core/base.proto.sk/#asyncdatasource) | The Wasm code that Envoy will execute. |  |
| `configuration` | `string` | The Wasm configuration string used on initialization of a new VM (proxy_onStart). |  |
| `allowPrecompiled` | `bool` | Allow the wasm file to include pre-compiled code on VMs which support it. |  |




---
### PluginConfig

 
Base Configuration for Wasm Plugins, e.g. filters and services.

```yaml
"name": string
"rootId": string
"vmConfig": .envoy.config.wasm.v2.VmConfig
"configuration": string

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `name` | `string` | A unique name for a filters/services in a VM for use in identifiying the filter/service if multiple filters/services are handled by the same vm_id and root_id and for logging/debugging. |  |
| `rootId` | `string` | A unique ID for a set of filters/services in a VM which will share a RootContext and Contexts if applicable (e.g. an Wasm HttpFilter and an Wasm AccessLog). If left blank, all filters/services with a blank root_id with the same vm_id will share Context(s). |  |
| `vmConfig` | [.envoy.config.wasm.v2.VmConfig](../wasm.proto.sk/#vmconfig) | Configuration for finding or starting VM. |  |
| `configuration` | `string` | Filter/service configuration string e.g. a serialized protobuf which will be the argument to the proxy_onConfigure() call. |  |




---
### WasmService

 
WasmService is configured as a built-in *envoy.wasm_service* :ref:`ServiceConfig
<envoy_api_msg_config.wasm.v2.ServiceConfig>`. This opaque configuration will be used to
create a Wasm Service.

```yaml
"config": .envoy.config.wasm.v2.PluginConfig
"singleton": bool
"statPrefix": string

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `config` | [.envoy.config.wasm.v2.PluginConfig](../wasm.proto.sk/#pluginconfig) | General plugin configuration. |  |
| `singleton` | `bool` | If true, create a single VM rather than creating one VM per silo. Such a singleton can not be used with filters. |  |
| `statPrefix` | `string` | If set add 'stat_prefix' as a prefix to all stats. |  |





<!-- Start of HubSpot Embed Code -->
<script type="text/javascript" id="hs-script-loader" async defer src="//js.hs-scripts.com/5130874.js"></script>
<!-- End of HubSpot Embed Code -->
