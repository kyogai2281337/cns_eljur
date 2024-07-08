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
          return { status: false, error: true }
        }

        const data = await response.json()

        document.cookie = `auth=${data.token}; path=/; max-age=7200`
        
        return { status: true, error: false }
      } catch (error) {
        console.error('Error logging in:', error)
        throw error
      }
    }
  }
  
  export default userAPI
  