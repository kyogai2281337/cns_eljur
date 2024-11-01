import axios from "axios";

export const backendURL = process.env.VUE_APP_BACKEND_URL || "/api";
const baseUrl = backendURL + "/auth";

interface UserResponse {
  token: string;
  user: {
    id: string;
    email: string;
    firstName: string;
    lastName: string;
  };
}

const signin = async (
  email: string,
  password: string
): Promise<{ data?: UserResponse; error: boolean; errorMsg?: string }> => {
  try {
    const response = await axios.post<UserResponse>(baseUrl + "/signin", {
      email: email,
      password: password,
    });
    return {
      data: response.data,
      error: false,
    };
  } catch (error) {
    if (axios.isAxiosError(error)) {
      return {
        error: true,
        errorMsg: error.response?.data?.message || "Authentication failed",
      };
    }
    return {
      error: true,
      errorMsg: "Unexpected error occurred",
    };
  }
};

const signup = async (
  email: string,
  password: string,
  first: string,
  last: string
): Promise<{ data?: UserResponse; error: boolean; errorMsg?: string }> => {
  try {
    const response = await axios.post<UserResponse>(baseUrl + "/signup", {
      email: email,
      password: password,
      first: first,
      last: last,
    });
    return {
      data: response.data,
      error: false,
    };
  } catch (error) {
    if (axios.isAxiosError(error)) {
      return {
        error: true,
        errorMsg: error.response?.data?.message || "Registration failed",
      };
    }
    return {
      error: true,
      errorMsg: "Unexpected error occurred",
    };
  }
};

const getProfile = async (): Promise<{
  data?: UserResponse;
  error: boolean;
  errorMsg?: string;
}> => {
  try {
    const response = await axios.get<UserResponse>(
      baseUrl + "/private/profile"
    );
    return {
      data: response.data,
      error: false,
    };
  } catch (error) {
    if (axios.isAxiosError(error)) {
      return {
        error: true,
        errorMsg: error.response?.data?.message || "Failed to fetch profile",
      };
    }
    return {
      error: true,
      errorMsg: "Unexpected error occurred",
    };
  }
};

export default {
  signin,
  signup,
  getProfile,
};
