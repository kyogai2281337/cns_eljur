interface UserAPI {
  loginUser: (email: string, password: string) => Promise<{ status: boolean }>;
  signupUser: (email: string, password: string, firstName: string, lastName: string) => Promise<{ id: number, email: string, role: string }>;
  getUserProfile: () => Promise<{
    id: number;
    email: string;
    role: string;
    first_name: string;
    last_name: string;
  }>;
}

const userAPI: UserAPI = {
  loginUser: async (email: string, password: string) => {
    try {
      const response = await fetch('/api/auth/signin', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({ email, password })
      });
      
      if (!response.ok) {
        console.log('error auth');
        return { status: false, error: true };
      }

      const data = await response.json();

      document.cookie = `auth=${data.token}; path=/; max-age=3600`;
      
      return { status: true, error: false };
    } catch (error) {
      console.error('Error logging in:', error);
      throw error;
    }
  },

  signupUser: async (email: string, password: string, firstName: string, lastName: string) => {
    try {
      const response = await fetch('/api/auth/signup', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({
          email,
          password,
          first: firstName,
          last: lastName
        })
      });

      if (!response.ok) {
        console.log('error signing up');
        throw new Error('Failed to sign up');
      }

      const userData = await response.json();
      return userData;
    } catch (error) {
      console.error('Error signing up:', error);
      throw error;
    }
  },

  getUserProfile: async () => {
    try {
      const response = await fetch('/api/auth/private/profile', {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json'
        }
      });

      if (!response.ok) {
        if (response.status === 401) {
          // Redirect to login page or home page
          console.log('Token expired, redirecting to login');
          window.location.href = '/'; // Redirect to home page
          throw new Error('Unauthorized: Token expired');
        } else {
          console.log('error fetching profile');
          throw new Error('Failed to fetch user profile');
        }
      }

      const profileData = await response.json();
      return profileData;
    } catch (error) {
      console.error('Error fetching user profile:', error);
      throw error;
    }
  }
};

export default userAPI;
