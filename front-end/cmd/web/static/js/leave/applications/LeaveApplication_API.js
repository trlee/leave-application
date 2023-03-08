class LeaveApplication_API {
  // fetch all leave application
  async getAllLeaveApplication() {
        const url = broker + 'getAllLeaveApplication'
        const response = await fetch(url)
        const result = await response.json()
        return result
  }

  // cancel leave application
  async cancelLeaveApplication(stringifiedJSON) {
        const url = broker + 'cancelLeaveApplication'
        const body = {
            method: 'POST',
            body: stringifiedJSON,
        }
        const response = await fetch(url, body)
        const result = await response.json()
        return result
  }

  // approve leave application
  async approveLeaveApplication(stringifiedJSON) {
    const url = broker + 'approveLeaveApplication'
        const body = {
            method: 'POST',
            body: stringifiedJSON,
        }
        const response = await fetch(url, body)
        const result = await response.json()
        return result
  }

  // reject leave application
  async rejectLeaveApplication(stringifiedJSON) {
    const url = broker + 'rejectLeaveApplication'
        const body = {
            method: 'POST',
            body: stringifiedJSON,
        }
        const response = await fetch(url, body)
        const result = await response.json()
        return result
  }

  async deleteLeaveApplication(ID) {
    const url = broker + 'leaveApplicationDelete'
    const body = {
        method: 'POST',
        body: JSON.stringify(ID),
    }
    const response = await fetch(url, body)
    const result = await response.json()
    return result
  }
}