class LeaveApplication_API {
    // fetch all leave application
    async getAllLeaveApplication() {
        const url = broker + 'getAllLeaveApplication'
        const response = await fetch(url)
        const result = await response.json()
        return result
  }
}