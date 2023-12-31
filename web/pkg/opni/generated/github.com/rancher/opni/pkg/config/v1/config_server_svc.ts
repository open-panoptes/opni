// @generated by service-generator v0.0.1 with parameter "target=ts,import_extension=none,ts_nocheck=false"
// @generated from file github.com/rancher/opni/pkg/config/v1/config_server.proto (package config.v1, syntax proto3)
/* eslint-disable */

import { ConfigurationHistoryRequest, GetRequest } from "../../plugins/driverutil/types_pb";
import { GatewayConfigSpec } from "./gateway_config_pb";
import { axios } from "@pkg/opni/utils/axios";
import { DryRunRequest, DryRunResponse, HistoryResponse, ResetRequest, SetRequest } from "./config_server_pb";
import { ReactiveEvents, ReactiveWatchRequest } from "../../apis/core/v1/core_pb";
import { Socket } from "@pkg/opni/utils/socket";
import { EVENT_CONNECT_ERROR, EVENT_CONNECTED, EVENT_CONNECTING, EVENT_DISCONNECT_ERROR, EVENT_MESSAGE } from "@shell/utils/socket";


export async function GetDefaultConfiguration(input: GetRequest): Promise<GatewayConfigSpec> {
  try {
    
    if (input) {
      console.info('Here is the input for a request to GatewayConfig-GetDefaultConfiguration:', input);
    }
  
    const rawResponse = (await axios.request({
      method: 'get',
      responseType: 'arraybuffer',
      headers: {
        'Content-Type': 'application/octet-stream',
        'Accept': 'application/octet-stream',
      },
      url: `/opni-api/GatewayConfig/configuration/default`,
    data: input?.toBinary() as ArrayBuffer
    })).data;

    const response = GatewayConfigSpec.fromBinary(new Uint8Array(rawResponse));
    console.info('Here is the response for a request to GatewayConfig-GetDefaultConfiguration:', response);
    return response;
  } catch (ex: any) {
    if (ex?.response?.data) {
      const s = String.fromCharCode.apply(null, Array.from(new Uint8Array(ex?.response?.data)));
      console.error(s);
    }
    throw ex;
  }
}


export async function SetDefaultConfiguration(input: SetRequest): Promise<void> {
  try {
    
    if (input) {
      console.info('Here is the input for a request to GatewayConfig-SetDefaultConfiguration:', input);
    }
  
    const rawResponse = (await axios.request({
      method: 'put',
      responseType: 'arraybuffer',
      headers: {
        'Content-Type': 'application/octet-stream',
        'Accept': 'application/octet-stream',
      },
      url: `/opni-api/GatewayConfig/configuration/default`,
    data: input?.toBinary() as ArrayBuffer
    })).data;

    const response = rawResponse;
    console.info('Here is the response for a request to GatewayConfig-SetDefaultConfiguration:', response);
    return response;
  } catch (ex: any) {
    if (ex?.response?.data) {
      const s = String.fromCharCode.apply(null, Array.from(new Uint8Array(ex?.response?.data)));
      console.error(s);
    }
    throw ex;
  }
}


export async function GetConfiguration(input: GetRequest): Promise<GatewayConfigSpec> {
  try {
    
    if (input) {
      console.info('Here is the input for a request to GatewayConfig-GetConfiguration:', input);
    }
  
    const rawResponse = (await axios.request({
      method: 'get',
      responseType: 'arraybuffer',
      headers: {
        'Content-Type': 'application/octet-stream',
        'Accept': 'application/octet-stream',
      },
      url: `/opni-api/GatewayConfig/configuration`,
    data: input?.toBinary() as ArrayBuffer
    })).data;

    const response = GatewayConfigSpec.fromBinary(new Uint8Array(rawResponse));
    console.info('Here is the response for a request to GatewayConfig-GetConfiguration:', response);
    return response;
  } catch (ex: any) {
    if (ex?.response?.data) {
      const s = String.fromCharCode.apply(null, Array.from(new Uint8Array(ex?.response?.data)));
      console.error(s);
    }
    throw ex;
  }
}


export async function SetConfiguration(input: SetRequest): Promise<void> {
  try {
    
    if (input) {
      console.info('Here is the input for a request to GatewayConfig-SetConfiguration:', input);
    }
  
    const rawResponse = (await axios.request({
      method: 'put',
      responseType: 'arraybuffer',
      headers: {
        'Content-Type': 'application/octet-stream',
        'Accept': 'application/octet-stream',
      },
      url: `/opni-api/GatewayConfig/configuration`,
    data: input?.toBinary() as ArrayBuffer
    })).data;

    const response = rawResponse;
    console.info('Here is the response for a request to GatewayConfig-SetConfiguration:', response);
    return response;
  } catch (ex: any) {
    if (ex?.response?.data) {
      const s = String.fromCharCode.apply(null, Array.from(new Uint8Array(ex?.response?.data)));
      console.error(s);
    }
    throw ex;
  }
}


