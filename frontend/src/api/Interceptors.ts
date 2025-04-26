import axios, { AxiosInstance, AxiosRequestConfig } from 'axios';
import createAuthRefreshInterceptor from 'axios-auth-refresh';

const API_URL = import.meta.env.VITE_API_URL;

const api: AxiosInstance = axios.create({
  baseURL: API_URL,
});

// Refresh logic to handle token refresh
const refreshAuthLogic = (failedRequest: { response: { config: AxiosRequestConfig } | undefined }) => {
  if (failedRequest.response && failedRequest.response.config) {
    return axios
      // Refresh token request using cookie
      .get<{ accessToken: string }>(API_URL + '/user/refresh', { withCredentials: true })
      .then(() => {
        return Promise.resolve();
      }).catch(() => {
        return Promise.reject();
      });
  }

  return Promise.reject(new Error('No response found in failed request.'));
};

createAuthRefreshInterceptor(api, refreshAuthLogic);

export default api;
