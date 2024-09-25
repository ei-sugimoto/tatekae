import { ServiceType } from '@bufbuild/protobuf';
import {
  createPromiseClient,
  Interceptor,
  PromiseClient,
} from '@connectrpc/connect';
import { createConnectTransport } from '@connectrpc/connect-web';
import { getCookie } from './utils/cookie';

const baseUrl = 'http://localhost:8080';

const authInterceptor: Interceptor = (next) => async (request) => {
  const token = getCookie('token');
  if (token) {
    request.header.set('Authorization', `Bearer ${token}`);
  }
  return next(request);
};

const transport = createConnectTransport({
  baseUrl: baseUrl,
  interceptors: [authInterceptor],
});

export function Client<T extends ServiceType>(service: T): PromiseClient<T> {
  return createPromiseClient(service, transport);
}