export async function ResetDefaultConfiguration(): Promise<void> {
  try {
    
    const rawResponse = (await axios.request({
      method: 'post',
      responseType: 'arraybuffer',
      headers: {
        'Content-Type': 'application/octet-stream',
        'Accept': 'application/octet-stream',
      },
      url: `/opni-api/GatewayConfig/configuration/default/reset`
    })).data;

    const response = rawResponse;
    console.info('Here is the response for a request to GatewayConfig-ResetDefaultConfiguration:', response);
    return response;
  } catch (ex: any) {
    if (ex?.response?.data) {
      const s = String.fromCharCode.apply(null, Array.from(new Uint8Array(ex?.response?.data)));
      console.error(s);
    }
    throw ex;
  }
}


export async function ResetConfiguration(input: ResetRequest): Promise<void> {
  try {
    
    if (input) {
      console.info('Here is the input for a request to GatewayConfig-ResetConfiguration:', input);
    }
  
    const rawResponse = (await axios.request({
      method: 'post',
      responseType: 'arraybuffer',
      headers: {
        'Content-Type': 'application/octet-stream',
        'Accept': 'application/octet-stream',
      },
      url: `/opni-api/GatewayConfig/configuration/reset`,
    data: input?.toBinary() as ArrayBuffer
    })).data;

    const response = rawResponse;
    console.info('Here is the response for a request to GatewayConfig-ResetConfiguration:', response);
    return response;
  } catch (ex: any) {
    if (ex?.response?.data) {
      const s = String.fromCharCode.apply(null, Array.from(new Uint8Array(ex?.response?.data)));
      console.error(s);
    }
    throw ex;
  }
}


export async function DryRun(input: DryRunRequest): Promise<DryRunResponse> {
  try {
    
    if (input) {
      console.info('Here is the input for a request to GatewayConfig-DryRun:', input);
    }
  
    const rawResponse = (await axios.request({
      method: 'post',
      responseType: 'arraybuffer',
      headers: {
        'Content-Type': 'application/octet-stream',
        'Accept': 'application/octet-stream',
      },
      url: `/opni-api/GatewayConfig/dry-run`,
    data: input?.toBinary() as ArrayBuffer
    })).data;

    const response = DryRunResponse.fromBinary(new Uint8Array(rawResponse));
    console.info('Here is the response for a request to GatewayConfig-DryRun:', response);
    return response;
  } catch (ex: any) {
    if (ex?.response?.data) {
      const s = String.fromCharCode.apply(null, Array.from(new Uint8Array(ex?.response?.data)));
      console.error(s);
    }
    throw ex;
  }
}


export async function ConfigurationHistory(input: ConfigurationHistoryRequest): Promise<HistoryResponse> {
  try {
    
    if (input) {
      console.info('Here is the input for a request to GatewayConfig-ConfigurationHistory:', input);
    }
  
    const rawResponse = (await axios.request({
      method: 'get',
      responseType: 'arraybuffer',
      headers: {
        'Content-Type': 'application/octet-stream',
        'Accept': 'application/octet-stream',
      },
      url: `/opni-api/GatewayConfig/configuration/history`,
    data: input?.toBinary() as ArrayBuffer
    })).data;

    const response = HistoryResponse.fromBinary(new Uint8Array(rawResponse));
    console.info('Here is the response for a request to GatewayConfig-ConfigurationHistory:', response);
    return response;
  } catch (ex: any) {
    if (ex?.response?.data) {
      const s = String.fromCharCode.apply(null, Array.from(new Uint8Array(ex?.response?.data)));
      console.error(s);
    }
    throw ex;
  }
}


export function WatchReactive(input: ReactiveWatchRequest, callback: (data: ReactiveEvents) => void): () => Promise<any> {
  const socket = new Socket('/opni-api/GatewayConfig/configuration/watch', true);
  Object.assign(socket, { frameTimeout: null })
  socket.addEventListener(EVENT_MESSAGE, (e: any) => {
    const event = e.detail;
    if (event.data) {
      callback(ReactiveEvents.fromBinary(new Uint8Array(event.data)));
    }
  });
  socket.addEventListener(EVENT_CONNECTING, () => {
    if (socket.socket) {
      socket.socket.binaryType = 'arraybuffer';
    }
  });
  socket.addEventListener(EVENT_CONNECTED, () => {
    socket.send(input.toBinary());
  });
  socket.addEventListener(EVENT_CONNECT_ERROR, (e) => {
    console.error(e);
  })
  socket.addEventListener(EVENT_DISCONNECT_ERROR, (e) => {
    console.error(e);
  })
  socket.connect();
  return () => {
    return socket.disconnect(null);
  };
}

