interface ApiResponse<T> {
  data: T | null;
  error: boolean;
  error_msg: string | null;
}

interface SignupData {
  email: string;
  password: string;
  first: string;
  last: string;
}

interface SigninData {
  email: string;
  password: string;
}

interface SignupResponse {
  id: number;
  email: string;
  role: string;
}

interface SigninResponse {
  status: number;
  token: string;
}

interface ProfileResponse {
  id: number;
  email: string;
  role: string;
  first_name: string;
  last_name: string;
}

const backend = "http://localhost"

const useUser = () => {
  const signup = async (data: SignupData): Promise<ApiResponse<SignupResponse>> => {
    try {
      const response = await fetch(backend+'/api/auth/signup', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(data),
        credentials: 'include'
      });

      if (!response.ok) {
        const error = await response.json();
        return { data: null, error: true, error_msg: error.message || 'Signup failed' };
      }

      const result: SignupResponse = await response.json();
      return { data: result, error: false, error_msg: null };
    } catch (error: unknown) {
      if (error instanceof Error) {
        return { data: null, error: true, error_msg: error.message || 'Network error' };
      }
      return { data: null, error: true, error_msg: 'Unknown error occurred' };
    }
  };

  const signin = async (data: SigninData): Promise<ApiResponse<SigninResponse>> => {
    try {
      const response = await fetch(backend+'/api/auth/signin', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(data),
        credentials: 'include'
      });

      if (!response.ok) {
        const error = await response.json();
        return { data: null, error: true, error_msg: error.message || 'Signin failed' };
      }

      const result: SigninResponse = await response.json();
      return { data: result, error: false, error_msg: null };
    } catch (error: unknown) {
      if (error instanceof Error) {
        return { data: null, error: true, error_msg: error.message || 'Network error' };
      }
      return { data: null, error: true, error_msg: 'Unknown error occurred' };
    }
  };

  const getProfile = async (): Promise<ApiResponse<ProfileResponse>> => {
    try {
      const response = await fetch(backend+'/api/auth/private/profile', {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
          'AuthToken': localStorage.getItem('token') || '',
        },
      });

      if (!response.ok) {
        const error = await response.json();
        return { data: null, error: true, error_msg: error.message || 'Failed to fetch profile' };
      }

      const result: ProfileResponse = await response.json();
      return { data: result, error: false, error_msg: null };
    } catch (error: unknown) {
      if (error instanceof Error) {
        return { data: null, error: true, error_msg: error.message || 'Network error' };
      }
      return { data: null, error: true, error_msg: 'Unknown error occurred' };
    }
  };

  return {
    signup,
    signin,
    getProfile
  };
};

export default useUser;
