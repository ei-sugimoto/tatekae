import { ServiceType } from '@bufbuild/protobuf';
import { createPromiseClient, PromiseClient } from '@connectrpc/connect';
import { createConnectTransport } from '@connectrpc/connect-web';
import { useMemo } from 'react';

const baseUrl = 'http://localhost:8080';

const transport = createConnectTransport({
  baseUrl: baseUrl,
});

export function Client<T extends ServiceType>(service: T): PromiseClient<T> {
  return useMemo(() => createPromiseClient(service, transport), [service]);
}
