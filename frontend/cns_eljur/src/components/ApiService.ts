export default class ApiService {
  static async get<T>(endpoint: string): Promise<T> {
    try {
      const response = await fetch(`${endpoint}`)
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`)
      }
      return (await response.json()) as T
    } catch (error) {
      console.error('Error fetching data:', error)
      throw error
    }
  }

  static async post<T>(endpoint: string, data: any): Promise<T> {
    try {
      const response = await fetch(`${endpoint}`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(data)
      })
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`)
      }
      return (await response.json()) as T
    } catch (error) {
      console.error('Error posting data:', error)
      throw error
    }
  }
}