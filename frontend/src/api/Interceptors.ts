import axios, { AxiosInstance, AxiosRequestConfig } from 'axios';
import createAuthRefreshInterceptor from 'axios-auth-refresh';


const api: AxiosInstance = axios.create({
  baseURL: 'http://localhost:8080/api',
});

// Refresh logic to handle token refresh
const refreshAuthLogic = (failedRequest: { response: { config: AxiosRequestConfig } | undefined }) => {
  if (failedRequest.response && failedRequest.response.config) {
    return axios
      // Refresh token request using cookie
      .get<{ accessToken: string }>('http://localhost:8080/api/user/refresh', { withCredentials: true })
      .then((response) => {
        console.log(response);
        return Promise.resolve();
      }).catch((err) => {
        console.log(err);
        return Promise.reject();
      });
  }

  return Promise.reject(new Error('No response found in failed request.'));
};

createAuthRefreshInterceptor(api, refreshAuthLogic);

export default api;
