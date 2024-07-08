interface UserAPI {
    loginUser: (email: string, password: string) => Promise<{ status: boolean }>
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
        })
        
        if (!response.ok) {
          console.log('error auth')
          return { status: false }
        }

        const data = await response.json()

        localStorage.setItem('token', data.token);
        
        return { status: true }
      } catch (error) {
        console.error('Error logging in:', error)
        throw error
      }
    }
  }
  
  export default userAPI
  